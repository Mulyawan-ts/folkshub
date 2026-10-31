<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useKaryawanStore } from "@/stores/karyawan";
import { useToastStore } from "@/stores/toast";
import AppInput from "@/components/AppInput.vue";
import AppLabel from "@/components/AppLabel.vue";
import AppAlert from "@/components/AppAlert.vue";
import AppConfirm from "@/components/AppConfirm.vue";
import AppButton from "@/components/AppButton.vue";
import AppCard from "@/components/AppCard.vue";
import AppCardItem from "@/components/AppCardItem.vue";
import DaftarData from "@/components/DaftarData.vue";
import { cocok } from "@/utils/cari";
import { statusVariant, KARYAWAN_STATUS } from "@/utils/status";
import { useForm } from "@/composables/useForm";

const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, error } = storeToRefs(karyawanStore);

const form = useForm({
  awal: { nama: "", email: "", status: "tetap" },
  validasi: (nilai) =>
    !nilai.nama.trim() || !nilai.email.trim()
      ? "Nama dan email wajib diisi"
      : "",
  aksi: (nilai) =>
    karyawanStore.simpan({
      nama: nilai.nama.trim(),
      email: nilai.email.trim(),
      status: nilai.status,
    }),
  pesanSukses: "Karyawan berhasil disimpan",
  toast: toastStore,
});

const cari = ref("");
const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);

const filteredItems = computed(() =>
  items.value.filter((karyawan) =>
    cocok(cari.value, karyawan.nama, karyawan.email, karyawan.status),
  ),
);

const statusOptions = [
  { value: "tetap", label: "Tetap" },
  { value: "outsourcing", label: "Outsourcing" },
  { value: "magang", label: "Magang" },
];

onMounted(() => {
  karyawanStore.fetchAll();
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
        <p class="text-base-content/60">
          Data karyawan (tetap, outsourcing, magang)
        </p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="karyawanStore.bersihkanError"
    />

    <AppCard as="form" padding="6" @submit.prevent="form.simpan">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4 items-end">
        <div class="form-control w-full">
          <AppLabel text="Nama" />
          <AppInput
            v-model="form.nilai.nama"
            placeholder="Masukkan nama"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Email" />
          <AppInput
            v-model="form.nilai.email"
            type="email"
            placeholder="nama@contoh.com"
            required
          />
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
            :loading="form.sedangSimpan.value"
          >
            Simpan
          </AppButton>
        </div>
      </div>
    </AppCard>

    <DaftarData
      v-model:cari="cari"
      placeholder="Cari nama, email, atau status"
      kosong="Belum ada data karyawan"
      kosong-cari="Tidak ada karyawan yang cocok"
      :jumlah="filteredItems.length"
    >
      <AppCardItem
        v-for="karyawan in filteredItems"
        :key="karyawan.id"
        :judul="karyawan.nama"
        :subjudul="karyawan.email"
        :status="karyawan.status"
        :status-variant="statusVariant(karyawan.status, KARYAWAN_STATUS)"
        hapus-label="Hapus"
        @hapus="mintaHapus(karyawan.id)"
      />
    </DaftarData>

    <AppConfirm
      :terbuka="konfirmasiHapus"
      judul="Hapus karyawan"
      pesan="Yakin ingin menghapus karyawan ini?"
      confirm-label="Hapus"
      variant="error"
      @konfirmasi="jalankanHapus"
      @batal="konfirmasiHapus = false"
    />
  </div>
</template>
