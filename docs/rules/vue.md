# VUE 3 COMPOSITION API & SFC STANDARDS

## 1. Struktur & Urutan SFC

- Urutan blok di dalam file `.vue` WAJIB:
  1. `<script setup lang="ts">`
  2. `<template>`
  3. `<style scoped>`
- Blok `<style>` WAJIB berada di top-level file, DILARANG di dalam `<template>`.

## 2. Skrip & Reaktivitas

- HANYA gunakan `<script setup lang="ts">`. DILARANG menggunakan Options API (`export default`) atau `defineComponent()`.
- State Reaktif:
  - Gunakan `ref()` untuk tipe data primitif dan variabel reaktif umum.
  - Gunakan `reactive()` HANYA untuk objek/array kompleks yang nilainya tidak akan di-reassign secara utuh.
- Compiler Macros:
  - Deklarasikan props secara terstruktur: `defineProps<{ title: string; count?: number }>()`.
  - Deklarasikan emits secara terstruktur dengan tuple: `defineEmits<{ "update:modelValue": [value: string] }>()`.
    Bentuk call-signature lama (`{ (e: 'update', v: string): void }`) tetap sah, tetapi tuple lebih ringkas dan konsisten dengan `v-model`.
  - Gunakan `withDefaults()` jika membutuhkan default value pada props.
  - Props objek/array yang punya default WAJIB memakai factory: `items: () => []`.
- Computeds & Watchers:
  - Gunakan `computed()` untuk semua data turunan.
  - DILARANG menggunakan `watch()` untuk menghitung state sinkron yang bisa ditangani oleh `computed()`.
  - Hapus watcher/event listener manual di `onUnmounted()` jika ada.
- Template Refs:
  - Gunakan `useTemplateRef('elementName')` atau `ref<HTMLElement | null>(null)` untuk referensi elemen DOM. DILARANG menggunakan `document.querySelector` atau `document.getElementById`.

## 3. Aturan Template

- Tag `v-for` WAJIB menggunakan `:key` berupa identifier unik (misal: `:key="item.id"`). DILARANG menggunakan `index` array jika list dapat bertambah, berkurang, atau diurutkan.
- DILARANG menggabungkan `v-if` dan `v-for` pada satu tag elemen yang sama. Gunakan `computed()` untuk memfilter list terlebih dahulu.
- Hindari mengeksekusi fungsi berat di dalam interpolation template `{{ calculateData() }}`. Alihkan ke `computed`.

## 4. Pembagian Peran Komponen

- **Komponen presentasional** (`src/components/`): hanya menerima props, memancarkan emit, dan menaruh konten lewat slot. DILARANG meng-import Pinia store atau memanggil binding Wails langsung.
- **Halaman container** (`src/pages/`): memegang store, menjalankan action, dan meneruskan data ke komponen lewat props serta menangani emit.
- Handler DOM internal komponen WAJIB disalurkan sebagai emit, bukan dibocorkan ke pemanggil lewat attribute mentah.

## 5. Atribut Fallthrough

- Attribute yang tidak dideklarasikan sebagai prop (mis. `class`, `@submit`, `id`) diteruskan Vue ke elemen root komponen.
- Komponen pembungkus dengan lebih dari satu elemen root WAJIB memakai `inheritAttrs: false` + `v-bind="$attrs"` pada elemen yang dituju, atau didokumentasikan bahwa fallthrough berlaku ke root tunggal.
