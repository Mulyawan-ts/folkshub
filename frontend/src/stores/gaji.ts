import { defineStore } from "pinia";
import {
  SimpanGaji,
  SemuaGaji,
  HapusGaji,
  UpdateGaji,
  AmbilGaji,
} from "../../wailsjs/go/main/App";

export type StatusGaji = "draft" | "diproses" | "dibayar" | "ditolak";

export interface Gaji {
  id: number;
  karyawan_id: number;
  bulan: string;
  gaji_pokok: number;
  tunjangan: number;
  potongan: number;
  total?: number;
  status: string;
  tanggal_bayar?: string;
  keterangan?: string;
}

export const useGajiStore = defineStore("gaji", {
  state: () => ({
    items: [] as Gaji[],
    loading: false,
    error: "" as string,
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((g) => g.id === id),
    byKaryawan: (state) => (karyawan_id: number) =>
      state.items.filter((g) => g.karyawan_id === karyawan_id),
    byBulan: (state) => (bulan: string) =>
      state.items.filter((g) => g.bulan === bulan),
    byStatus: (state) => (status: StatusGaji) =>
      state.items.filter((g) => g.status === status),
    totalGaji: (state) =>
      state.items.reduce((acc, g) => acc + (g.total ?? 0), 0),
    totalDibayar: (state) =>
      state.items
        .filter((g) => g.status === "dibayar")
        .reduce((acc, g) => acc + (g.total ?? 0), 0),
    belumDibayarCount: (state) =>
      state.items.filter((g) => g.status !== "dibayar").length,
  },

  actions: {
    async fetchAll() {
      this.loading = true;
      this.error = "";
      try {
        this.items = await SemuaGaji();
      } catch (err: any) {
        this.error = err?.message || "Gagal memuat data gaji";
      } finally {
        this.loading = false;
      }
    },

    async simpan(gaji: Omit<Gaji, "id">) {
      this.loading = true;
      this.error = "";
      try {
        const created = await SimpanGaji(
          gaji.karyawan_id,
          gaji.bulan,
          gaji.gaji_pokok,
          gaji.tunjangan,
          gaji.potongan,
          gaji.status,
        );
        this.items.push(created);
        return created;
      } catch (err: any) {
        this.error = err?.message || "Gagal menyimpan gaji";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async hapus(id: number) {
      this.loading = true;
      this.error = "";
      try {
        await HapusGaji(id);
        this.items = this.items.filter((g) => g.id !== id);
      } catch (err: any) {
        this.error = err?.message || "Gagal menghapus gaji";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async update(id: number, gaji: Omit<Gaji, "id">) {
      this.loading = true;
      this.error = "";
      try {
        const updated = await UpdateGaji(
          id,
          gaji.karyawan_id,
          gaji.bulan,
          gaji.gaji_pokok,
          gaji.tunjangan,
          gaji.potongan,
          gaji.status,
        );
        const index = this.items.findIndex((g) => g.id === id);
        if (index !== -1) {
          this.items[index] = updated;
        }
        return updated;
      } catch (err: any) {
        this.error = err?.message || "Gagal mengupdate gaji";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async ambil(id: number) {
      this.loading = true;
      this.error = "";
      try {
        return await AmbilGaji(id);
      } catch (err: any) {
        this.error = err?.message || "Gagal memuat detail gaji";
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async setStatus(id: number, status: string) {
      this.loading = true;
      this.error = "";
      try {
        const updated = await UpdateGaji(id, 0, "", 0, 0, 0, status);
        const index = this.items.findIndex((g) => g.id === id);
        if (index !== -1) {
          this.items[index] = updated;
        }
        return this.items[index];
      } catch (err: any) {
        this.error = err?.message || "Gagal mengupdate status gaji";
        throw err;
      } finally {
        this.loading = false;
      }
    },
  },
});
