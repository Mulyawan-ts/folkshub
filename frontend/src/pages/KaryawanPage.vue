<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useKaryawanStore } from "../stores/karyawan";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppAlert from "../components/AppAlert.vue";
import AppButton from "../components/AppButton.vue";
import AppCard from "../components/AppCard.vue";
import AppCardItem from "../components/AppCardItem.vue";
import DaftarData from "../components/DaftarData.vue";
import { cocok } from "../utils/cari";
import { statusVariant, KARYAWAN_STATUS } from "../utils/status";

const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(karyawanStore);

const nama = ref("");
const email = ref("");
const status = ref("tetap");
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((k) => cocok(cari.value, k.nama, k.email, k.status)),
);

const statusOptions = [
  { value: "tetap", label: "Tetap" },
  { value: "outsourcing", label: "Outsourcing" },
  { value: "magang", label: "Magang" },
];

onMounted(() => {
  karyawanStore.fetchAll();
});

async function handleSimpan() {
  if (!nama.value.trim() || !email.value.trim()) {
    toastStore.error("Nama dan email wajib diisi");
    return;
  }

  try {
    await karyawanStore.simpan({
      nama: nama.value.trim(),
      email: email.value.trim(),
      status: status.value,
    });
    toastStore.success("Karyawan berhasil disimpan");
    nama.value = "";
    email.value = "";
    status.value = "tetap";
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm("Yakin ingin menghapus karyawan ini?")) return;
  try {
    await karyawanStore.hapus(id);
    toastStore.success("Karyawan berhasil dihapus");
  } catch {
    // error already set in store
  }
}
</script>

<template>
  <div class="max-w-4xl mx-auto space-y-6">
    <div
      class="flex flex-col md:flex-row md:items-center justify-between gap-4"
    >
      <div>
        <h1 class="text-2xl font-bold">Kelola Karyawan</h1>
        <p class="text-text-muted">
          Data karyawan (tetap, outsourcing, magang)
        </p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="karyawanStore.bersihkanError"
    />

    <AppCard as="form" padding="6" @submit.prevent="handleSimpan">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4 items-end">
        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Nama</span>
          </label>
          <AppInput v-model="nama" placeholder="Masukkan nama" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Email</span>
          </label>
          <AppInput
            v-model="email"
            type="email"
            placeholder="nama@contoh.com"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Status</span>
          </label>
          <select v-model="status" class="select select-bordered w-full">
            <option
              v-for="opt in statusOptions"
              :key="opt.value"
              :value="opt.value"
            >
              {{ opt.label }}
            </option>
          </select>
        </div>

        <div class="w-full">
          <AppButton type="submit" variant="primary" :loading="loading">
            Simpan
          </AppButton>
        </div>
      </div>
    </AppCard>

    <DaftarData
      v-model:cari="cari"
      placeholder="Cari nama, email, atau status"
      kosong="Belum ada data karyawan"
      kosong-cari="Tidak ada karyawan yang cocok"
      :jumlah="filteredItems.length"
    >
      <AppCardItem
        v-for="k in filteredItems"
        :key="k.id"
        :judul="k.nama"
        :subjudul="k.email"
        :status="k.status"
        :status-variant="statusVariant(k.status, KARYAWAN_STATUS)"
        hapus-label="Hapus"
        @hapus="handleHapus(k.id)"
      />
    </DaftarData>
  </div>
</template>
