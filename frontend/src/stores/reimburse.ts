import { defineStore } from 'pinia'
import { SimpanReimburse, SemuaReimburse, HapusReimburse, UpdateReimburse, AmbilReimburse } from '../../wailsjs/go/main/App'

export type StatusReimburse = 'draft' | 'submitted' | 'approved' | 'rejected' | 'paid'

export interface Reimburse {
  id: number
  karyawanId: number
  tanggal: string
  jenis: string
  jumlah: number
  keterangan: string
  bukti: string
  status: StatusReimburse
  catatanApproval?: string
}

export const useReimburseStore = defineStore('reimburse', {
  state: () => ({
    items: [] as Reimburse[],
    loading: false,
    error: '' as string
  }),

  getters: {
    itemById: (state) => (id: number) => state.items.find((r) => r.id === id),
    byKaryawan: (state) => (karyawanId: number) =>
      state.items.filter((r) => r.karyawanId === karyawanId),
    byStatus: (state) => (status: StatusReimburse) =>
      state.items.filter((r) => r.status === status),
    totalDiajujkan: (state) =>
      state.items.reduce((acc, r) => acc + r.jumlah, 0),
    totalDibayar: (state) =>
      state.items
        .filter((r) => r.status === 'paid')
        .reduce((acc, r) => acc + r.jumlah, 0)
  },

  actions: {
    async fetchAll() {
      this.loading = true
      this.error = ''
      try {
        this.items = await SemuaReimburse()
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat data reimburse'
      } finally {
        this.loading = false
      }
    },

    async simpan(reimburse: Omit<Reimburse, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const created = await SimpanReimburse(
          reimburse.karyawanId,
          reimburse.tanggal,
          reimburse.jenis,
          reimburse.keterangan,
          reimburse.bukti,
          reimburse.status,
          reimburse.jumlah
        )
        this.items.push(created)
        return created
      } catch (err: any) {
        this.error = err?.message || 'Gagal menyimpan reimburse'
        throw err
      } finally {
        this.loading = false
      }
    },

    async hapus(id: number) {
      this.loading = true
      this.error = ''
      try {
        await HapusReimburse(id)
        this.items = this.items.filter((r) => r.id !== id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal menghapus reimburse'
        throw err
      } finally {
        this.loading = false
      }
    },

    async update(id: number, reimburse: Omit<Reimburse, 'id'>) {
      this.loading = true
      this.error = ''
      try {
        const updated = await UpdateReimburse(
          id,
          reimburse.karyawanId,
          reimburse.tanggal,
          reimburse.jenis,
          reimburse.keterangan,
          reimburse.bukti,
          reimburse.status,
          reimburse.jumlah
        )
        const index = this.items.findIndex((r) => r.id === id)
        if (index !== -1) {
          this.items[index] = updated
        }
        return updated
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate reimburse'
        throw err
      } finally {
        this.loading = false
      }
    },

    async ambil(id: number) {
      this.loading = true
      this.error = ''
      try {
        return await AmbilReimburse(id)
      } catch (err: any) {
        this.error = err?.message || 'Gagal memuat detail reimburse'
        throw err
      } finally {
        this.loading = false
      }
    },

    async setStatus(id: number, status: StatusReimburse, catatan?: string) {
      this.loading = true
      this.error = ''
      try {
        const updated = await UpdateReimburse(
          id,
          0,
          '',
          '',
          catatan || '',
          '',
          status,
          0
        )
        const index = this.items.findIndex((r) => r.id === id)
        if (index !== -1) {
          this.items[index] = updated
        }
        return this.items[index]
      } catch (err: any) {
        this.error = err?.message || 'Gagal mengupdate status reimburse'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
