<script setup lang="ts">
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import AppToast from './components/AppToast.vue'
import { useToastStore } from './stores/toast'
import KaryawanPage from './pages/KaryawanPage.vue'
import AbsensiPage from './pages/AbsensiPage.vue'
import CutiPage from './pages/CutiPage.vue'
import ReimbursePage from './pages/ReimbursePage.vue'
import GajiPage from './pages/GajiPage.vue'
import JabatanPage from './pages/JabatanPage.vue'
import PelatihanPage from './pages/PelatihanPage.vue'
import DepartemenPage from './pages/DepartemenPage.vue'

const toastStore = useToastStore()
const { items } = storeToRefs(toastStore)

const pages = [
  { id: 'karyawan', label: 'Karyawan', component: KaryawanPage },
  { id: 'absensi', label: 'Absensi', component: AbsensiPage },
  { id: 'cuti', label: 'Cuti', component: CutiPage },
  { id: 'reimburse', label: 'Reimburse', component: ReimbursePage },
  { id: 'gaji', label: 'Gaji', component: GajiPage },
  { id: 'jabatan', label: 'Jabatan', component: JabatanPage },
  { id: 'departemen', label: 'Departemen', component: DepartemenPage },
  { id: 'pelatihan', label: 'Pelatihan', component: PelatihanPage }
]

const activePage = ref('karyawan')

const PageComponent = pages.find((p) => p.id === activePage.value)?.component || KaryawanPage
</script>

<template>
  <main class="p-8">
    <h1 class="text-2xl font-semibold mb-6">Folkshub</h1>

    <div class="tabs tabs-bordered mb-8">
      <div class="tab" v-for="page in pages" :key="page.id">
        <button
          :class="{ 'tab-active': activePage === page.id }"
          @click="activePage = page.id"
        >
          {{ page.label }}
        </button>
      </div>
    </div>

    <component :is="PageComponent" />
  </main>

  <AppToast :items="items" @dismiss="toastStore.dismiss" />
</template>
