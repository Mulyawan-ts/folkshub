<script setup lang="ts">
import { ref, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useAbsensiStore, useKaryawanStore } from "../stores/index";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppButton from "../components/AppButton.vue";
import AppBadge from "../components/AppBadge.vue";
import { formatTanggal } from "../utils/tanggal";

const absensiStore = useAbsensiStore();
const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(absensiStore);

const karyawanId = ref("");
const tanggal = ref("");
const checkIn = ref("");
const checkOut = ref("");
const status = ref("hadir");
const keterangan = ref("");

const statusOptions = [
  { value: "hadir", label: "Hadir" },
  { value: "izin", label: "Izin" },
  { value: "sakit", label: "Sakit" },
  { value: "alfa", label: "Alfa" },
  { value: "lembur", label: "Lembur" },
];

const statusVariant = (s: string) => {
  switch (s) {
    case "hadir":
      return "success";
    case "izin":
      return "info";
    case "sakit":
      return "warning";
    case "alfa":
      return "error";
    case "lembur":
      return "neutral";
    default:
      return "neutral";
  }
};

onMounted(() => {
  karyawanStore.fetchAll();
  absensiStore.fetchAll();
});

async function handleSimpan() {
  if (!karyawanId.value.trim() || !tanggal.value.trim()) {
    toastStore.error("Karyawan dan tanggal wajib diisi");
    return;
  }
  try {
    await absensiStore.simpan({
      karyawan_id: Number(karyawanId.value),
      tanggal: tanggal.value,
      check_in: checkIn.value || undefined,
      check_out: checkOut.value || undefined,
      status: status.value as any,
      keterangan: keterangan.value,
    });
    toastStore.success("Absensi berhasil disimpan");
    karyawanId.value = "";
    tanggal.value = "";
    checkIn.value = "";
    checkOut.value = "";
    keterangan.value = "";
    status.value = "hadir";
    absensiStore.fetchAll();
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm("Yakin ingin menghapus absensi ini?")) return;
  try {
    await absensiStore.hapus(id);
    toastStore.success("Absensi berhasil dihapus");
    absensiStore.fetchAll();
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
        <h1 class="text-2xl font-bold">Kelola Absensi</h1>
        <p class="text-text-muted">Data kehadiran karyawan</p>
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
          <AppInput
            v-model="karyawanId"
            type="number"
            placeholder="ID karyawan"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Tanggal</span>
          </label>
          <AppInput v-model="tanggal" type="date" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Check In</span>
          </label>
          <AppInput v-model="checkIn" type="time" />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Check Out</span>
          </label>
          <AppInput v-model="checkOut" type="time" />
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

      <div class="form-control w-full mt-4">
        <label class="label">
          <span class="label-text">Keterangan</span>
        </label>
        <AppInput v-model="keterangan" placeholder="Catatan tambahan" />
      </div>
    </form>

    <div class="card p-4 bg-base-100 shadow-sm overflow-x-auto">
      <table class="table w-full text-left">
        <thead>
          <tr>
            <th>Karyawan ID</th>
            <th>Tanggal</th>
            <th>Check In</th>
            <th>Check Out</th>
            <th>Status</th>
            <th class="text-right">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="a in items" :key="a.id" class="border-b">
            <td>{{ a.karyawan_id }}</td>
            <td>{{ formatTanggal(a.tanggal) }}</td>
            <td>{{ a.check_in || "-" }}</td>
            <td>{{ a.check_out || "-" }}</td>
            <td>
              <AppBadge :variant="statusVariant(a.status)" size="sm">
                {{ a.status }}
              </AppBadge>
            </td>
            <td class="text-right">
              <AppButton
                variant="error"
                size="sm"
                outline
                @click="handleHapus(a.id)"
              >
                Hapus
              </AppButton>
            </td>
          </tr>
          <tr v-if="items.length === 0">
            <td colspan="6" class="text-center text-text-muted py-4">
              Belum ada data absensi
            </td>
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
