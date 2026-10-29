import { defineStore } from 'pinia'
import { SimpanPelatihan, SemuaPelatihan, HapusPelatihan, UpdatePelatihan, AmbilPelatihan } from '../../wailsjs/go/main/App'

export type StatusPelatihan = 'terjadwal' | 'berlangsung' | 'selesai' | 'dibatalkan'

export interface Pelatihan {
  id: number
  karyawan_id: number
  nama: string
  jenis: string
  tanggal_mulai: string
  tanggal_selesai: string
  lokasi: string
  biaya: number
  status: string
  sertifikat: string
  catatan: string
}

export const usePelatihanStore = defineStore('pelatihan', {
  state: () => ({
    items: [] as Pelatihan[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((p) => p.id === id),
    byKaryawan: (state) => (karyawan_id: number) =>
      state.items.filter((p) => p.karyawan_id === karyawan_id),
    byStatus: (state) => (status: StatusPelatihan) =>
      state.items.filter((p) => p.status === status)
  },

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        this.items = await SemuaPelatihan()
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
        const created = await SimpanPelatihan(
          pelatihan.karyawan_id,
          pelatihan.nama,
          pelatihan.jenis,
          pelatihan.tanggal_mulai,
          pelatihan.tanggal_selesai,
          pelatihan.lokasi,
          pelatihan.sertifikat,
          pelatihan.catatan,
          pelatihan.status,
          pelatihan.biaya
        )
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
        await HapusPelatihan(id)
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
        const updated = await UpdatePelatihan(
          id,
          pelatihan.karyawan_id,
          pelatihan.nama,
          pelatihan.jenis,
          pelatihan.tanggal_mulai,
          pelatihan.tanggal_selesai,
          pelatihan.lokasi,
          pelatihan.sertifikat,
          pelatihan.catatan,
          pelatihan.status,
          pelatihan.biaya
        )
        const index = this.items.findIndex((p) => p.id === id)
        if (index !== -1) {
          this.items[index] = updated
        }
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
        return await AmbilPelatihan(id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail pelatihan'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
