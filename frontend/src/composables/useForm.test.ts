import { test } from "node:test";
import assert from "node:assert/strict";
import { useForm } from "./useForm.ts";

function buatToast() {
  const dicatat: { tipe: string; pesan: string }[] = [];
  return {
    dicatat,
    error: (pesan: string) => dicatat.push({ tipe: "error", pesan }),
    success: (pesan: string) => dicatat.push({ tipe: "success", pesan }),
  };
}

test("useForm: validasi gagal tidak memanggil aksi", async () => {
  const toast = buatToast();
  let dipanggil = 0;
  const f = useForm({
    awal: { nama: "" },
    validasi: (v) => (v.nama.trim() ? "" : "Nama wajib diisi"),
    aksi: async () => {
      dipanggil++;
    },
    pesanSukses: "ok",
    toast,
  });

  await f.simpan();

  assert.equal(dipanggil, 0);
  assert.deepEqual(toast.dicatat, [
    { tipe: "error", pesan: "Nama wajib diisi" },
  ]);
});

test("useForm: simpan sukses memberi notifikasi dan mereset field", async () => {
  const toast = buatToast();
  const terkirim: { nama: string }[] = [];
  const f = useForm({
    awal: { nama: "" },
    validasi: () => "",
    aksi: async (v) => {
      terkirim.push({ nama: v.nama });
    },
    pesanSukses: "Karyawan disimpan",
    toast,
  });

  f.nilai.nama = "Budi";
  await f.simpan();

  assert.deepEqual(terkirim, [{ nama: "Budi" }]);
  assert.equal(f.nilai.nama, "");
  assert.equal(f.sedangSimpan.value, false);
  assert.deepEqual(toast.dicatat, [
    { tipe: "success", pesan: "Karyawan disimpan" },
  ]);
});

test("useForm: aksi gagal tetap reset status simpan tanpa mereset field", async () => {
  const toast = buatToast();
  const f = useForm({
    awal: { nama: "" },
    validasi: () => "",
    aksi: async () => {
      throw new Error("backend menolak");
    },
    pesanSukses: "tidak dipakai",
    toast,
  });

  f.nilai.nama = "Ani";
  await f.simpan();

  assert.equal(f.sedangSimpan.value, false);
  assert.equal(f.nilai.nama, "Ani");
  assert.deepEqual(toast.dicatat, []);
});
