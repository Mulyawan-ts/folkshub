import { defineStore } from 'pinia'
import { SimpanCuti, SemuaCuti, HapusCuti, UpdateCuti, AmbilCuti } from '../../wailsjs/go/main/App'

export type StatusCuti = 'draft' | 'submitted' | 'approved' | 'rejected' | 'cancelled'

export type JenisCuti = 'tahunan' | 'sakit' | 'izin' | 'melahirkan' | 'lainnya'

export interface Cuti {
  id: number
  karyawanId: number
  jenis: JenisCuti
  tanggalMulai: string
  tanggalSelesai: string
  jumlahHari: number
  keterangan: string
  status: StatusCuti
  catatanApproval?: string
  disetujuiOleh?: number
}

export const useCutiStore = defineStore('cuti', {
  state: () => ({
    items: [] as Cuti[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((c) => c.id === id),
    byKaryawan: (state) => (karyawanId: number) =>
      state.items.filter((c) => c.karyawanId === karyawanId),
    byStatus: (state) => (status: StatusCuti) =>
      state.items.filter((c) => c.status === status),
    byJenis: (state) => (jenis: JenisCuti) =>
      state.items.filter((c) => c.jenis === jenis),
    pendingCount: (state) =>
      state.items.filter((c) => c.status === 'submitted').length,
    approvedCount: (state) =>
      state.items.filter((c) => c.status === 'approved').length
  },

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        this.items = await SemuaCuti()
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat data cuti'
      } finally {
        this.loading = false
      }
    },

    async simpan(cuti: Omit<Cuti, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const created = await SimpanCuti(
          cuti.karyawanId,
          cuti.jenis,
          cuti.tanggalMulai,
          cuti.tanggalSelesai,
          cuti.keterangan,
          cuti.status,
          cuti.jumlahHari
        )
        this.items.push(created)
        return created
      } catch (err: any) {
        this.error = err?.message || 'Gagal menyimpan cuti'
        throw err
      } finally {
        this.loading = false
      }
    },

    async hapus(id: number) {
      this.loading = true
      this.error = ''
      try {
        await HapusCuti(id)
        this.items = this.items.filter((c) => c.id !== id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal menghapus cuti'
        throw err
      } finally {
        this.loading = false
      }
    },

    async update(id: number, cuti: Omit<Cuti, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const updated = await UpdateCuti(
          id,
          cuti.karyawanId,
          cuti.jenis,
          cuti.tanggalMulai,
          cuti.tanggalSelesai,
          cuti.keterangan,
          cuti.status,
          cuti.jumlahHari
        )
        const index = this.items.findIndex((c) => c.id === id)
        if (index !== -1) {
          this.items[index] = updated
        }
        return updated
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate cuti'
        throw err
      } finally {
        this.loading = false
      }
    },

    async ambil(id: number) {
      this.loading = true
      this.error = ''
      try {
        return await AmbilCuti(id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail cuti'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
