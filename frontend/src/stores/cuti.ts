import { defineStore } from 'pinia'

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

let nextCutiId = 1

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
        await new Promise((r) => setTimeout(r, 300))
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
        const created: Cuti = { id: nextCutiId++, ...cuti }
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
        const index = this.items.findIndex((c) => c.id === id)
        if (index === -1) {
          throw new Error('Cuti tidak ditemukan')
        }
        const updated: Cuti = { id, ...cuti }
        this.items[index] = updated
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
        const cuti = this.items.find((c) => c.id === id)
        if (!cuti) {
          throw new Error('Cuti tidak ditemukan')
        }
        return cuti
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail cuti'
        throw err
      } finally {
        this.loading = false
      }
    },

    async setStatus(id: number, status: StatusCuti, catatan?: string, disetujuiOleh?: number) {
      this.loading = true
      this.error = ''
      try {
        const index = this.items.findIndex((c) => c.id === id)
        if (index === -1) {
          throw new Error('Cuti tidak ditemukan')
        }
        this.items[index].status = status
        if (catatan !== undefined) {
          this.items[index].catatanApproval = catatan
        }
        if (disetujuiOleh !== undefined) {
          this.items[index].disetujuiOleh = disetujuiOleh
        }
        return this.items[index]
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate status cuti'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
