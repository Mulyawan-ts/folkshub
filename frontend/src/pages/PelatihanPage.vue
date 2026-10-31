<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { usePelatihanStore } from "../stores/index";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppButton from "../components/AppButton.vue";
import AppCard from "../components/AppCard.vue";
import AppBadge from "../components/AppBadge.vue";
import TabelData from "../components/TabelData.vue";
import TabelKosong from "../components/TabelKosong.vue";
import { formatTanggal } from "../utils/tanggal";
import { cocok } from "../utils/cari";
import { statusVariant, PELATIHAN_STATUS } from "../utils/status";

const pelatihanStore = usePelatihanStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(pelatihanStore);

const karyawanId = ref("");
const nama = ref("");
const jenis = ref("");
const tanggalMulai = ref("");
const tanggalSelesai = ref("");
const lokasi = ref("");
const biaya = ref("");
const sertifikat = ref("");
const catatan = ref("");
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((p) =>
    cocok(
      cari.value,
      p.karyawan_id,
      p.nama,
      p.jenis,
      p.lokasi,
      p.status,
      p.tanggal_mulai,
      p.tanggal_selesai,
      p.catatan,
    ),
  ),
);

onMounted(() => {
  pelatihanStore.fetchAll();
});

async function handleSimpan() {
  if (
    !karyawanId.value.trim() ||
    !nama.value.trim() ||
    !tanggalMulai.value.trim()
  ) {
    toastStore.error("Karyawan, nama pelatihan, dan tanggal mulai wajib diisi");
    return;
  }
  try {
    await pelatihanStore.simpan({
      karyawan_id: Number(karyawanId.value),
      nama: nama.value,
      jenis: jenis.value,
      tanggal_mulai: tanggalMulai.value,
      tanggal_selesai: tanggalSelesai.value,
      lokasi: lokasi.value,
      biaya: Number(biaya.value) || 0,
      status: "terjadwal",
      sertifikat: sertifikat.value,
      catatan: catatan.value,
    });
    toastStore.success("Pelatihan berhasil disimpan");
    karyawanId.value = "";
    nama.value = "";
    jenis.value = "";
    tanggalMulai.value = "";
    tanggalSelesai.value = "";
    lokasi.value = "";
    biaya.value = "";
    sertifikat.value = "";
    catatan.value = "";
    pelatihanStore.fetchAll();
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm("Yakin ingin menghapus pelatihan ini?")) return;
  try {
    await pelatihanStore.hapus(id);
    toastStore.success("Pelatihan berhasil dihapus");
    pelatihanStore.fetchAll();
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
        <h1 class="text-2xl font-bold">Kelola Pelatihan</h1>
        <p class="text-text-muted">Data pelatihan karyawan</p>
      </div>
    </div>

    <div v-if="error" class="alert alert-error">
      <span>{{ error }}</span>
    </div>

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
            <span class="label-text">Nama Pelatihan</span>
          </label>
          <AppInput v-model="nama" placeholder="Nama pelatihan" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Jenis</span>
          </label>
          <AppInput v-model="jenis" placeholder="Contoh: Teknis, Soft Skill" />
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
          <AppInput v-model="tanggalSelesai" type="date" />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Lokasi</span>
          </label>
          <AppInput v-model="lokasi" placeholder="Lokasi pelatihan" />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Biaya</span>
          </label>
          <AppInput v-model="biaya" type="number" placeholder="0" />
        </div>

        <div class="w-full col-span-6">
          <AppButton type="submit" variant="primary" :loading="loading">
            Simpan
          </AppButton>
        </div>
      </div>

      <div class="form-control w-full">
        <label class="label">
          <span class="label-text">Sertifikat</span>
        </label>
        <AppInput
          v-model="sertifikat"
          placeholder="Nomor atau URL sertifikat"
        />
      </div>

      <div class="form-control w-full mt-4">
        <label class="label">
          <span class="label-text">Catatan</span>
        </label>
        <AppInput v-model="catatan" placeholder="Catatan pelatihan" />
      </div>
    </AppCard>

    <TabelData
      v-model:cari="cari"
      placeholder="Cari nama, jenis, lokasi, tanggal, atau status"
    >
      <template #head>
        <th>Karyawan ID</th>
        <th>Nama Pelatihan</th>
        <th>Jenis</th>
        <th>Tanggal Mulai</th>
        <th>Tanggal Selesai</th>
        <th>Biaya</th>
        <th>Status</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr v-for="p in filteredItems" :key="p.id" class="border-b">
        <td>{{ p.karyawan_id }}</td>
        <td>{{ p.nama }}</td>
        <td>{{ p.jenis || "-" }}</td>
        <td>{{ formatTanggal(p.tanggal_mulai) }}</td>
        <td>
          {{ p.tanggal_selesai ? formatTanggal(p.tanggal_selesai) : "-" }}
        </td>
        <td>{{ p.biaya }}</td>
        <td>
          <AppBadge
            :variant="statusVariant(p.status, PELATIHAN_STATUS)"
            size="sm"
          >
            {{ p.status }}
          </AppBadge>
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="handleHapus(p.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="8"
        kosong="Belum ada data pelatihan"
        kosong-cari="Tidak ada pelatihan yang cocok"
        :mencari="!!cari"
      />
    </TabelData>
  </div>
</template>
