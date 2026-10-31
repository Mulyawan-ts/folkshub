# ANTI-AI-SLOP & STRICT CODING RULES

## 1. Gaya Komunikasi (No Fluff / No Slop)

- Langsung ke solusi. DILARANG menggunakan kalimat pembuka basa-basi seperti "Tentu!", "Berikut kodenya:", atau "Sebagai AI...".
- Jangan mengulang penjelasan dari apa yang baru saja Anda edit. Berikan ringkasan teknis singkat maksimal 2 kalimat jika diperlukan.
- Jangan gunakan komentar kode AI Slop yang tidak berguna (contoh dilarang: `// variabel nama`, `// fungsi untuk login`). Hanya tulis komentar untuk logika bisnis yang sangat kompleks.

## 2. Kualitas & Lingkup Kode (Clean Code)

- Prinsip YAGNI (You Ain't Gonna Need It): HANYA ubah atau tambahkan kode yang diminta. DILARANG menambahkan abstraksi, dependensi baru, atau refactoring pada file yang tidak berhubungan.
- DILARANG menggunakan kode _placeholder_ atau potongan menggantung seperti `// ... tambahkan sisa kode di sini ...`. Tuliskan kode secara utuh dan siap jalan.
- Selalu tangani error handling dengan benar, jangan gunakan `try-catch` kosong atau sekadar `console.log(err)`.

## 3. Eksekusi File & Terminal

- Jangan jalankan perintah terminal yang berisiko merusak (seperti `git reset --hard` atau `rm -rf`) tanpa persetujuan eksplisit.
- Sebelum mengedit file, pastikan Anda telah membaca konteks file terkait terlebih dahulu.
