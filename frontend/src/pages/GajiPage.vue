<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useGajiStore, useKaryawanStore } from "../stores/index";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppAlert from "../components/AppAlert.vue";
import AppButton from "../components/AppButton.vue";
import AppCard from "../components/AppCard.vue";
import AppBadge from "../components/AppBadge.vue";
import TabelData from "../components/TabelData.vue";
import TabelKosong from "../components/TabelKosong.vue";
import { formatBulan } from "../utils/tanggal";
import { cocok } from "../utils/cari";
import { statusVariant, GAJI_STATUS } from "../utils/status";

const gajiStore = useGajiStore();
const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(gajiStore);

const karyawanId = ref("");
const bulan = ref("");
const gajiPokok = ref("");
const tunjangan = ref("");
const potongan = ref("");
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((g) =>
    cocok(
      cari.value,
      g.karyawan_id,
      g.bulan,
      formatBulan(g.bulan),
      g.status,
      g.keterangan,
    ),
  ),
);

onMounted(() => {
  karyawanStore.fetchAll();
  gajiStore.fetchAll();
});

async function handleSimpan() {
  if (
    !karyawanId.value.trim() ||
    !bulan.value.trim() ||
    !gajiPokok.value.trim()
  ) {
    toastStore.error("Karyawan, bulan, dan gaji pokok wajib diisi");
    return;
  }
  try {
    await gajiStore.simpan({
      karyawan_id: Number(karyawanId.value),
      bulan: bulan.value,
      gaji_pokok: Number(gajiPokok.value),
      tunjangan: Number(tunjangan.value) || 0,
      potongan: Number(potongan.value) || 0,
      status: "draft",
    });
    toastStore.success("Data gaji berhasil disimpan");
    karyawanId.value = "";
    bulan.value = "";
    gajiPokok.value = "";
    tunjangan.value = "";
    potongan.value = "";
    gajiStore.fetchAll();
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm("Yakin ingin menghapus data gaji ini?")) return;
  try {
    await gajiStore.hapus(id);
    toastStore.success("Gaji berhasil dihapus");
    gajiStore.fetchAll();
  } catch {
    // error already set in store
  }
}
</script>

<template>
  <div class="max-w-5xl mx-auto space-y-6">
    <div
      class="flex flex-col md:flex-row md:items-center justify-between gap-4"
    >
      <div>
        <h1 class="text-2xl font-bold">Kelola Gaji</h1>
        <p class="text-base-content/60">Data slip gaji karyawan</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="gajiStore.bersihkanError"
    />

    <AppCard as="form" padding="6" @submit.prevent="handleSimpan">
      <div class="grid grid-cols-1 md:grid-cols-6 gap-4 items-end">
        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Karyawan ID</span>
          </label>
          <AppInput
            v-model="karyawanId"
            type="number"
            placeholder="ID karyawan"
            required
          />
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
          <AppInput
            v-model="gajiPokok"
            type="number"
            placeholder="0"
            required
          />
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
    </AppCard>

    <TabelData
      v-model:cari="cari"
      placeholder="Cari karyawan, bulan, status, atau keterangan"
    >
      <template #head>
        <th>Karyawan ID</th>
        <th>Bulan</th>
        <th>Gaji Pokok</th>
        <th>Tunjangan</th>
        <th>Potongan</th>
        <th>Total</th>
        <th>Status</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr v-for="g in filteredItems" :key="g.id" class="border-b">
        <td>{{ g.karyawan_id }}</td>
        <td>{{ formatBulan(g.bulan) }}</td>
        <td>{{ g.gaji_pokok }}</td>
        <td>{{ g.tunjangan }}</td>
        <td>{{ g.potongan }}</td>
        <td class="font-bold">{{ g.total }}</td>
        <td>
          <AppBadge :variant="statusVariant(g.status, GAJI_STATUS)" size="sm">
            {{ g.status }}
          </AppBadge>
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="handleHapus(g.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="8"
        kosong="Belum ada data gaji"
        kosong-cari="Tidak ada gaji yang cocok"
        :mencari="!!cari"
      />
    </TabelData>
  </div>
</template>
