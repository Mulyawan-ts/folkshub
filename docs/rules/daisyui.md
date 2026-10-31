# DAISYUI & STYLING STANDARDS

## 1. Penggunaan Component Class

- Selalu prioritaskan class komponen resmi daisyUI daripada menyusun utility Tailwind mentah dari nol.
  - _Benar:_ `<button class="btn btn-primary">Submit</button>`
  - _Salah:_ `<button class="bg-blue-600 text-white px-4 py-2 rounded">Submit</button>`

## 2. Warna Semantik & Dark Mode (Theming)

- Tema aktif diatur lewat `document.documentElement.dataset.theme` (`corporate` / `business`), lihat `src/App.vue`.
- DILARANG KERAS menggunakan warna hardcoded Tailwind (seperti `bg-white`, `text-black`, `bg-gray-100`, `text-slate-800`).
- WAJIB menggunakan warna semantik daisyUI:
  - **Background Halaman / Kartu:** `bg-base-100`, `bg-base-200`, `bg-base-300`
  - **Teks:** `text-base-content`, `text-primary`, `text-secondary`, `text-accent`
  - **Tombol / Elemen Aksi:** `btn-primary`, `btn-secondary`, `btn-accent`, `btn-ghost`
  - **Status / Alert:** `alert-success`, `alert-error`, `alert-warning`, `alert-info`
  - **Badge / Status Tag:** `badge-primary`, `badge-outline`, `badge-success`
- **Teks sekunder** DILARANG memakai class karangan seperti `text-text-muted` (tidak terdefinisi di Tailwind/daisyUI dan tidak menghasilkan CSS apa pun). Gunakan `text-base-content/60` atau `opacity-70`.
- Class yang belum terdefinisi di tema WAJIB dicek dulu lewat `src/style.css` sebelum dipakai. Utility Tailwind yang tidak menghasilkan CSS adalah bug senyap.

## 3. Komponen Form

- Elemen form tunggal WAJIB dibungkus struktur `form-control` resmi daisyUI:
  ```html
  <div class="form-control w-full max-w-xs">
    <label class="label">
      <span class="label-text font-medium">Nama Pengguna</span>
    </label>
    <input
      type="text"
      placeholder="Ketik di sini..."
      class="input input-bordered w-full"
    />
    <label class="label" v-if="error">
      <span class="label-text-alt text-error">{{ error }}</span>
    </label>
  </div>
  ```
- **Pengecualian — mode inline/card:** untuk field yang ditampilkan di dalam card (label di kiri, kontrol di kanan), struktur `form-control` bertumpuk diganti layout mendatar: `flex items-center gap-3` dengan label `w-32 shrink-0`. Ini disengaja dan tidak dianggap pelanggaran.

## 4. Komponen Pembungkus Proyek

- Pakai komponen yang sudah ada sebelum menulis markup baru:
  - `AppCard` — wadah card (`padding`, `scroll`, `as`)
  - `AppCardItem` — card satu record
  - `AppInput` — input/select/textarea berlabel, dengan `icon`, `tipe`, `inline`
  - `AppButton`, `AppBadge`, `AppAlert`, `AppConfirm`
  - `DaftarData` — card + pencarian + grid
- DILARANG menyalin ulang markup yang sudah ada di komponen tersebut.

## 5. Konteks Desktop

- Aplikasi berjalan di jendela Wails, bukan browser umum. Hindari elemen yang mengandalkan scroll halaman penuh atau layout `100vh` bertumpuk.
- Dialog konfirmasi WAJIB memakai `AppConfirm`, bukan `window.confirm()`.
