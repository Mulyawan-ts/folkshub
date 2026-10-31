<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useJabatanStore } from "@/stores/index";
import { useToastStore } from "@/stores/toast";
import AppInput from "@/components/AppInput.vue";
import AppLabel from "@/components/AppLabel.vue";
import AppAlert from "@/components/AppAlert.vue";
import AppConfirm from "@/components/AppConfirm.vue";
import AppButton from "@/components/AppButton.vue";
import AppCard from "@/components/AppCard.vue";
import TabelData from "@/components/TabelData.vue";
import TabelKosong from "@/components/TabelKosong.vue";
import { cocok } from "@/utils/cari";
import { useForm } from "@/composables/useForm";

const jabatanStore = useJabatanStore();
const toastStore = useToastStore();

const { items, error } = storeToRefs(jabatanStore);

const form = useForm({
  awal: {
    kode: "",
    nama: "",
    departemen: "",
    gajiPokok: "",
    tunjangan: "",
    deskripsi: "",
  },
  validasi: (v) =>
    !v.kode.trim() || !v.nama.trim() ? "Kode dan nama jabatan wajib diisi" : "",
  aksi: (v) =>
    jabatanStore.simpan({
      kode: v.kode,
      nama: v.nama,
      departemen: v.departemen,
      deskripsi: v.deskripsi,
      gaji_pokok: Number(v.gajiPokok) || 0,
      tunjangan: Number(v.tunjangan) || 0,
    }),
  pesanSukses: "Jabatan berhasil disimpan",
  toast: toastStore,
});

const konfirmasiHapus = ref(false);
const hapusId = ref<number | null>(null);
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((j) =>
    cocok(cari.value, j.kode, j.nama, j.departemen, j.deskripsi),
  ),
);

onMounted(() => {
  jabatanStore.fetchAll();
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
    await jabatanStore.hapus(id);
    toastStore.success("Jabatan berhasil dihapus");
    jabatanStore.fetchAll();
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
        <h1 class="text-2xl font-bold">Kelola Jabatan</h1>
        <p class="text-base-content/60">Data jabatan dan departemen</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="jabatanStore.bersihkanError"
    />

    <AppCard as="form" padding="6" @submit.prevent="form.simpan">
      <div class="grid grid-cols-1 md:grid-cols-6 gap-4 items-end">
        <div class="form-control w-full">
          <AppLabel text="Kode Jabatan" />
          <AppInput
            v-model="form.nilai.kode"
            placeholder="Contoh: MGR, STAFF"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Nama Jabatan" />
          <AppInput
            v-model="form.nilai.nama"
            placeholder="Contoh: Manager"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Departemen" />
          <AppInput
            v-model="form.nilai.departemen"
            placeholder="Contoh: HRD, Keuangan"
            required
          />
        </div>

        <div class="form-control w-full">
          <AppLabel text="Gaji Pokok" />
          <AppInput
            v-model="form.nilai.gajiPokok"
            type="number"
            placeholder="0"
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

      <div class="form-control w-full mt-4">
        <AppLabel text="Deskripsi" />
        <AppInput
          v-model="form.nilai.deskripsi"
          placeholder="Deskripsi jabatan"
        />
      </div>
    </AppCard>

    <TabelData
      v-model:cari="cari"
      placeholder="Cari kode, nama, atau departemen"
    >
      <template #head>
        <th>Kode</th>
        <th>Nama Jabatan</th>
        <th>Departemen</th>
        <th>Gaji Pokok</th>
        <th>Tunjangan</th>
        <th>Deskripsi</th>
        <th class="text-right">Aksi</th>
      </template>

      <tr v-for="j in filteredItems" :key="j.id" class="border-b">
        <td>{{ j.kode }}</td>
        <td>{{ j.nama }}</td>
        <td>{{ j.departemen }}</td>
        <td>{{ j.gaji_pokok }}</td>
        <td>{{ j.tunjangan }}</td>
        <td>{{ j.deskripsi || "-" }}</td>
        <td class="text-right">
          <AppButton
            variant="error"
            size="sm"
            outline
            @click="mintaHapus(j.id)"
          >
            Hapus
          </AppButton>
        </td>
      </tr>

      <TabelKosong
        v-if="filteredItems.length === 0"
        :colspan="7"
        kosong="Belum ada data jabatan"
        kosong-cari="Tidak ada jabatan yang cocok"
        :mencari="!!cari"
      />
    </TabelData>

    <AppConfirm
      :terbuka="konfirmasiHapus"
      judul="Hapus jabatan"
      pesan="Yakin ingin menghapus jabatan ini?"
      confirm-label="Hapus"
      variant="error"
      @konfirmasi="jalankanHapus"
      @batal="konfirmasiHapus = false"
    />
  </div>
</template>
