<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useReimburseStore, useKaryawanStore } from "@/stores/index";
import { useToastStore } from "@/stores/toast";
import AppInput from "@/components/AppInput.vue";
import AppAlert from "@/components/AppAlert.vue";
import AppConfirm from "@/components/AppConfirm.vue";
import AppButton from "@/components/AppButton.vue";
import AppCard from "@/components/AppCard.vue";
import AppBadge from "@/components/AppBadge.vue";
import TabelData from "@/components/TabelData.vue";
import TabelKosong from "@/components/TabelKosong.vue";
import { formatTanggal } from "@/utils/tanggal";
import { cocok } from "@/utils/cari";
import { statusVariant, REIMBURSE_STATUS } from "@/utils/status";
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
  validasi: (v) =>
    !v.karyawanId.trim() ||
    !v.tanggal.trim() ||
    !v.jenis.trim() ||
    !v.jumlah.trim()
      ? "Karyawan, tanggal, jenis, dan jumlah wajib diisi"
      : "",
  aksi: (v) =>
    reimburseStore.simpan({
      karyawan_id: Number(v.karyawanId),
      tanggal: v.tanggal,
      jenis: v.jenis,
      jumlah: Number(v.jumlah),
      keterangan: v.keterangan,
      bukti: v.bukti,
      status: "draft",
    }),
  pesanSukses: "Reimburse berhasil diajukan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((r) =>
    cocok(
      cari.value,
      r.karyawan_id,
      r.tanggal,
      r.jenis,
      r.status,
      r.keterangan,
      r.bukti,
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
          <label class="label">
            <span class="label-text">Karyawan ID</span>
          </label>
          <AppInput
            v-model="form.nilai.karyawanId"
            type="number"
            placeholder="ID karyawan"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Tanggal</span>
          </label>
          <AppInput v-model="form.nilai.tanggal" type="date" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Jenis Biaya</span>
          </label>
          <AppInput
            v-model="form.nilai.jenis"
            placeholder="Contoh: Transport, Makan, Akomodasi"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Jumlah</span>
          </label>
          <AppInput
            v-model="form.nilai.jumlah"
            type="number"
            placeholder="0"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Bukti</span>
          </label>
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
        <label class="label">
          <span class="label-text">Keterangan</span>
        </label>
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

      <tr v-for="r in filteredItems" :key="r.id" class="border-b">
        <td>{{ r.karyawan_id }}</td>
        <td>{{ formatTanggal(r.tanggal) }}</td>
        <td>{{ r.jenis }}</td>
        <td>{{ r.jumlah }}</td>
        <td>{{ r.bukti || "-" }}</td>
        <td>
          <AppBadge
            :variant="statusVariant(r.status, REIMBURSE_STATUS)"
            size="sm"
          >
            {{ r.status }}
          </AppBadge>
        </td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(r.id)"
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
