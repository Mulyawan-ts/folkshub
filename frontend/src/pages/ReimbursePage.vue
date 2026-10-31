<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useReimburseStore, useKaryawanStore } from "../stores/index";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppButton from "../components/AppButton.vue";
import AppBadge from "../components/AppBadge.vue";
import { formatTanggal } from "../utils/tanggal";

const reimburseStore = useReimburseStore();
const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(reimburseStore);

const karyawanId = ref("");
const tanggal = ref("");
const jenis = ref("");
const jumlah = ref("");
const keterangan = ref("");
const bukti = ref("");
const cari = ref("");

const filteredItems = computed(() => {
  const q = cari.value.trim().toLowerCase();
  if (!q) return items.value;
  return items.value.filter(
    (r) =>
      String(r.karyawan_id).includes(q) ||
      r.tanggal.includes(q) ||
      r.jenis.toLowerCase().includes(q) ||
      r.status.toLowerCase().includes(q) ||
      (r.keterangan ?? "").toLowerCase().includes(q) ||
      (r.bukti ?? "").toLowerCase().includes(q),
  );
});

const statusVariant = (s: string) => {
  switch (s) {
    case "submitted":
      return "warning";
    case "approved":
      return "success";
    case "rejected":
      return "error";
    case "paid":
      return "success";
    case "draft":
      return "neutral";
    default:
      return "neutral";
  }
};

onMounted(() => {
  karyawanStore.fetchAll();
  reimburseStore.fetchAll();
});

async function handleSimpan() {
  if (
    !karyawanId.value.trim() ||
    !tanggal.value.trim() ||
    !jenis.value.trim() ||
    !jumlah.value.trim()
  ) {
    toastStore.error("Karyawan, tanggal, jenis, dan jumlah wajib diisi");
    return;
  }
  try {
    await reimburseStore.simpan({
      karyawan_id: Number(karyawanId.value),
      tanggal: tanggal.value,
      jenis: jenis.value,
      jumlah: Number(jumlah.value),
      keterangan: keterangan.value,
      bukti: bukti.value,
      status: "draft",
    });
    toastStore.success("Reimburse berhasil diajukan");
    karyawanId.value = "";
    tanggal.value = "";
    jenis.value = "";
    jumlah.value = "";
    keterangan.value = "";
    bukti.value = "";
    reimburseStore.fetchAll();
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm("Yakin ingin menghapus reimburse ini?")) return;
  try {
    await reimburseStore.hapus(id);
    toastStore.success("Reimburse berhasil dihapus");
    reimburseStore.fetchAll();
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
        <h1 class="text-2xl font-bold">Kelola Reimburse</h1>
        <p class="text-text-muted">Klaim biaya karyawan</p>
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
            <span class="label-text">Jenis Biaya</span>
          </label>
          <AppInput
            v-model="jenis"
            placeholder="Contoh: Transport, Makan, Akomodasi"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Jumlah</span>
          </label>
          <AppInput v-model="jumlah" type="number" placeholder="0" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Bukti</span>
          </label>
          <AppInput v-model="bukti" placeholder="URL atau nomor bukti" />
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
        <AppInput
          v-model="keterangan"
          placeholder="Alasan pengajuan reimburse"
        />
      </div>
    </form>

    <div class="card p-4 bg-base-100 shadow-sm overflow-x-auto">
      <AppInput
        v-model="cari"
        placeholder="Cari karyawan, jenis, status, atau keterangan"
        class="mb-4"
      />
      <table class="table w-full text-left">
        <thead>
          <tr>
            <th>Karyawan ID</th>
            <th>Tanggal</th>
            <th>Jenis</th>
            <th>Jumlah</th>
            <th>Bukti</th>
            <th>Status</th>
            <th class="text-right">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="r in filteredItems" :key="r.id" class="border-b">
            <td>{{ r.karyawan_id }}</td>
            <td>{{ formatTanggal(r.tanggal) }}</td>
            <td>{{ r.jenis }}</td>
            <td>{{ r.jumlah }}</td>
            <td>{{ r.bukti || "-" }}</td>
            <td>
              <AppBadge :variant="statusVariant(r.status)" size="sm">
                {{ r.status }}
              </AppBadge>
            </td>
            <td class="text-right">
              <AppButton
                variant="error"
                size="sm"
                outline
                @click="handleHapus(r.id)"
              >
                Hapus
              </AppButton>
            </td>
          </tr>
          <tr v-if="filteredItems.length === 0">
            <td colspan="7" class="text-center text-text-muted py-4">
              {{
                cari
                  ? "Tidak ada reimburse yang cocok"
                  : "Belum ada data reimburse"
              }}
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
