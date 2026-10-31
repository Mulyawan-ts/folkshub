import { defineStore } from "pinia";
import {
  SimpanReimburse,
  SemuaReimburse,
  HapusReimburse,
  UpdateReimburse,
} from "../../wailsjs/go/main/App";

export type StatusReimburse =
  "draft" | "submitted" | "approved" | "rejected" | "paid";

export interface Reimburse {
  id: number;
  karyawan_id: number;
  tanggal: string;
  jenis: string;
  jumlah: number;
  keterangan: string;
  bukti: string;
  status: string;
  catatan_approval?: string;
}

export const useReimburseStore = defineStore("reimburse", {
  state: () => ({
    items: [] as Reimburse[],
    loading: false,
    error: "" as string,
  }),

  actions: {
    async fetchAll() {
      this.loading = true;
      this.error = "";
      try {
        this.items = await SemuaReimburse();
      } catch (err: any) {
        this.error = err?.message || "Gagal memuat data reimburse";
      } finally {
        this.loading = false;
      }
    },

    async simpan(reimburse: Omit<Reimburse, "id">) {
      this.loading = true;
      this.error = "";
      try {
        const created = await SimpanReimburse(
          reimburse.karyawan_id,
          reimburse.tanggal,
          reimburse.jenis,
          reimburse.keterangan,
          reimburse.bukti,
          reimburse.status,
          reimburse.jumlah,
        );
        this.items.push(created);
        return created;
      } catch (err: any) {
        this.error = err?.message || "Gagal menyimpan reimburse";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async hapus(id: number) {
      this.loading = true;
      this.error = "";
      try {
        await HapusReimburse(id);
        this.items = this.items.filter((r) => r.id !== id);
      } catch (err: any) {
        this.error = err?.message || "Gagal menghapus reimburse";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async update(id: number, reimburse: Omit<Reimburse, "id">) {
      this.loading = true;
      this.error = "";
      try {
        const updated = await UpdateReimburse(
          id,
          reimburse.karyawan_id,
          reimburse.tanggal,
          reimburse.jenis,
          reimburse.keterangan,
          reimburse.bukti,
          reimburse.status,
          reimburse.jumlah,
        );
        const index = this.items.findIndex((r) => r.id === id);
        if (index !== -1) {
          this.items[index] = updated;
        }
        return updated;
      } catch (err: any) {
        this.error = err?.message || "Gagal mengupdate reimburse";
        throw err;
      } finally {
        this.loading = false;
      }
    },
  },
});
