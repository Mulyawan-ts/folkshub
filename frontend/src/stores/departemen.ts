import { defineStore } from 'pinia'

export interface Departemen {
  id: number
  kode: string
  nama: string
  kepala: string
  deskripsi: string
}

let nextDepartemenId = 1

export const useDepartemenStore = defineStore('departemen', {
  state: () => ({
    items: [] as Departemen[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((d) => d.id === id),
    byKode: (state) => (kode: string) => state.items.find((d) => d.kode === kode)
  },

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        await new Promise((r) => setTimeout(r, 300))
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat data departemen'
      } finally {
        this.loading = false
      }
    },

    async simpan(departemen: Omit<Departemen, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const created: Departemen = { id: nextDepartemenId++, ...departemen }
        this.items.push(created)
        return created
      } catch (err: any) {
        this.error = err?.message || 'Gagal menyimpan departemen'
        throw err
      } finally {
        this.loading = false
      }
    },

    async hapus(id: number) {
      this.loading = true
      this.error = ''
      try {
        this.items = this.items.filter((d) => d.id !== id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal menghapus departemen'
        throw err
      } finally {
        this.loading = false
      }
    },

    async update(id: number, departemen: Omit<Departemen, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const index = this.items.findIndex((d) => d.id === id)
        if (index === -1) {
          throw new Error('Departemen tidak ditemukan')
        }
        const updated: Departemen = { id, ...departemen }
        this.items[index] = updated
        return updated
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate departemen'
        throw err
      } finally {
        this.loading = false
      }
    },

    async ambil(id: number) {
      this.loading = true
      this.error = ''
      try {
        const departemen = this.items.find((d) => d.id === id)
        if (!departemen) {
          throw new Error('Departemen tidak ditemukan')
        }
        return departemen
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail departemen'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
