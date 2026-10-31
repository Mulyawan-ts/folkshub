<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useReimburseStore, useKaryawanStore } from "@/stores/index";
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
import { REIMBURSE_STATUS } from "@/utils/status";
import { useForm } from "@/composables/useForm";

const reimburseStore = useReimburseStore();
const karyawanStore = useKaryawanStore();
const toastStore = useToastStore();

const { items, error } = storeToRefs(reimburseStore);

const form = useForm({
  awal: {
    karyawanId: "",
    tanggal: "",
    jenis: "",
    jumlah: "",
    keterangan: "",
    bukti: "",
  },
  validasi: (nilai) =>
    !nilai.karyawanId.trim() ||
    !nilai.tanggal.trim() ||
    !nilai.jenis.trim() ||
    !nilai.jumlah.trim()
      ? "Karyawan, tanggal, jenis, dan jumlah wajib diisi"
      : "",
  aksi: (nilai) =>
    reimburseStore.simpan({
      karyawan_id: Number(nilai.karyawanId),
      tanggal: nilai.tanggal,
      jenis: nilai.jenis,
      jumlah: Number(nilai.jumlah),
      keterangan: nilai.keterangan,
      bukti: nilai.bukti,
      status: "draft",
    }),
  pesanSukses: "Reimburse berhasil diajukan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((reimburse) =>
    cocok(
      cari.value,
      reimburse.karyawan_id,
      reimburse.tanggal,
      reimburse.jenis,
      reimburse.status,
      reimburse.keterangan,
      reimburse.bukti,
    ),
  ),
);

onMounted(() => {
  karyawanStore.fetchAll();
  reimburseStore.fetchAll();
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
        <p class="text-base-content/60">Klaim biaya karyawan</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="reimburseStore.bersihkanError"
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
          <AppLabel text="Tanggal" />
          <AppInput v-model="form.nilai.tanggal" type="date" required />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Jenis Biaya" />
          <AppInput
            v-model="form.nilai.jenis"
            placeholder="Contoh: Transport, Makan, Akomodasi"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Jumlah" />
          <AppInput
            v-model="form.nilai.jumlah"
            type="number"
            placeholder="0"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Bukti" />
          <AppInput
            v-model="form.nilai.bukti"
            placeholder="URL atau nomor bukti"
          />
        </div>

        <div class="w-full">
          <AppButton
            type="submit"
            variant="primary"
            :loading="form.sedangSimpan.value"
          >
            Ajukan
          </AppButton>
        </div>
      </div>

      <div class="form-control w-full mt-4">
        <AppLabel text="Keterangan" />
        <AppInput
          v-model="form.nilai.keterangan"
          placeholder="Alasan pengajuan reimburse"
        />
      </div>
    </AppCard>

    <TabelData
      v-model:cari="cari"
      placeholder="Cari karyawan, jenis, status, atau keterangan"
    >
      <template #head>
        <th>Karyawan ID</th>
        <th>Tanggal</th>
        <th>Jenis</th>
        <th>Jumlah</th>
        <th>Bukti</th>
        <th>Status</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr
        v-for="reimburse in filteredItems"
        :key="reimburse.id"
        class="border-b"
      >
        <td>{{ reimburse.karyawan_id }}</td>
        <td>{{ formatTanggal(reimburse.tanggal) }}</td>
        <td>{{ reimburse.jenis }}</td>
        <td>{{ reimburse.jumlah }}</td>
        <td>{{ reimburse.bukti || "-" }}</td>
        <td>
          <AppStatus :status="reimburse.status" :map="REIMBURSE_STATUS" />
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(reimburse.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="7"
        kosong="Belum ada data reimburse"
        kosong-cari="Tidak ada reimburse yang cocok"
        :mencari="!!cari"
      />
    </TabelData>

    <AppConfirm
      :terbuka="konfirmasiHapus"
      judul="Hapus reimburse"
      pesan="Yakin ingin menghapus reimburse ini?"
      confirm-label="Hapus"
      variant="error"
      @konfirmasi="jalankanHapus"
      @batal="konfirmasiHapus = false"
    />
  </div>
</template>
