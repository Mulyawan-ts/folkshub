import { defineStore } from 'pinia'
import { SimpanCuti, SemuaCuti, HapusCuti, UpdateCuti, AmbilCuti } from '../../wailsjs/go/main/App'

export type StatusCuti = 'draft' | 'submitted' | 'approved' | 'rejected' | 'cancelled'

export type JenisCuti = 'tahunan' | 'sakit' | 'izin' | 'melahirkan' | 'lainnya'

export interface Cuti {
  id: number
  karyawan_id: number
  jenis: string
  tanggal_mulai: string
  tanggal_selesai: string
  jumlah_hari: number
  keterangan: string
  status: string
  catatan_approval?: string
  disetujui_oleh?: number
}

export const useCutiStore = defineStore('cuti', {
  state: () => ({
    items: [] as Cuti[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((c) => c.id === id),
    byKaryawan: (state) => (karyawan_id: number) =>
      state.items.filter((c) => c.karyawan_id === karyawan_id),
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
          cuti.karyawan_id,
          cuti.jenis,
          cuti.tanggal_mulai,
          cuti.tanggal_selesai,
          cuti.keterangan,
          cuti.status,
          cuti.jumlah_hari
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
          cuti.karyawan_id,
          cuti.jenis,
          cuti.tanggal_mulai,
          cuti.tanggal_selesai,
          cuti.keterangan,
          cuti.status,
          cuti.jumlah_hari
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
