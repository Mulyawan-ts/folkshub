# STRICT DEVELOPMENT & CODING RULES

## 1. Anti-AI-Slop & Gaya Komunikasi

- **Langsung ke Kode:** DILARANG menggunakan basa-basi pembuka/penutup (seperti "Tentu!", "Berikut kodenya:", "Sebagai AI...").
- **Tanpa Penjelasan Ulang:** Jangan menerangkan ulang kode yang baru saja diubah kecuali diminta. Berikan ringkasan teknis maksimal 2 kalimat jika benar-benar diperlukan.
- **Bebas Komentar Sepele:** DILARANG menulis komentar AI Slop yang tidak berguna (contoh: `// fungsi untuk submit`). Komentar HANYA untuk logika bisnis yang sangat kompleks.
- **Prinsip YAGNI:** HANYA ubah atau tambahkan kode yang diminta. DILARANG membuat abstraksi berlebihan, menambahkan dependensi baru, atau merefactor file di luar cakupan tugas.
- **Kode Utuh:** DILARANG menggunakan placeholder menggantung seperti `// ... sisa kode di sini ...`. Tuliskan kode secara lengkap dan siap dijalankan.

## 2. Vue 3 Composition API & SFC Standards

- **Struktur SFC:** Urutan tag di dalam file `.vue` WAJIB konsisten:
  1. `<script setup lang="ts">` (atau `lang="js"`)
  2. `<template>`
  3. `<style scoped>`
- **Script Setup Only:** WAJIB menggunakan `<script setup>`. DILARANG menggunakan Options API (`export default { ... }`) atau `defineComponent()` biasa.
- **Aturan Reaktivitas:**
  - Gunakan `ref()` sebagai standar state reaktif. Gunakan `reactive()` hanya untuk objek kompleks yang tidak akan di-reassign.
  - Deklarasikan props via compiler macro: `defineProps<{ ... }>()`.
  - Deklarasikan emits via compiler macro: `defineEmits<{ ... }>()`.
  - Gunakan `computed()` untuk semua nilai turunan sinkron. DILARANG menggunakan `watch()` untuk menghitung ulang state turunan.
  - Hindari manipulasi DOM langsung (`document.getElementById`). Gunakan `useTemplateRef()` atau elemen `ref()`.
- **Aturan Template:**
  - Tag `v-for` WAJIB memiliki `:key` unik (utamakan ID unik, hindari index array jika data dinamis).
  - DILARANG menggabungkan `v-if` dan `v-for` pada satu tag elemen yang sama.
  - Hindari pemanggilan fungsi berat secara langsung di dalam interpolation template `{{ computeSomething() }}`. Gunakan `computed` property.

## 3. daisyUI & Styling Standards

- **Gunakan Component Class:** Selalu gunakan class komponen resmi daisyUI alih-alih membangun UI dari utility Tailwind mentah.
  - _Contoh Benar:_ `class="btn btn-primary"`, `class="card bg-base-100 shadow-xl"`
  - _Contoh Salah:_ `class="bg-blue-500 text-white px-4 py-2 rounded-lg"`
- **Warna Semantik & Theming:**
  - DILARANG keras menggunakan warna hardcoded Tailwind (`bg-white`, `text-black`, `bg-slate-800`).
  - WAJIB memakai utility warna semantik daisyUI agar fitur `data-theme` berjalan otomatis:
    - Background: `bg-base-100`, `bg-base-200`, `bg-base-300`
    - Teks: `text-base-content`, `text-primary`, `text-secondary`
    - Status / Feedback: `alert-success`, `alert-error`, `alert-warning`, `alert-info`
- **Struktur Form:** Wajib menggunakan pola struktur `form-control` resmi daisyUI:
  ```html
  <div class="form-control w-full">
    <label class="label"><span class="label-text">Label</span></label>
    <input type="text" class="input input-bordered w-full" />
  </div>
  ```
