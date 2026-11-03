import { defineStore } from "pinia";
import { RingkasanDashboard } from "../../wailsjs/go/main/App";

export interface Ringkasan {
  totalKaryawan: number;
  karyawanTetap: number;
  karyawanOutsourcing: number;
  karyawanMagang: number;
  totalAbsensi: number;
  totalCuti: number;
  cutiPending: number;
  totalReimburse: number;
  reimbursePending: number;
  totalGaji: number;
  totalJabatan: number;
  totalDepartemen: number;
  totalPelatihan: number;
}

export const useDashboardStore = defineStore("dashboard", {
  state: () => ({
    data: null as Ringkasan | null,
    loading: false,
    error: "" as string,
  }),

  actions: {
    bersihkanError() {
      this.error = "";
    },
    async fetch() {
      this.loading = true;
      this.error = "";
      try {
        this.data = (await RingkasanDashboard()) as Ringkasan;
      } catch (err: any) {
        this.error = err?.message || "Gagal memuat ringkasan";
      } finally {
        this.loading = false;
      }
    },
  },
});
