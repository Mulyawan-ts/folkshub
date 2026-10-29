import { defineStore } from 'pinia'

export type StatusPelatihan = 'terjadwal' | 'berlangsung' | 'selesai' | 'dibatalkan'

export interface Pelatihan {
  id: number
  karyawanId: number
  nama: string
  jenis: string
  tanggalMulai: string
  tanggalSelesai: string
  lokasi: string
  biaya: number
  status: StatusPelatihan
  sertifikat: string
  catatan: string
}

let nextPelatihanId = 1

export const usePelatihanStore = defineStore('pelatihan', {
  state: () => ({
    items: [] as Pelatihan[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((p) => p.id === id),
    byKaryawan: (state) => (karyawanId: number) =>
      state.items.filter((p) => p.karyawanId === karyawanId),
    byStatus: (state) => (status: StatusPelatihan) =>
      state.items.filter((p) => p.status === status)
  },

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        await new Promise((r) => setTimeout(r, 300))
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat data pelatihan'
      } finally {
        this.loading = false
      }
    },

    async simpan(pelatihan: Omit<Pelatihan, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const created: Pelatihan = { id: nextPelatihanId++, ...pelatihan }
        this.items.push(created)
        return created
      } catch (err: any) {
        this.error = err?.message || 'Gagal menyimpan pelatihan'
        throw err
      } finally {
        this.loading = false
      }
    },

    async hapus(id: number) {
      this.loading = true
      this.error = ''
      try {
        this.items = this.items.filter((p) => p.id !== id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal menghapus pelatihan'
        throw err
      } finally {
        this.loading = false
      }
    },

    async update(id: number, pelatihan: Omit<Pelatihan, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const index = this.items.findIndex((p) => p.id === id)
        if (index === -1) {
          throw new Error('Pelatihan tidak ditemukan')
        }
        const updated: Pelatihan = { id, ...pelatihan }
        this.items[index] = updated
        return updated
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate pelatihan'
        throw err
      } finally {
        this.loading = false
      }
    },

    async ambil(id: number) {
      this.loading = true
      this.error = ''
      try {
        const pelatihan = this.items.find((p) => p.id === id)
        if (!pelatihan) {
          throw new Error('Pelatihan tidak ditemukan')
        }
        return pelatihan
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail pelatihan'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
