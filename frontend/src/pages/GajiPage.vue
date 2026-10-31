<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useGajiStore, useKaryawanStore } from "@/stores/index";
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
import { formatBulan } from "@/utils/tanggal";
import { cocok } from "@/utils/cari";
import { GAJI_STATUS } from "@/utils/status";
import { useForm } from "@/composables/useForm";

const gajiStore = useGajiStore();
const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, error } = storeToRefs(gajiStore);

const form = useForm({
  awal: {
    karyawanId: "",
    bulan: "",
    gajiPokok: "",
    tunjangan: "",
    potongan: "",
  },
  validasi: (nilai) =>
    !nilai.karyawanId.trim() || !nilai.bulan.trim() || !nilai.gajiPokok.trim()
      ? "Karyawan, bulan, dan gaji pokok wajib diisi"
      : "",
  aksi: (nilai) =>
    gajiStore.simpan({
      karyawan_id: Number(nilai.karyawanId),
      bulan: nilai.bulan,
      gaji_pokok: Number(nilai.gajiPokok),
      tunjangan: Number(nilai.tunjangan) || 0,
      potongan: Number(nilai.potongan) || 0,
      status: "draft",
    }),
  pesanSukses: "Data gaji berhasil disimpan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((gaji) =>
    cocok(
      cari.value,
      gaji.karyawan_id,
      gaji.bulan,
      formatBulan(gaji.bulan),
      gaji.status,
      gaji.keterangan,
    ),
  ),
);

onMounted(() => {
  karyawanStore.fetchAll();
  gajiStore.fetchAll();
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
    await gajiStore.hapus(id);
    toastStore.success("Gaji berhasil dihapus");
    gajiStore.fetchAll();
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
        <h1 class="text-2xl font-bold">Kelola Gaji</h1>
        <p class="text-base-content/60">Data slip gaji karyawan</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="gajiStore.bersihkanError"
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
          <AppLabel text="Bulan" />
          <AppInput v-model="form.nilai.bulan" type="month" required />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Gaji Pokok" />
          <AppInput
            v-model="form.nilai.gajiPokok"
            type="number"
            placeholder="0"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Tunjangan" />
          <AppInput
            v-model="form.nilai.tunjangan"
            type="number"
            placeholder="0"
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Potongan" />
          <AppInput
            v-model="form.nilai.potongan"
            type="number"
            placeholder="0"
          />
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
    </AppCard>

    <TabelData
      v-model:cari="cari"
      placeholder="Cari karyawan, bulan, status, atau keterangan"
    >
      <template #head>
        <th>Karyawan ID</th>
        <th>Bulan</th>
        <th>Gaji Pokok</th>
        <th>Tunjangan</th>
        <th>Potongan</th>
        <th>Total</th>
        <th>Status</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr v-for="gaji in filteredItems" :key="gaji.id" class="border-b">
        <td>{{ gaji.karyawan_id }}</td>
        <td>{{ formatBulan(gaji.bulan) }}</td>
        <td>{{ gaji.gaji_pokok }}</td>
        <td>{{ gaji.tunjangan }}</td>
        <td>{{ gaji.potongan }}</td>
        <td class="font-bold">{{ gaji.total }}</td>
        <td>
          <AppStatus :status="gaji.status" :map="GAJI_STATUS" />
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(gaji.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="8"
        kosong="Belum ada data gaji"
        kosong-cari="Tidak ada gaji yang cocok"
        :mencari="!!cari"
      />
    </TabelData>

    <AppConfirm
      :terbuka="konfirmasiHapus"
      judul="Hapus gaji"
      pesan="Yakin ingin menghapus data gaji ini?"
      confirm-label="Hapus"
      variant="error"
      @konfirmasi="jalankanHapus"
      @batal="konfirmasiHapus = false"
    />
  </div>
</template>
