# folkshub

Aplikasi desktop manajemen karyawan berbasis Wails (Go + Vue 3 + Pinia). Data disimpan dalam file JSON per modul di direktori aplikasi.

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
app.go              # Titik masuk Wails, pendaftaran semua service & binding endpoint
main.go             # Inisialisasi aplikasi Wails
go.mod / go.sum     # Dependensi Go
internal/
  model/            # Struct domain + TableName()
  service/          # Service per modul (CRUD berbasis JSON file)
frontend/
  src/
    stores/         # Pinia store per modul
    wailsjs/        # Binding Go yang di-generate ke TypeScript
wails.json          # Konfigurasi project Wails
build/              # Hasil build per platform
```

## Persiapan

- [Go](https://go.dev/) (untuk backend)
- [Node.js](https://nodejs.org/) + npm (untuk frontend)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation) (opsional, untuk `wails dev`/`wails build`)

## Menjalankan

Backend (Go):

```sh
go build ./...
go vet ./...
go test ./...
```

Frontend (type-check):

```sh
cd frontend
npm install
npm run type-check
```

Menjalankan dalam mode development:

```sh
wails dev
```

Build produksi:

```sh
wails build
```

Hasil build tersedia di `build/bin`, `build/darwin`, atau `build/windows` sesuai platform.

## Data

Setiap modul menyimpan data ke file JSON sendiri di direktori aplikasi:

- `karyawan.db`, `absensi.db`, `cuti.db`, `reimburse.db`
- `gaji.db`, `jabatan.db`, `departemen.db`, `pelatihan.db`

## Lisensi

Tidak ditetapkan.
