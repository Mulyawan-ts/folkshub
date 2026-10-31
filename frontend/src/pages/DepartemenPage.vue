<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useDepartemenStore } from "../stores/index";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppButton from "../components/AppButton.vue";

const departemenStore = useDepartemenStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(departemenStore);

const kode = ref("");
const nama = ref("");
const kepala = ref("");
const deskripsi = ref("");
const cari = ref("");

const filteredItems = computed(() => {
  const q = cari.value.trim().toLowerCase();
  if (!q) return items.value;
  return items.value.filter(
    (d) =>
      d.kode.toLowerCase().includes(q) ||
      d.nama.toLowerCase().includes(q) ||
      (d.kepala ?? "").toLowerCase().includes(q) ||
      (d.deskripsi ?? "").toLowerCase().includes(q),
  );
});

onMounted(() => {
  departemenStore.fetchAll();
});

async function handleSimpan() {
  if (!kode.value.trim() || !nama.value.trim()) {
    toastStore.error("Kode dan nama departemen wajib diisi");
    return;
  }
  try {
    await departemenStore.simpan({
      kode: kode.value,
      nama: nama.value,
      kepala: kepala.value,
      deskripsi: deskripsi.value,
    });
    toastStore.success("Departemen berhasil disimpan");
    kode.value = "";
    nama.value = "";
    kepala.value = "";
    deskripsi.value = "";
    departemenStore.fetchAll();
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm("Yakin ingin menghapus departemen ini?")) return;
  try {
    await departemenStore.hapus(id);
    toastStore.success("Departemen berhasil dihapus");
    departemenStore.fetchAll();
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
        <h1 class="text-2xl font-bold">Kelola Departemen</h1>
        <p class="text-text-muted">Data departemen organisasi</p>
      </div>
    </div>

    <div v-if="error" class="alert alert-error">
      <span>{{ error }}</span>
    </div>

    <form @submit.prevent="handleSimpan" class="card p-6 bg-base-100 shadow-sm">
      <div class="grid grid-cols-1 md:grid-cols-6 gap-4 items-end">
        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Kode Departemen</span>
          </label>
          <AppInput v-model="kode" placeholder="Contoh: HRD, FIN" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Nama Departemen</span>
          </label>
          <AppInput
            v-model="nama"
            placeholder="Contoh: Human Resources"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Kepala</span>
          </label>
          <AppInput v-model="kepala" placeholder="Nama kepala departemen" />
        </div>

        <div class="w-full col-span-6">
          <AppButton type="submit" variant="primary" :loading="loading">
            Simpan
          </AppButton>
        </div>
      </div>

      <div class="form-control w-full mt-4">
        <label class="label">
          <span class="label-text">Deskripsi</span>
        </label>
        <AppInput v-model="deskripsi" placeholder="Deskripsi departemen" />
      </div>
    </form>

    <div class="card p-4 bg-base-100 shadow-sm overflow-x-auto">
      <AppInput
        v-model="cari"
        placeholder="Cari kode, nama, atau kepala"
        class="mb-4"
      />
      <table class="table w-full text-left">
        <thead>
          <tr>
            <th>Kode</th>
            <th>Nama Departemen</th>
            <th>Kepala</th>
            <th>Deskripsi</th>
            <th class="text-right">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="d in filteredItems" :key="d.id" class="border-b">
            <td>{{ d.kode }}</td>
            <td>{{ d.nama }}</td>
            <td>{{ d.kepala || "-" }}</td>
            <td>{{ d.deskripsi || "-" }}</td>
            <td class="text-right">
              <AppButton
                variant="error"
                size="sm"
                outline
                @click="handleHapus(d.id)"
              >
                Hapus
              </AppButton>
            </td>
          </tr>
          <tr v-if="filteredItems.length === 0">
            <td colspan="5" class="text-center text-text-muted py-4">
              {{
                cari
                  ? "Tidak ada departemen yang cocok"
                  : "Belum ada data departemen"
              }}
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
