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

## 6. Pola yang Sering Membingungkan (Bacaan Wajib untuk Pemula)

### 6.1 Kapan menulis `.value`

`ref()` dan `reactive()` menghasilkan objek pembungkus. Aksesnya berbeda tergantung tempat:

| Lokasi | Cara akses | Contoh |
| --- | --- | --- |
| `<script setup>` | wajib `.value` | `cari.value = "budi"` |
| `<template>` | tanpa `.value` | `v-model="cari"` |
| di dalam `computed()` | wajib `.value` | `items.value.filter(...)` |

Aturan ini berlaku juga untuk nilai hasil `useForm`: di script tulis `form.sedangSimpan.value`, di template tulis `form.sedangSimpan`.

### 6.2 Kenapa `useForm` menerima `toast`

`useForm` adalah composable, bukan komponen. Ia tidak boleh meng-import store sendiri, karena store hanya boleh disentuh oleh halaman container (lihat §4). Karena itu dependensi `toast` dioper sebagai parameter:

```ts
const form = useForm({
  awal: { nama: "" },
  validasi: (nilai) => (nilai.nama.trim() ? "" : "Nama wajib diisi"),
  aksi: (nilai) => karyawanStore.simpan({ nama: nilai.nama, email: "", status: "tetap" }),
  pesanSukses: "Karyawan berhasil disimpan",
  toast: toastStore,
});
```

Alur di dalamnya: `validasi` dijalankan dulu; bila mengembalikan teks, proses berhenti dan `toast.error` dipanggil. Bila lolos, `aksi` dijalankan, lalu `toast.success`, lalu semua field kembali ke `awal`.

### 6.3 Kenapa `AppLabel` memakai prop `untuk`, bukan `for`

`for` adalah kata kunci JavaScript. Prop dengan nama itu gagal saat diakses di template. `untuk` adalah padanan Indonesia dan dipakai sebagai nilai atribut `for` di HTML.

Tanpa prop `untuk`, `AppLabel` merender `<span class="label">`, bukan `<label>`. Alasannya: `<label for="...">` hanya sah bila `for` menunjuk id elemen yang benar-benar ada. Karena `<input>` berada di luar komponen, id-nya tidak bisa diketahui otomatis, jadi label sengaja dibuat pasif.

### 6.4 Kenapa alias `@` tidak berlaku untuk `wailsjs`

Alias `@` dipetakan ke `src/` (lihat `vite.config.ts` dan `paths` di `tsconfig.json`). Folder `wailsjs/` berada di luar `src/`, sehingga import binding tetap memakai path relatif:

```ts
import { SemuaKaryawan } from "../../wailsjs/go/main/App";
```

### 6.5 Kenapa `catch` di halaman tampak kosong

```ts
try {
  await karyawanStore.hapus(id);
  toastStore.success("Karyawan berhasil dihapus");
} catch {
  // error sudah disetel oleh store
}
```

Store sudah menyimpan pesan ke `error` sebelum melempar ulang. Halaman menampilkan `error` itu lewat `AppAlert`. Jadi `catch` kosong di sini bukan penelantaran error, melainkan titik henti agar toast sukses tidak muncul saat gagal.

### 6.6 Kenapa `:key` wajib memakai `id`

```html
<tr v-for="karyawan in filteredItems" :key="karyawan.id">
```

Memakai indeks array membuat Vue salah mengenali baris saat data dihapus atau diurutkan, sehingga tampilan bisa tidak sesuai data. `id` dari database selalu unik dan stabil.

### 6.7 Kenapa variabel ditulis penuh

DILARANG memakai satu huruf untuk data domain (`a`, `c`, `g`, `k`). Gunakan nama entitasnya: `absensi`, `cuti`, `gaji`, `karyawan`. Singkatan hanya boleh untuk indeks (`i`) atau akumulator numerik pendek di lingkup yang sangat sempit.

