import { defineStore } from 'pinia'
import { SimpanKaryawan, SemuaKaryawan, HapusKaryawan } from '../../wailsjs/go/main/App'

export interface Karyawan {
  id: number
  nama: string
  email: string
  status: string
}

export const useKaryawanStore = defineStore('karyawan', {
  state: () => ({
    items: [] as Karyawan[],
    loading: false,
    error: '' as string
  }),

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        this.items = await SemuaKaryawan()
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat data karyawan'
      } finally {
        this.loading = false
      }
    },

    async simpan(karyawan: Omit<Karyawan, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const created = await SimpanKaryawan(karyawan.nama, karyawan.email, karyawan.status)
        this.items.push(created)
        return created
      } catch (err: any) {
        this.error = err?.message || 'Gagal menyimpan karyawan'
        throw err
      } finally {
        this.loading = false
      }
    },

    async hapus(id: number) {
      this.loading = true
      this.error = ''
      try {
        await HapusKaryawan(id)
        this.items = this.items.filter((k) => k.id !== id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal menghapus karyawan'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
