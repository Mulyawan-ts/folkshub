import { test } from "node:test";
import assert from "node:assert/strict";
import { karyawanOptions } from "./karyawan.ts";

test("karyawanOptions: value adalah id, label adalah nama", () => {
  assert.deepEqual(
    karyawanOptions([
      { id: 3, nama: "Budi" },
      { id: 7, nama: "Ani" },
    ]),
    [
      { value: 3, label: "Budi" },
      { value: 7, label: "Ani" },
    ],
  );
});

test("karyawanOptions: daftar kosong menghasilkan array kosong", () => {
  assert.deepEqual(karyawanOptions([]), []);
});
