import { defineStore } from "pinia";
import {
  SimpanDepartemen,
  SemuaDepartemen,
  HapusDepartemen,
  UpdateDepartemen,
} from "../../wailsjs/go/main/App";

export interface Departemen {
  id: number;
  kode: string;
  nama: string;
  kepala: string;
  deskripsi: string;
}

export const useDepartemenStore = defineStore("departemen", {
  state: () => ({
    items: [] as Departemen[],
    loading: false,
    error: "" as string,
  }),

  actions: {
    async fetchAll() {
      this.loading = true;
      this.error = "";
      try {
        this.items = await SemuaDepartemen();
      } catch (err: any) {
        this.error = err?.message || "Gagal memuat data departemen";
      } finally {
        this.loading = false;
      }
    },

    async simpan(departemen: Omit<Departemen, "id">) {
      this.loading = true;
      this.error = "";
      try {
        const created = await SimpanDepartemen(
          departemen.kode,
          departemen.nama,
          departemen.kepala,
          departemen.deskripsi,
        );
        this.items.push(created);
        return created;
      } catch (err: any) {
        this.error = err?.message || "Gagal menyimpan departemen";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async hapus(id: number) {
      this.loading = true;
      this.error = "";
      try {
        await HapusDepartemen(id);
        this.items = this.items.filter((d) => d.id !== id);
      } catch (err: any) {
        this.error = err?.message || "Gagal menghapus departemen";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async update(id: number, departemen: Omit<Departemen, "id">) {
      this.loading = true;
      this.error = "";
      try {
        const updated = await UpdateDepartemen(
          id,
          departemen.kode,
          departemen.nama,
          departemen.kepala,
          departemen.deskripsi,
        );
        const index = this.items.findIndex((d) => d.id === id);
        if (index !== -1) {
          this.items[index] = updated;
        }
        return updated;
      } catch (err: any) {
        this.error = err?.message || "Gagal mengupdate departemen";
        throw err;
      } finally {
        this.loading = false;
      }
    },
  },
});
