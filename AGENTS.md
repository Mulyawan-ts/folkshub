# AGENT RULES & INSTRUCTIONS

## 0. Konteks Proyek

**Folkshub adalah aplikasi desktop HRIS (HR Management System).**
Stack: **Vue 3 + TypeScript + Pinia + daisyUI**, dibungkus **Wails v2 (Go + SQLite)**.

- Bukan web app. Tidak ada server, tidak ada REST API, tidak ada HTTP client.
- Data disimpan per modul dalam file SQLite (`karyawan.db`, `absensi.db`, dst) lewat service Go.
- Semua komunikasi frontend ↔ backend memakai binding Wails di `frontend/wailsjs/go/main/App`.
- Setiap perubahan atau aksi di UI berbasis **card dan button**, bukan halaman web panjang.

## 1. Aturan Dasar (Anti-Slop)

- Langsung berikan kode utuh tanpa basa-basi pembuka/penutup.
- DILARANG menggunakan komentar AI slop yang tidak perlu.
- Terapkan prinsip YAGNI (HANYA ubah file yang relevan).
- DILARANG menambah dependensi baru jika stdlib, API native, atau dependensi yang sudah terpasang bisa menyelesaikannya.

## 2. Modul Aturan Spesifik (BACA SAAT TERKAIT)

Sebelum membuat atau mengubah kode, WAJIB membaca file aturan modular di `docs/rules/`:

- **Vue 3 SFC:** jika membuat atau mengedit komponen `.vue`, BACA `docs/rules/vue.md`.
- **daisyUI & Styling:** jika membuat layout, form, atau komponen UI, BACA `docs/rules/daisyui.md`.
- **State & Binding Wails:** jika mengelola Pinia store atau memanggil backend, BACA `docs/rules/wails.md`.

## 3. Eksekusi

- Pastikan kode yang dihasilkan sesuai dengan standar pada file aturan modular di atas.
- Urutan verifikasi wajib sebelum menyatakan selesai:
  - Frontend: `cd frontend && pnpm run type-check`, `pnpm test`, `pnpm exec prettier --check src`
  - Backend: `go build ./...`, `go vet ./...`, `go test ./...`

## 4. Incremental & Single-File Execution (1 File Per Tahap)

- **Fokus 1 File:** DILARANG MERUBAH ATAU MEMBUAT BANYAK FILE SEKALIGUS dalam satu langkah/respon.
- **Alur Kerja Bertahap:**
  1. HANYA buat atau sesuaikan 1 file pada satu waktu.
  2. Pastikan file tersebut bebas dari syntax error atau broken import.
  3. Konfirmasikan secara singkat file yang telah diubah sebelum beralih ke file berikutnya.
- **Urutan Dependencies (Bottom-Up):** Jika tugas membutuhkan beberapa file, kerjakan dengan urutan:
  1. File Type/Interface (`types/`)
  2. File API Service (`services/`)
  3. File State (`stores/`)
  4. File Komponen UI (`components/` atau `views/`)
