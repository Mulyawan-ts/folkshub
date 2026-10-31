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
  validasi: (v) =>
    !v.karyawanId.trim() || !v.bulan.trim() || !v.gajiPokok.trim()
      ? "Karyawan, bulan, dan gaji pokok wajib diisi"
      : "",
  aksi: (v) =>
    gajiStore.simpan({
      karyawan_id: Number(v.karyawanId),
      bulan: v.bulan,
      gaji_pokok: Number(v.gajiPokok),
      tunjangan: Number(v.tunjangan) || 0,
      potongan: Number(v.potongan) || 0,
      status: "draft",
    }),
  pesanSukses: "Data gaji berhasil disimpan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((g) =>
    cocok(
      cari.value,
      g.karyawan_id,
      g.bulan,
      formatBulan(g.bulan),
      g.status,
      g.keterangan,
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
            :loading="form.sedangSimpan.value"
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

      <tr v-for="g in filteredItems" :key="g.id" class="border-b">
        <td>{{ g.karyawan_id }}</td>
        <td>{{ formatBulan(g.bulan) }}</td>
        <td>{{ g.gaji_pokok }}</td>
        <td>{{ g.tunjangan }}</td>
        <td>{{ g.potongan }}</td>
        <td class="font-bold">{{ g.total }}</td>
        <td>
          <AppStatus :status="g.status" :map="GAJI_STATUS" />
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(g.id)"
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
