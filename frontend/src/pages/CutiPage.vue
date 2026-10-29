<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import {
  useCutiStore,
  useKaryawanStore
} from '../stores/index'
import { useToastStore } from '../stores/toast'
import AppInput from '../components/AppInput.vue'
import AppButton from '../components/AppButton.vue'
import AppBadge from '../components/AppBadge.vue'

const cutiStore = useCutiStore()
const karyawanStore = useKaryawanStore()
const toastStore = useToastStore()

const { items, loading, error } = storeToRefs(cutiStore)

const karyawanId = ref('')
const jenis = ref('tahunan')
const tanggalMulai = ref('')
const tanggalSelesai = ref('')
const jumlahHari = ref('')
const keterangan = ref('')

const jenisOptions = [
  { value: 'tahunan', label: 'Tahunan' },
  { value: 'sakit', label: 'Sakit' },
  { value: 'izin', label: 'Izin' },
  { value: 'melahirkan', label: 'Melahirkan' },
  { value: 'lainnya', label: 'Lainnya' }
]

const statusVariant = (s: string) => {
  switch (s) {
    case 'submitted': return 'warning'
    case 'approved': return 'success'
    case 'rejected': return 'error'
    case 'cancelled': return 'neutral'
    case 'draft': return 'neutral'
    default: return 'neutral'
  }
}

onMounted(() => {
  karyawanStore.fetchAll()
  cutiStore.fetchAll()
})

async function handleSimpan() {
  if (!karyawanId.value.trim() || !tanggalMulai.value.trim() || !tanggalSelesai.value.trim()) {
    toastStore.error('Karyawan, tanggal mulai, dan tanggal selesai wajib diisi')
    return
  }
  try {
    await cutiStore.simpan({
      karyawanId: Number(karyawanId.value),
      jenis: jenis.value as any,
      tanggalMulai: tanggalMulai.value,
      tanggalSelesai: tanggalSelesai.value,
      jumlahHari: Number(jumlahHari.value) || 0,
      keterangan: keterangan.value,
      status: 'draft'
    })
    toastStore.success('Cuti berhasil diajukan')
    karyawanId.value = ''
    tanggalMulai.value = ''
    tanggalSelesai.value = ''
    jumlahHari.value = ''
    keterangan.value = ''
    cutiStore.fetchAll()
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm('Yakin ingin menghapus pengajuan cuti ini?')) return
  try {
    await cutiStore.hapus(id)
    toastStore.success('Cuti berhasil dihapus')
    cutiStore.fetchAll()
  } catch {
    // error already set in store
  }
}
</script>

<template>
  <div class="max-w-5xl mx-auto space-y-6">
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-bold">Kelola Cuti</h1>
        <p class="text-text-muted">Pengajuan cuti karyawan</p>
      </div>
    </div>

    <div v-if="error" class="alert alert-error">
      <span>{{ error }}</span>
    </div>

    <form @submit.prevent="handleSimpan" class="card p-6 bg-base-100 shadow-sm">
      <div class="grid grid-cols-1 md:grid-cols-6 gap-4 items-end">
        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Karyawan ID</span>
          </label>
          <AppInput v-model="karyawanId" type="number" placeholder="ID karyawan" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Jenis Cuti</span>
          </label>
          <select v-model="jenis" class="select select-bordered w-full">
            <option v-for="opt in jenisOptions" :key="opt.value" :value="opt.value">
              {{ opt.label }}
            </option>
          </select>
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Tanggal Mulai</span>
          </label>
          <AppInput v-model="tanggalMulai" type="date" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Tanggal Selesai</span>
          </label>
          <AppInput v-model="tanggalSelesai" type="date" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Jumlah Hari</span>
          </label>
          <AppInput v-model="jumlahHari" type="number" placeholder="0" />
        </div>

        <div class="w-full">
          <AppButton type="submit" variant="primary" :loading="loading">
            Ajukan
          </AppButton>
        </div>
      </div>

      <div class="form-control w-full mt-4">
        <label class="label">
          <span class="label-text">Keterangan</span>
        </label>
        <AppInput v-model="keterangan" placeholder="Alasan pengajuan cuti" />
      </div>
    </form>

    <div class="card p-4 bg-base-100 shadow-sm overflow-x-auto">
      <table class="table w-full text-left">
        <thead>
          <tr>
            <th>Karyawan ID</th>
            <th>Jenis</th>
            <th>Tanggal Mulai</th>
            <th>Tanggal Selesai</th>
            <th>Jumlah Hari</th>
            <th>Status</th>
            <th class="text-right">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="c in items" :key="c.id" class="border-b">
            <td>{{ c.karyawanId }}</td>
            <td>{{ c.jenis }}</td>
            <td>{{ c.tanggalMulai }}</td>
            <td>{{ c.tanggalSelesai }}</td>
            <td>{{ c.jumlahHari }}</td>
            <td>
              <AppBadge :variant="statusVariant(c.status)" size="sm">
                {{ c.status }}
              </AppBadge>
            </td>
            <td class="text-right">
              <AppButton variant="error" size="sm" outline @click="handleHapus(c.id)">
                Hapus
              </AppButton>
            </td>
          </tr>
          <tr v-if="items.length === 0">
            <td colspan="7" class="text-center text-text-muted py-4">Belum ada data cuti</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>

  <style scoped>
  .card {
    @apply border;
  }
  </style>
</template>
