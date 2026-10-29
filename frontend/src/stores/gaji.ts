import { defineStore } from 'pinia'

export type StatusGaji = 'draft' | 'diproses' | 'dibayar' | 'ditolak'

export interface Gaji {
  id: number
  karyawanId: number
  bulan: string
  gajiPokok: number
  tunjangan: number
  potongan: number
  total?: number
  status: StatusGaji
  tanggalBayar?: string
  keterangan?: string
}

let nextGajiId = 1

export const useGajiStore = defineStore('gaji', {
  state: () => ({
    items: [] as Gaji[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((g) => g.id === id),
    byKaryawan: (state) => (karyawanId: number) =>
      state.items.filter((g) => g.karyawanId === karyawanId),
    byBulan: (state) => (bulan: string) =>
      state.items.filter((g) => g.bulan === bulan),
    byStatus: (state) => (status: StatusGaji) =>
      state.items.filter((g) => g.status === status),
    totalGaji: (state) =>
      state.items.reduce((acc, g) => acc + (g.total ?? 0), 0),
    totalDibayar: (state) =>
      state.items
        .filter((g) => g.status === 'dibayar')
        .reduce((acc, g) => acc + (g.total ?? 0), 0),
    belumDibayarCount: (state) =>
      state.items.filter((g) => g.status !== 'dibayar').length
  },

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        await new Promise((r) => setTimeout(r, 300))
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat data gaji'
      } finally {
        this.loading = false
      }
    },

    async simpan(gaji: Omit<Gaji, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const created: Gaji = {
          id: nextGajiId++,
          ...gaji,
          total: gaji.gajiPokok + gaji.tunjangan - gaji.potongan
        }
        this.items.push(created)
        return created
      } catch (err: any) {
        this.error = err?.message || 'Gagal menyimpan gaji'
        throw err
      } finally {
        this.loading = false
      }
    },

    async hapus(id: number) {
      this.loading = true
      this.error = ''
      try {
        this.items = this.items.filter((g) => g.id !== id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal menghapus gaji'
        throw err
      } finally {
        this.loading = false
      }
    },

    async update(id: number, gaji: Omit<Gaji, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const index = this.items.findIndex((g) => g.id === id)
        if (index === -1) {
          throw new Error('Gaji tidak ditemukan')
        }
        const updated: Gaji = {
          id,
          ...gaji,
          total: gaji.gajiPokok + gaji.tunjangan - gaji.potongan
        }
        this.items[index] = updated
        return updated
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate gaji'
        throw err
      } finally {
        this.loading = false
      }
    },

    async ambil(id: number) {
      this.loading = true
      this.error = ''
      try {
        const gaji = this.items.find((g) => g.id === id)
        if (!gaji) {
          throw new Error('Gaji tidak ditemukan')
        }
        return gaji
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail gaji'
        throw err
      } finally {
        this.loading = false
      }
    },

    async setStatus(id: number, status: StatusGaji, tanggalBayar?: string) {
      this.loading = true
      this.error = ''
      try {
        const index = this.items.findIndex((g) => g.id === id)
        if (index === -1) {
          throw new Error('Gaji tidak ditemukan')
        }
        this.items[index].status = status
        if (tanggalBayar !== undefined) {
          this.items[index].tanggalBayar = tanggalBayar
        }
        return this.items[index]
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate status gaji'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
