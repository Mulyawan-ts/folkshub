<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { usePelatihanStore, useKaryawanStore } from "@/stores/index";
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
import { karyawanOptions } from "@/utils/karyawan";

const pelatihanStore = usePelatihanStore();
const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, error } = storeToRefs(pelatihanStore);
const { items: daftarKaryawan } = storeToRefs(karyawanStore);

const opsiKaryawan = computed(() => karyawanOptions(daftarKaryawan.value));

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
  validasi: (nilai) =>
    !nilai.karyawanId.trim() || !nilai.nama.trim() || !nilai.tanggalMulai.trim()
      ? "Karyawan, nama pelatihan, dan tanggal mulai wajib diisi"
      : "",
  aksi: (nilai) =>
    pelatihanStore.simpan({
      karyawan_id: Number(nilai.karyawanId),
      nama: nilai.nama,
      jenis: nilai.jenis,
      tanggal_mulai: nilai.tanggalMulai,
      tanggal_selesai: nilai.tanggalSelesai,
      lokasi: nilai.lokasi,
      biaya: Number(nilai.biaya) || 0,
      status: "terjadwal",
      sertifikat: nilai.sertifikat,
      catatan: nilai.catatan,
    }),
  pesanSukses: "Pelatihan berhasil disimpan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((pelatihan) =>
    cocok(
      cari.value,
      pelatihan.karyawan_id,
      pelatihan.nama,
      pelatihan.jenis,
      pelatihan.lokasi,
      pelatihan.status,
      pelatihan.tanggal_mulai,
      pelatihan.tanggal_selesai,
      pelatihan.catatan,
    ),
  ),
);

onMounted(() => {
  karyawanStore.fetchAll();
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
        <AppInput
          v-model="form.nilai.karyawanId"
          label="Karyawan"
          tipe="select"
          :options="opsiKaryawan"
          placeholder="Pilih karyawan"
          required
        />

        <AppInput
          v-model="form.nilai.nama"
          label="Nama Pelatihan"
          placeholder="Nama pelatihan"
          required
        />

        <AppInput
          v-model="form.nilai.jenis"
          label="Jenis"
          placeholder="Contoh: Teknis, Soft Skill"
        />

        <AppInput
          v-model="form.nilai.tanggalMulai"
          label="Tanggal Mulai"
          type="date"
          required
        />

        <AppInput
          v-model="form.nilai.tanggalSelesai"
          label="Tanggal Selesai"
          type="date"
        />

        <AppInput
          v-model="form.nilai.lokasi"
          label="Lokasi"
          placeholder="Lokasi pelatihan"
        />

        <AppInput
          v-model="form.nilai.biaya"
          label="Biaya"
          type="number"
          placeholder="0"
        />

        <div class="w-full col-span-6">
          <AppButton
            type="submit"
            variant="primary"
            :loading="form.sedangSimpan"
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

      <tr
        v-for="pelatihan in filteredItems"
        :key="pelatihan.id"
        class="border-b"
      >
        <td>{{ pelatihan.karyawan_id }}</td>
        <td>{{ pelatihan.nama }}</td>
        <td>{{ pelatihan.jenis || "-" }}</td>
        <td>{{ formatTanggal(pelatihan.tanggal_mulai) }}</td>
        <td>
          {{
            pelatihan.tanggal_selesai
              ? formatTanggal(pelatihan.tanggal_selesai)
              : "-"
          }}
        </td>
        <td>{{ pelatihan.biaya }}</td>
        <td>
          <AppStatus :status="pelatihan.status" :map="PELATIHAN_STATUS" />
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(pelatihan.id)"
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
