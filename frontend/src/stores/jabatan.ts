import { defineStore } from 'pinia'
import { SimpanJabatan, SemuaJabatan, HapusJabatan, UpdateJabatan, AmbilJabatan } from '../../wailsjs/go/main/App'

export interface Jabatan {
  id: number
  kode: string
  nama: string
  departemen: string
  gajiPokok: number
  tunjangan: number
  deskripsi: string
}

export const useJabatanStore = defineStore('jabatan', {
  state: () => ({
    items: [] as Jabatan[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((j) => j.id === id),
    byDepartemen: (state) => (departemen: string) =>
      state.items.filter((j) => j.departemen === departemen)
  },

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        this.items = await SemuaJabatan()
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat data jabatan'
      } finally {
        this.loading = false
      }
    },

    async simpan(jabatan: Omit<Jabatan, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const created = await SimpanJabatan(
          jabatan.kode,
          jabatan.nama,
          jabatan.departemen,
          jabatan.deskripsi,
          jabatan.gajiPokok,
          jabatan.tunjangan
        )
        this.items.push(created)
        return created
      } catch (err: any) {
        this.error = err?.message || 'Gagal menyimpan jabatan'
        throw err
      } finally {
        this.loading = false
      }
    },

    async hapus(id: number) {
      this.loading = true
      this.error = ''
      try {
        await HapusJabatan(id)
        this.items = this.items.filter((j) => j.id !== id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal menghapus jabatan'
        throw err
      } finally {
        this.loading = false
      }
    },

    async update(id: number, jabatan: Omit<Jabatan, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const updated = await UpdateJabatan(
          id,
          jabatan.kode,
          jabatan.nama,
          jabatan.departemen,
          jabatan.deskripsi,
          jabatan.gajiPokok,
          jabatan.tunjangan
        )
        const index = this.items.findIndex((j) => j.id === id)
        if (index !== -1) {
          this.items[index] = updated
        }
        return updated
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate jabatan'
        throw err
      } finally {
        this.loading = false
      }
    },

    async ambil(id: number) {
      this.loading = true
      this.error = ''
      try {
        return await AmbilJabatan(id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail jabatan'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
