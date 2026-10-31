<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useJabatanStore } from "../stores/index";
import { useToastStore } from "../stores/toast";
import AppInput from "../components/AppInput.vue";
import AppAlert from "../components/AppAlert.vue";
import AppButton from "../components/AppButton.vue";
import AppCard from "../components/AppCard.vue";
import TabelData from "../components/TabelData.vue";
import TabelKosong from "../components/TabelKosong.vue";
import { cocok } from "../utils/cari";

const jabatanStore = useJabatanStore();
const toastStore = useToastStore();

const { items, loading, error } = storeToRefs(jabatanStore);

const kode = ref("");
const nama = ref("");
const departemen = ref("");
const gajiPokok = ref("");
const tunjangan = ref("");
const deskripsi = ref("");
const cari = ref("");

const filteredItems = computed(() =>
  items.value.filter((j) =>
    cocok(cari.value, j.kode, j.nama, j.departemen, j.deskripsi),
  ),
);

onMounted(() => {
  jabatanStore.fetchAll();
});

async function handleSimpan() {
  if (!kode.value.trim() || !nama.value.trim() || !departemen.value.trim()) {
    toastStore.error("Kode, nama, dan departemen wajib diisi");
    return;
  }
  try {
    await jabatanStore.simpan({
      kode: kode.value,
      nama: nama.value,
      departemen: departemen.value,
      gaji_pokok: Number(gajiPokok.value) || 0,
      tunjangan: Number(tunjangan.value) || 0,
      deskripsi: deskripsi.value,
    });
    toastStore.success("Jabatan berhasil disimpan");
    kode.value = "";
    nama.value = "";
    departemen.value = "";
    gajiPokok.value = "";
    tunjangan.value = "";
    deskripsi.value = "";
    jabatanStore.fetchAll();
  } catch {
    // error already set in store
  }
}

async function handleHapus(id: number) {
  if (!confirm("Yakin ingin menghapus jabatan ini?")) return;
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
        <p class="text-text-muted">Data jabatan dan departemen</p>
      </div>
    </div>

    <AppAlert
      v-if="error"
      variant="error"
      :pesan="error"
      dismissible
      @dismiss="jabatanStore.bersihkanError"
    />

    <AppCard as="form" padding="6" @submit.prevent="handleSimpan">
      <div class="grid grid-cols-1 md:grid-cols-6 gap-4 items-end">
        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Kode Jabatan</span>
          </label>
          <AppInput v-model="kode" placeholder="Contoh: MGR, STAFF" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Nama Jabatan</span>
          </label>
          <AppInput v-model="nama" placeholder="Contoh: Manager" required />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Departemen</span>
          </label>
          <AppInput
            v-model="departemen"
            placeholder="Contoh: HRD, Keuangan"
            required
          />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Gaji Pokok</span>
          </label>
          <AppInput v-model="gajiPokok" type="number" placeholder="0" />
        </div>

        <div class="form-control w-full">
          <label class="label">
            <span class="label-text">Tunjangan</span>
          </label>
          <AppInput v-model="tunjangan" type="number" placeholder="0" />
        </div>

        <div class="w-full">
          <AppButton type="submit" variant="primary" :loading="loading">
            Simpan
          </AppButton>
        </div>
      </div>

      <div class="form-control w-full mt-4">
        <label class="label">
          <span class="label-text">Deskripsi</span>
        </label>
        <AppInput v-model="deskripsi" placeholder="Deskripsi jabatan" />
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
            @click="handleHapus(j.id)"
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
  </div>
</template>
