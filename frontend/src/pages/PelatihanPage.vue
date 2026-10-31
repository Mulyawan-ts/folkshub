<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { usePelatihanStore } from "@/stores/index";
import { useToastStore } from "@/stores/toast";
import AppInput from "@/components/AppInput.vue";
import AppLabel from "@/components/AppLabel.vue";
import AppAlert from "@/components/AppAlert.vue";
import AppConfirm from "@/components/AppConfirm.vue";
import AppButton from "@/components/AppButton.vue";
import AppCard from "@/components/AppCard.vue";
import AppStatus from "@/components/AppStatus.vue";
import TabelData from "@/components/TabelData.vue";
import TabelKosong from "@/components/TabelKosong.vue";
import { formatTanggal } from "@/utils/tanggal";
import { cocok } from "@/utils/cari";
import { PELATIHAN_STATUS } from "@/utils/status";
import { useForm } from "@/composables/useForm";

const pelatihanStore = usePelatihanStore();
const toastStore = useToastStore();

const { items, error } = storeToRefs(pelatihanStore);

const form = useForm({
  awal: {
    karyawanId: "",
    nama: "",
    jenis: "",
    tanggalMulai: "",
    tanggalSelesai: "",
    lokasi: "",
    biaya: "",
    sertifikat: "",
    catatan: "",
  },
  validasi: (v) =>
    !v.karyawanId.trim() || !v.nama.trim() || !v.tanggalMulai.trim()
      ? "Karyawan, nama pelatihan, dan tanggal mulai wajib diisi"
      : "",
  aksi: (v) =>
    pelatihanStore.simpan({
      karyawan_id: Number(v.karyawanId),
      nama: v.nama,
      jenis: v.jenis,
      tanggal_mulai: v.tanggalMulai,
      tanggal_selesai: v.tanggalSelesai,
      lokasi: v.lokasi,
      biaya: Number(v.biaya) || 0,
      status: "terjadwal",
      sertifikat: v.sertifikat,
      catatan: v.catatan,
    }),
  pesanSukses: "Pelatihan berhasil disimpan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
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
        <p class="text-base-content/60">Data pelatihan karyawan</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="pelatihanStore.bersihkanError"
    />

    <AppCard as="form" padding="6" @submit.prevent="form.simpan">
      <div class="grid grid-cols-1 md:grid-cols-6 gap-4 items-end">
        <div class="form-control w-full">
          <AppLabel text="Karyawan ID" />
          <AppInput
            v-model="form.nilai.karyawanId"
            type="number"
            placeholder="ID karyawan"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Nama Pelatihan" />
          <AppInput
            v-model="form.nilai.nama"
            placeholder="Nama pelatihan"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Jenis" />
          <AppInput
            v-model="form.nilai.jenis"
            placeholder="Contoh: Teknis, Soft Skill"
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Tanggal Mulai" />
          <AppInput v-model="form.nilai.tanggalMulai" type="date" required />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Tanggal Selesai" />
          <AppInput v-model="form.nilai.tanggalSelesai" type="date" />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Lokasi" />
          <AppInput
            v-model="form.nilai.lokasi"
            placeholder="Lokasi pelatihan"
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Biaya" />
          <AppInput v-model="form.nilai.biaya" type="number" placeholder="0" />
        </div>

        <div class="w-full col-span-6">
          <AppButton
            type="submit"
            variant="primary"
            :loading="form.sedangSimpan.value"
          >
            Simpan
          </AppButton>
        </div>
      </div>

      <div class="form-control w-full">
        <AppLabel text="Sertifikat" />
        <AppInput
          v-model="form.nilai.sertifikat"
          placeholder="Nomor atau URL sertifikat"
        />
      </div>

      <div class="form-control w-full mt-4">
        <AppLabel text="Catatan" />
        <AppInput
          v-model="form.nilai.catatan"
          placeholder="Catatan pelatihan"
        />
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
          <AppStatus :status="p.status" :map="PELATIHAN_STATUS" />
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(p.id)"
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

    <AppConfirm
      :terbuka="konfirmasiHapus"
      judul="Hapus pelatihan"
      pesan="Yakin ingin menghapus pelatihan ini?"
      confirm-label="Hapus"
      variant="error"
      @konfirmasi="jalankanHapus"
      @batal="konfirmasiHapus = false"
    />
  </div>
</template>
