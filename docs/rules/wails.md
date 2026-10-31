# WAILS BINDING & PINIA STORE STANDARDS

Aplikasi ini **desktop (Wails v2, Go + SQLite)**, bukan web app dengan HTTP API.
Tidak ada Axios, tidak ada `import.meta.env.VITE_API_*`, tidak ada `fetch()` ke server.

## 1. Sumber Data

- Semua data diambil lewat binding Go yang di-generate Wails di `frontend/wailsjs/go/main/App`:
  `SimpanX`, `SemuaX`, `AmbilX`, `UpdateX`, `HapusX`.
- DILARANG menulis HTTP client, endpoint REST, atau URL API. Backend menangani SQLite secara lokal.
- Binding adalah hasil generate. DILARANG mengedit `frontend/wailsjs/**` secara manual; jalankan `wails generate module` atau `wails dev` untuk regenerasi.
- Method baru harus dideklarasikan di `app.go` agar ikut ter-bind.
- Setiap payload ke binding WAJIB bertipe jelas di TypeScript (interface di `src/stores/<modul>.ts`).

## 2. Pola Pinia Store

- Proyek memakai **Option Store** (`state` / `getters` / `actions`). Ini konsisten di seluruh 8 modul dan dipertahankan.
- Setiap store mengekspos bentuk yang sama:
  - `state`: `items`, `loading`, `error`
  - `actions`: `fetchAll`, `simpan`, `update`, `hapus`, `bersihkanError`
- Error dari binding masuk ke `error` dengan fallback pesan bahasa Indonesia: `err?.message || "Gagal ..."`. Action yang mengubah data WAJIB `throw err` setelah menyetel `error`, agar pemanggil bisa bereaksi.
- `loading` selalu di-reset di blok `finally`.

  ```typescript
  import { defineStore } from "pinia";
  import { SemuaKaryawan } from "../../wailsjs/go/main/App";

  export const useKaryawanStore = defineStore("karyawan", {
    state: () => ({ items: [] as Karyawan[], loading: false, error: "" }),
    actions: {
      bersihkanError() {
        this.error = "";
      },
      async fetchAll() {
        this.loading = true;
        this.error = "";
        try {
          this.items = await SemuaKaryawan();
        } catch (err: any) {
          this.error = err?.message || "Gagal memuat data karyawan";
        } finally {
          this.loading = false;
        }
      },
    },
  });
  ```

## 3. Sinkronisasi State Lokal

- Setelah `simpan`, dorong hasil kembalian ke `items` tanpa memanggil `fetchAll` ulang bila binding mengembalikan entitas penuh.
- Setelah `hapus`, buang item dari `items` dengan `filter`, jangan panggil ulang seluruh daftar.
- Setelah `update`, ganti item pada indeks hasil `findIndex`.

## 4. Konfigurasi Aplikasi

- Konfigurasi jendela ada di `main.go` (`options.App`), bukan di file env atau JSON frontend.
- Penanganan menu, ukuran minimum jendela, dan shutdown berada di sisi Go.
