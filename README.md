# folkshub

Aplikasi desktop HRIS (HR Management System) berbasis **Wails v2** dengan frontend **Vue 3 + TypeScript + Pinia + daisyUI** dan backend **Go + SQLite**.

Bukan aplikasi web. Tidak ada server, tidak ada REST API, tidak ada HTTP client. Frontend dan backend berkomunikasi lewat binding Wails yang di-generate ke `frontend/wailsjs/`.

## Fitur

- **Karyawan**: CRUD data karyawan
- **Absensi**: absensi masuk/keluar per karyawan
- **Cuti**: pengajuan dan persetujuan cuti
- **Reimburse**: pengajuan pengeluaran dan bukti
- **Gaji**: perhitungan gaji pokok, tunjangan, potongan
- **Jabatan**: master jabatan dengan gaji pokok & tunjangan
- **Departemen**: master departemen dan kepala
- **Pelatihan**: pelatihan karyawan (jenis, tanggal, lokasi, biaya, sertifikat, status)

## Struktur

```
app.go                  # Titik masuk Wails, pendaftaran service & binding endpoint
main.go                 # Konfigurasi jendela (ukuran, judul) dan inisialisasi Wails
go.mod / go.sum         # Dependensi Go
internal/
  model/                # Struct domain + TableName()
  service/              # Service per modul (CRUD SQLite + validasi)
frontend/
  src/
    components/          # Komponen UI presentasional (props + emit)
    composables/         # Logika bersama antar halaman (mis. useForm)
    pages/               # Halaman container, pemegang store
    stores/              # Pinia store per modul
    utils/               # Helper murni + test (tanggal, pencarian, status)
    types/               # Tipe UI bersama
  wailsjs/               # Binding Go hasil generate, JANGAN diedit manual
build/                   # Hasil build per platform
docs/rules/              # Aturan pengembangan (dibaca agent sebelum menulis kode)
```

## Persiapan

- [Go](https://go.dev/) 1.25+
- [Node.js](https://nodejs.org/) 18+
- [pnpm](https://pnpm.io/) (project memakai `pnpm-lock.yaml`)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation) v2

## Menjalankan

Backend (Go):

```sh
go build ./...
go vet ./...
go test ./...
```

Frontend:

```sh
cd frontend
pnpm install
pnpm run type-check
pnpm test
pnpm run format:check
```

Menjalankan aplikasi dalam mode development (Wails menjalankan Vite otomatis):

```sh
wails dev
```

Build produksi:

```sh
wails build
```

Hasil build tersedia di `build/bin` sesuai platform.

## Data

Setiap modul menyimpan datanya di file SQLite terpisah pada direktori aplikasi:

- `karyawan.db`, `absensi.db`, `cuti.db`, `reimburse.db`
- `gaji.db`, `jabatan.db`, `departemen.db`, `pelatihan.db`

File `*.db`, `*.db-shm`, dan `*.db-wal` tidak masuk repository (lihat `.gitignore`). Mode WAL aktif, jadi file `-wal` dapat muncul selama aplikasi berjalan.

## Lisensi

Tidak ditetapkan.
