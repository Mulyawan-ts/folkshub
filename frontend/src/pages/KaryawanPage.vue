<script setup lang="ts">
import { ref, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useKaryawanStore } from "../stores/karyawan";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppButton from "../components/AppButton.vue";
import AppBadge from "../components/AppBadge.vue";

const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(karyawanStore);

const nama = ref("");
const email = ref("");
const status = ref("tetap");

const statusOptions = [
  { value: "tetap", label: "Tetap" },
  { value: "outsourcing", label: "Outsourcing" },
  { value: "magang", label: "Magang" },
];

const itemStatusVariant = (s: string) => {
  switch (s) {
    case "tetap":
      return "success";
    case "outsourcing":
      return "warning";
    case "magang":
      return "info";
    default:
      return "neutral";
  }
};

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

    <div class="card p-4 bg-base-100 shadow-sm overflow-x-auto">
      <table class="table w-full text-left">
        <thead>
          <tr>
            <th>Nama</th>
            <th>Email</th>
            <th>Status</th>
            <th class="text-right">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="k in items" :key="k.id" class="border-b">
            <td>{{ k.nama }}</td>
            <td>{{ k.email }}</td>
            <td>
              <AppBadge :variant="itemStatusVariant(k.status)" size="sm">
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
          <tr v-if="items.length === 0">
            <td colspan="4" class="text-center text-text-muted py-4">
              Belum ada data karyawan
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.card {
  @apply border;
}
</style>
