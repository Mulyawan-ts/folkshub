<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useDepartemenStore } from "@/stores/index";
import { useToastStore } from "@/stores/toast";
import AppInput from "@/components/AppInput.vue";
import AppAlert from "@/components/AppAlert.vue";
import AppConfirm from "@/components/AppConfirm.vue";
import AppButton from "@/components/AppButton.vue";
import AppCard from "@/components/AppCard.vue";
import TabelData from "@/components/TabelData.vue";
import TabelKosong from "@/components/TabelKosong.vue";
import { cocok } from "@/utils/cari";
import { useForm } from "@/composables/useForm";

const departemenStore = useDepartemenStore();
const toastStore = useToastStore();

const { items, error } = storeToRefs(departemenStore);

const form = useForm({
  awal: { kode: "", nama: "", kepala: "", deskripsi: "" },
  validasi: (v) =>
    !v.kode.trim() || !v.nama.trim()
      ? "Kode dan nama departemen wajib diisi"
      : "",
  aksi: (v) =>
    departemenStore.simpan({
      kode: v.kode,
      nama: v.nama,
      kepala: v.kepala,
      deskripsi: v.deskripsi,
    }),
  pesanSukses: "Departemen berhasil disimpan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((d) =>
    cocok(cari.value, d.kode, d.nama, d.kepala, d.deskripsi),
  ),
);

onMounted(() => {
  departemenStore.fetchAll();
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
    await departemenStore.hapus(id);
    toastStore.success("Departemen berhasil dihapus");
    departemenStore.fetchAll();
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
        <h1 class="text-2xl font-bold">Kelola Departemen</h1>
        <p class="text-base-content/60">Data departemen organisasi</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="departemenStore.bersihkanError"
    />

    <AppCard as="form" padding="6" @submit.prevent="form.simpan">
      <div class="grid grid-cols-1 md:grid-cols-6 gap-4 items-end">
        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Kode Departemen</span>
          </label>
          <AppInput
            v-model="form.nilai.kode"
            placeholder="Contoh: HRD, FIN"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Nama Departemen</span>
          </label>
          <AppInput
            v-model="form.nilai.nama"
            placeholder="Contoh: Human Resources"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Kepala</span>
          </label>
          <AppInput
            v-model="form.nilai.kepala"
            placeholder="Nama kepala departemen"
          />
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

      <div class="form-control w-full mt-4">
        <label class="label">
          <span class="label-text">Deskripsi</span>
        </label>
        <AppInput
          v-model="form.nilai.deskripsi"
          placeholder="Deskripsi departemen"
        />
      </div>
    </AppCard>

    <TabelData v-model:cari="cari" placeholder="Cari kode, nama, atau kepala">
      <template #head>
        <th>Kode</th>
        <th>Nama Departemen</th>
        <th>Kepala</th>
        <th>Deskripsi</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr v-for="d in filteredItems" :key="d.id" class="border-b">
        <td>{{ d.kode }}</td>
        <td>{{ d.nama }}</td>
        <td>{{ d.kepala || "-" }}</td>
        <td>{{ d.deskripsi || "-" }}</td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(d.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="5"
        kosong="Belum ada data departemen"
        kosong-cari="Tidak ada departemen yang cocok"
        :mencari="!!cari"
      />
    </TabelData>

    <AppConfirm
      :terbuka="konfirmasiHapus"
      judul="Hapus departemen"
      pesan="Yakin ingin menghapus departemen ini?"
      confirm-label="Hapus"
      variant="error"
      @konfirmasi="jalankanHapus"
      @batal="konfirmasiHapus = false"
    />
  </div>
</template>
