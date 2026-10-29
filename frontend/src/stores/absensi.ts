import { defineStore } from 'pinia'

export type StatusAbsensi = 'hadir' | 'izin' | 'sakit' | 'alfa' | 'lembur'

export interface Absensi {
  id: number
  karyawanId: number
  tanggal: string
  checkIn?: string
  checkOut?: string
  status: StatusAbsensi
  keterangan?: string
}

let nextAbsensiId = 1

export const useAbsensiStore = defineStore('absensi', {
  state: () => ({
    items: [] as Absensi[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((a) => a.id === id),
    byKaryawan: (state) => (karyawanId: number) =>
      state.items.filter((a) => a.karyawanId === karyawanId),
    byTanggal: (state) => (tanggal: string) =>
      state.items.filter((a) => a.tanggal === tanggal),
    hadirCount: (state) => state.items.filter((a) => a.status === 'hadir').length,
    alfaCount: (state) => state.items.filter((a) => a.status === 'alfa').length
  },

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        await new Promise((r) => setTimeout(r, 300))
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat data absensi'
      } finally {
        this.loading = false
      }
    },

    async simpan(absensi: Omit<Absensi, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const created: Absensi = { id: nextAbsensiId++, ...absensi }
        this.items.push(created)
        return created
      } catch (err: any) {
        this.error = err?.message || 'Gagal menyimpan absensi'
        throw err
      } finally {
        this.loading = false
      }
    },

    async hapus(id: number) {
      this.loading = true
      this.error = ''
      try {
        this.items = this.items.filter((a) => a.id !== id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal menghapus absensi'
        throw err
      } finally {
        this.loading = false
      }
    },

    async update(id: number, absensi: Omit<Absensi, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const index = this.items.findIndex((a) => a.id === id)
        if (index === -1) {
          throw new Error('Absensi tidak ditemukan')
        }
        const updated: Absensi = { id, ...absensi }
        this.items[index] = updated
        return updated
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate absensi'
        throw err
      } finally {
        this.loading = false
      }
    },

    async ambil(id: number) {
      this.loading = true
      this.error = ''
      try {
        const absensi = this.items.find((a) => a.id === id)
        if (!absensi) {
          throw new Error('Absensi tidak ditemukan')
        }
        return absensi
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail absensi'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
