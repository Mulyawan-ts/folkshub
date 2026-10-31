import { defineStore } from "pinia";
import {
  SimpanAbsensi,
  SemuaAbsensi,
  HapusAbsensi,
  UpdateAbsensi,
  AmbilAbsensi,
} from "../../wailsjs/go/main/App";

export type StatusAbsensi = "hadir" | "izin" | "sakit" | "alfa" | "lembur";

export interface Absensi {
  id: number;
  karyawan_id: number;
  tanggal: string;
  check_in?: string;
  check_out?: string;
  status: string;
  keterangan?: string;
}

export const useAbsensiStore = defineStore("absensi", {
  state: () => ({
    items: [] as Absensi[],
    loading: false,
    error: "" as string,
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((a) => a.id === id),
    byKaryawan: (state) => (karyawan_id: number) =>
      state.items.filter((a) => a.karyawan_id === karyawan_id),
    byTanggal: (state) => (tanggal: string) =>
      state.items.filter((a) => a.tanggal === tanggal),
    hadirCount: (state) =>
      state.items.filter((a) => a.status === "hadir").length,
    alfaCount: (state) => state.items.filter((a) => a.status === "alfa").length,
  },

  actions: {
    async fetchAll() {
      this.loading = true;
      this.error = "";
      try {
        this.items = await SemuaAbsensi();
      } catch (err: any) {
        this.error = err?.message || "Gagal memuat data absensi";
      } finally {
        this.loading = false;
      }
    },

    async simpan(absensi: Omit<Absensi, "id">) {
      this.loading = true;
      this.error = "";
      try {
        const created = await SimpanAbsensi(
          absensi.karyawan_id,
          absensi.tanggal,
          absensi.check_in || "",
          absensi.check_out || "",
          absensi.status,
          absensi.keterangan || "",
        );
        this.items.push(created);
        return created;
      } catch (err: any) {
        this.error = err?.message || "Gagal menyimpan absensi";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async hapus(id: number) {
      this.loading = true;
      this.error = "";
      try {
        await HapusAbsensi(id);
        this.items = this.items.filter((a) => a.id !== id);
      } catch (err: any) {
        this.error = err?.message || "Gagal menghapus absensi";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async update(id: number, absensi: Omit<Absensi, "id">) {
      this.loading = true;
      this.error = "";
      try {
        const updated = await UpdateAbsensi(
          id,
          absensi.karyawan_id,
          absensi.tanggal,
          absensi.check_in || "",
          absensi.check_out || "",
          absensi.status,
          absensi.keterangan || "",
        );
        const index = this.items.findIndex((a) => a.id === id);
        if (index !== -1) {
          this.items[index] = updated;
        }
        return updated;
      } catch (err: any) {
        this.error = err?.message || "Gagal mengupdate absensi";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async ambil(id: number) {
      this.loading = true;
      this.error = "";
      try {
        return await AmbilAbsensi(id);
      } catch (err: any) {
        this.error = err?.message || "Gagal memuat detail absensi";
        throw err;
      } finally {
        this.loading = false;
      }
    },
  },
});
