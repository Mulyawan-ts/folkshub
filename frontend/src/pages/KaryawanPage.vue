<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useKaryawanStore } from "../stores/karyawan";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppButton from "../components/AppButton.vue";
import AppBadge from "../components/AppBadge.vue";
import TabelData from "../components/TabelData.vue";
import TabelKosong from "../components/TabelKosong.vue";
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

    <div v-if="error" class="alert alert-error">
      <span>{{ error }}</span>
    </div>

    <form @submit.prevent="handleSimpan" class="card p-6 bg-base-100 shadow-sm">
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
    </form>

    <TabelData v-model:cari="cari" placeholder="Cari nama, email, atau status">
      <template #head>
        <th>Nama</th>
        <th>Email</th>
        <th>Status</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr v-for="k in filteredItems" :key="k.id" class="border-b">
        <td>{{ k.nama }}</td>
        <td>{{ k.email }}</td>
        <td>
          <AppBadge
            :variant="statusVariant(k.status, KARYAWAN_STATUS)"
            size="sm"
          >
            {{ k.status }}
          </AppBadge>
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="handleHapus(k.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="4"
        kosong="Belum ada data karyawan"
        kosong-cari="Tidak ada karyawan yang cocok"
        :mencari="!!cari"
      />
    </TabelData>
  </div>
</template>
