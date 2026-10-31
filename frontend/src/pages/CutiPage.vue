<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useCutiStore, useKaryawanStore } from "../stores/index";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppAlert from "../components/AppAlert.vue";
import AppConfirm from "../components/AppConfirm.vue";
import AppButton from "../components/AppButton.vue";
import AppCard from "../components/AppCard.vue";
import AppBadge from "../components/AppBadge.vue";
import TabelData from "../components/TabelData.vue";
import TabelKosong from "../components/TabelKosong.vue";
import { formatTanggal } from "../utils/tanggal";
import { cocok } from "../utils/cari";
import { statusVariant, CUTI_STATUS } from "../utils/status";

const cutiStore = useCutiStore();
const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(cutiStore);

const karyawanId = ref("");
const jenis = ref("tahunan");
const tanggalMulai = ref("");
const tanggalSelesai = ref("");
const jumlahHari = ref("");
const keterangan = ref("");
const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((c) =>
    cocok(
      cari.value,
      c.karyawan_id,
      c.jenis,
      c.status,
      c.tanggal_mulai,
      c.tanggal_selesai,
      c.keterangan,
    ),
  ),
);

const jenisOptions = [
  { value: "tahunan", label: "Tahunan" },
  { value: "sakit", label: "Sakit" },
  { value: "izin", label: "Izin" },
  { value: "melahirkan", label: "Melahirkan" },
  { value: "lainnya", label: "Lainnya" },
];

onMounted(() => {
  karyawanStore.fetchAll();
  cutiStore.fetchAll();
});

async function handleSimpan() {
  if (
    !karyawanId.value.trim() ||
    !tanggalMulai.value.trim() ||
    !tanggalSelesai.value.trim()
  ) {
    toastStore.error(
      "Karyawan, tanggal mulai, dan tanggal selesai wajib diisi",
    );
    return;
  }
  try {
    await cutiStore.simpan({
      karyawan_id: Number(karyawanId.value),
      jenis: jenis.value as any,
      tanggal_mulai: tanggalMulai.value,
      tanggal_selesai: tanggalSelesai.value,
      jumlah_hari: Number(jumlahHari.value) || 0,
      keterangan: keterangan.value,
      status: "draft",
    });
    toastStore.success("Cuti berhasil diajukan");
    karyawanId.value = "";
    tanggalMulai.value = "";
    tanggalSelesai.value = "";
    jumlahHari.value = "";
    keterangan.value = "";
    cutiStore.fetchAll();
  } catch {
    // error already set in store
  }
}

function mintaHapus(id: number) {
  hapusId.value = id;
  konfirmasiHapus.value = true;
}

async function jalankanHapus() {
  const id = hapusId.value;
  konfirmasiHapus.value = false;
  hapusId.value = null;
  if (id === null) return;

  try {
    await cutiStore.hapus(id);
    toastStore.success("Cuti berhasil dihapus");
    cutiStore.fetchAll();
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
        <h1 class="text-2xl font-bold">Kelola Cuti</h1>
        <p class="text-base-content/60">Pengajuan cuti karyawan</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="cutiStore.bersihkanError"
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
            <span class="label-text">Jenis Cuti</span>
          </label>
          <select v-model="jenis" class="select select-bordered w-full">
            <option
              v-for="opt in jenisOptions"
              :key="opt.value"
              :value="opt.value"
            >
              {{ opt.label }}
            </option>
          </select>
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Tanggal Mulai</span>
          </label>
          <AppInput v-model="tanggalMulai" type="date" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Tanggal Selesai</span>
          </label>
          <AppInput v-model="tanggalSelesai" type="date" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Jumlah Hari</span>
          </label>
          <AppInput v-model="jumlahHari" type="number" placeholder="0" />
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
        <AppInput v-model="keterangan" placeholder="Alasan pengajuan cuti" />
      </div>
    </AppCard>

    <TabelData
      v-model:cari="cari"
      placeholder="Cari karyawan, jenis, tanggal, status, atau keterangan"
    >
      <template #head>
        <th>Karyawan ID</th>
        <th>Jenis</th>
        <th>Tanggal Mulai</th>
        <th>Tanggal Selesai</th>
        <th>Jumlah Hari</th>
        <th>Status</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr v-for="c in filteredItems" :key="c.id" class="border-b">
        <td>{{ c.karyawan_id }}</td>
        <td>{{ c.jenis }}</td>
        <td>{{ formatTanggal(c.tanggal_mulai) }}</td>
        <td>{{ formatTanggal(c.tanggal_selesai) }}</td>
        <td>{{ c.jumlah_hari }}</td>
        <td>
          <AppBadge :variant="statusVariant(c.status, CUTI_STATUS)" size="sm">
            {{ c.status }}
          </AppBadge>
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(c.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="7"
        kosong="Belum ada data cuti"
        kosong-cari="Tidak ada cuti yang cocok"
        :mencari="!!cari"
      />
    </TabelData>

    <AppConfirm
      :terbuka="konfirmasiHapus"
      judul="Hapus cuti"
      pesan="Yakin ingin menghapus pengajuan cuti ini?"
      confirm-label="Hapus"
      variant="error"
      @konfirmasi="jalankanHapus"
      @batal="konfirmasiHapus = false"
    />
  </div>
</template>
