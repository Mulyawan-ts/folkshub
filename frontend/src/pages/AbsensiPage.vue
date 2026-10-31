<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useAbsensiStore, useKaryawanStore } from "@/stores/index";
import type { StatusAbsensi } from "@/stores/index";
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
import { ABSENSI_STATUS } from "@/utils/status";
import { useForm } from "@/composables/useForm";
import { karyawanOptions } from "@/utils/karyawan";

const absensiStore = useAbsensiStore();
const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, error } = storeToRefs(absensiStore);
const { items: daftarKaryawan } = storeToRefs(karyawanStore);

const opsiKaryawan = computed(() => karyawanOptions(daftarKaryawan.value));

const form = useForm({
  awal: {
    karyawanId: "",
    tanggal: "",
    checkIn: "",
    checkOut: "",
    status: "hadir",
    keterangan: "",
  },
  validasi: (nilai) =>
    !nilai.karyawanId.trim() || !nilai.tanggal.trim()
      ? "Karyawan dan tanggal wajib diisi"
      : "",
  aksi: (nilai) =>
    absensiStore.simpan({
      karyawan_id: Number(nilai.karyawanId),
      tanggal: nilai.tanggal,
      check_in: nilai.checkIn || undefined,
      check_out: nilai.checkOut || undefined,
      status: nilai.status as StatusAbsensi,
      keterangan: nilai.keterangan,
    }),
  pesanSukses: "Absensi berhasil disimpan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((absensi) =>
    cocok(
      cari.value,
      absensi.karyawan_id,
      absensi.tanggal,
      formatTanggal(absensi.tanggal),
      absensi.status,
      absensi.keterangan,
    ),
  ),
);

const statusOptions = [
  { value: "hadir", label: "Hadir" },
  { value: "izin", label: "Izin" },
  { value: "sakit", label: "Sakit" },
  { value: "alfa", label: "Alfa" },
  { value: "lembur", label: "Lembur" },
];

onMounted(() => {
  karyawanStore.fetchAll();
  absensiStore.fetchAll();
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
        <p class="text-base-content/60">Data kehadiran karyawan</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="absensiStore.bersihkanError"
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

        <div class="form-control w-full">
          <AppLabel text="Tanggal" />
          <AppInput v-model="form.nilai.tanggal" type="date" required />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Check In" />
          <AppInput v-model="form.nilai.checkIn" type="time" />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Check Out" />
          <AppInput v-model="form.nilai.checkOut" type="time" />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Status" />
          <select
            v-model="form.nilai.status"
            class="select select-bordered w-full"
          >
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
          <AppButton
            type="submit"
            variant="primary"
            :loading="form.sedangSimpan"
          >
            Simpan
          </AppButton>
        </div>
      </div>

      <div class="form-control w-full mt-4">
        <AppLabel text="Keterangan" />
        <AppInput
          v-model="form.nilai.keterangan"
          placeholder="Catatan tambahan"
        />
      </div>
    </AppCard>

    <TabelData
      v-model:cari="cari"
      placeholder="Cari karyawan, tanggal, status, atau keterangan"
    >
      <template #head>
        <th>Karyawan ID</th>
        <th>Tanggal</th>
        <th>Check In</th>
        <th>Check Out</th>
        <th>Status</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr v-for="absensi in filteredItems" :key="absensi.id" class="border-b">
        <td>{{ absensi.karyawan_id }}</td>
        <td>{{ formatTanggal(absensi.tanggal) }}</td>
        <td>{{ absensi.check_in || "-" }}</td>
        <td>{{ absensi.check_out || "-" }}</td>
        <td>
          <AppStatus :status="absensi.status" :map="ABSENSI_STATUS" />
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(absensi.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="6"
        kosong="Belum ada data absensi"
        kosong-cari="Tidak ada absensi yang cocok"
        :mencari="!!cari"
      />
    </TabelData>

    <AppConfirm
      :terbuka="konfirmasiHapus"
      judul="Hapus absensi"
      pesan="Yakin ingin menghapus absensi ini?"
      confirm-label="Hapus"
      variant="error"
      @konfirmasi="jalankanHapus"
      @batal="konfirmasiHapus = false"
    />
  </div>
</template>
