<script setup lang="ts">
import { onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useDashboardStore } from "@/stores/dashboard";
import AppCard from "@/components/AppCard.vue";
import AppAlert from "@/components/AppAlert.vue";

const dashboardStore = useDashboardStore();
const { data, error, loading } = storeToRefs(dashboardStore);

onMounted(() => {
  dashboardStore.fetch();
});
</script>

<template>
  <div class="max-w-6xl mx-auto space-y-6">
    <div>
      <h1 class="text-2xl font-bold">Dashboard</h1>
      <p class="text-base-content/60">Ringkasan data HRIS</p>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="dashboardStore.bersihkanError"
    />

    <div v-if="loading" class="text-center py-8">
      <span class="loading loading-spinner loading-lg"></span>
    </div>

    <div v-else-if="data" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <AppCard padding="4">
        <p class="text-sm text-base-content/60">Total Karyawan</p>
        <p class="text-3xl font-bold">{{ data.totalKaryawan }}</p>
        <p class="text-xs text-base-content/50 mt-1">
          Tetap {{ data.karyawanTetap }} · Outsourcing {{ data.karyawanOutsourcing }} · Magang {{ data.karyawanMagang }}
        </p>
      </AppCard>

      <AppCard padding="4">
        <p class="text-sm text-base-content/60">Absensi</p>
        <p class="text-3xl font-bold">{{ data.totalAbsensi }}</p>
      </AppCard>

      <AppCard padding="4">
        <p class="text-sm text-base-content/60">Cuti</p>
        <p class="text-3xl font-bold">{{ data.totalCuti }}</p>
        <p class="text-xs text-base-content/50 mt-1">Pending {{ data.cutiPending }}</p>
      </AppCard>

      <AppCard padding="4">
        <p class="text-sm text-base-content/60">Reimburse</p>
        <p class="text-3xl font-bold">{{ data.totalReimburse }}</p>
        <p class="text-xs text-base-content/50 mt-1">Pending {{ data.reimbursePending }}</p>
      </AppCard>

      <AppCard padding="4">
        <p class="text-sm text-base-content/60">Gaji (record)</p>
        <p class="text-3xl font-bold">{{ data.totalGaji }}</p>
      </AppCard>

      <AppCard padding="4">
        <p class="text-sm text-base-content/60">Jabatan</p>
        <p class="text-3xl font-bold">{{ data.totalJabatan }}</p>
      </AppCard>

      <AppCard padding="4">
        <p class="text-sm text-base-content/60">Departemen</p>
        <p class="text-3xl font-bold">{{ data.totalDepartemen }}</p>
      </AppCard>

      <AppCard padding="4">
        <p class="text-sm text-base-content/60">Pelatihan</p>
        <p class="text-3xl font-bold">{{ data.totalPelatihan }}</p>
      </AppCard>
    </div>
  </div>
</template>
