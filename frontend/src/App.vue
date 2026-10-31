<script setup lang="ts">
import { computed, ref } from "vue";
import AdminLayout from "@/layouts/AdminLayout.vue";
import KaryawanPage from "@/pages/KaryawanPage.vue";
import AbsensiPage from "@/pages/AbsensiPage.vue";
import CutiPage from "@/pages/CutiPage.vue";
import ReimbursePage from "@/pages/ReimbursePage.vue";
import GajiPage from "@/pages/GajiPage.vue";
import JabatanPage from "@/pages/JabatanPage.vue";
import DepartemenPage from "@/pages/DepartemenPage.vue";
import PelatihanPage from "@/pages/PelatihanPage.vue";

const halaman = [
  { id: "karyawan", label: "Karyawan", komponen: KaryawanPage },
  { id: "absensi", label: "Absensi", komponen: AbsensiPage },
  { id: "cuti", label: "Cuti", komponen: CutiPage },
  { id: "reimburse", label: "Reimburse", komponen: ReimbursePage },
  { id: "gaji", label: "Gaji", komponen: GajiPage },
  { id: "jabatan", label: "Jabatan", komponen: JabatanPage },
  { id: "departemen", label: "Departemen", komponen: DepartemenPage },
  { id: "pelatihan", label: "Pelatihan", komponen: PelatihanPage },
];

const aktif = ref("karyawan");

const halamanAktif = computed(
  () =>
    halaman.find((h) => h.id === aktif.value)?.komponen || halaman[0].komponen,
);
</script>

<template>
  <AdminLayout v-model:aktif="aktif" :halaman="halaman">
    <component :is="halamanAktif" />
  </AdminLayout>
</template>
