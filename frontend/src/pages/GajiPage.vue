<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import {
  useGajiStore,
  useKaryawanStore
} from '../stores/index'
import { useToastStore } from '../stores/toast'
import AppInput from '../components/AppInput.vue'
import AppButton from '../components/AppButton.vue'
import AppBadge from '../components/AppBadge.vue'
import { formatBulan } from '../utils/tanggal'

const gajiStore = useGajiStore()
const karyawanStore = useKaryawanStore()
const toastStore = useToastStore()

const { items, loading, error } = storeToRefs(gajiStore)

const karyawanId = ref('')
const bulan = ref('')
const gajiPokok = ref('')
const tunjangan = ref('')
const potongan = ref('')

const statusVariant = (s: string) => {
  switch (s) {
    case 'diproses': return 'warning'
    case 'dibayar': return 'success'
    case 'ditolak': return 'error'
    case 'draft': return 'neutral'
    default: return 'neutral'
  }
}

onMounted(() => {
  karyawanStore.fetchAll()
  gajiStore.fetchAll()
})

async function handleSimpan() {
  if (!karyawanId.value.trim() || !bulan.value.trim() || !gajiPokok.value.trim()) {
    toastStore.error('Karyawan, bulan, dan gaji pokok wajib diisi')
    return
  }
  try {
    await gajiStore.simpan({
      karyawan_id: Number(karyawanId.value),
      bulan: bulan.value,
      gaji_pokok: Number(gajiPokok.value),
      tunjangan: Number(tunjangan.value) || 0,
      potongan: Number(potongan.value) || 0,
      status: 'draft'
    })
    toastStore.success('Data gaji berhasil disimpan')
    karyawanId.value = ''
    bulan.value = ''
    gajiPokok.value = ''
    tunjangan.value = ''
    potongan.value = ''
    gajiStore.fetchAll()
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm('Yakin ingin menghapus data gaji ini?')) return
  try {
    await gajiStore.hapus(id)
    toastStore.success('Gaji berhasil dihapus')
    gajiStore.fetchAll()
  } catch {
    // error already set in store
  }
}
</script>

<template>
  <div class="max-w-5xl mx-auto space-y-6">
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-bold">Kelola Gaji</h1>
        <p class="text-text-muted">Data slip gaji karyawan</p>
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
            <span class="label-text">Bulan</span>
          </label>
          <AppInput v-model="bulan" type="month" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Gaji Pokok</span>
          </label>
          <AppInput v-model="gajiPokok" type="number" placeholder="0" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Tunjangan</span>
          </label>
          <AppInput v-model="tunjangan" type="number" placeholder="0" />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Potongan</span>
          </label>
          <AppInput v-model="potongan" type="number" placeholder="0" />
        </div>

        <div class="w-full">
          <AppButton type="submit" variant="primary" :loading="loading">
            Simpan
          </AppButton>
        </div>
      </div>
    </form>

    <div class="card p-4 bg-base-100 shadow-sm overflow-x-auto">
      <table class="table w-full text-left">
        <thead>
          <tr>
            <th>Karyawan ID</th>
            <th>Bulan</th>
            <th>Gaji Pokok</th>
            <th>Tunjangan</th>
            <th>Potongan</th>
            <th>Total</th>
            <th>Status</th>
            <th class="text-right">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="g in items" :key="g.id" class="border-b">
            <td>{{ g.karyawan_id }}</td>
            <td>{{ formatBulan(g.bulan) }}</td>
            <td>{{ g.gaji_pokok }}</td>
            <td>{{ g.tunjangan }}</td>
            <td>{{ g.potongan }}</td>
            <td class="font-bold">{{ g.total }}</td>
            <td>
              <AppBadge :variant="statusVariant(g.status)" size="sm">
                {{ g.status }}
              </AppBadge>
            </td>
            <td class="text-right">
              <AppButton variant="error" size="sm" outline @click="handleHapus(g.id)">
                Hapus
              </AppButton>
            </td>
          </tr>
          <tr v-if="items.length === 0">
            <td colspan="8" class="text-center text-text-muted py-4">Belum ada data gaji</td>
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
