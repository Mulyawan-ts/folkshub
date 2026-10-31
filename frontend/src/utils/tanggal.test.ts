import { test } from "node:test";
import assert from "node:assert/strict";
import { formatTanggal, formatBulan } from "./tanggal.ts";

test("formatTanggal: hari, dd-mm-yyyy", () => {
  assert.equal(formatTanggal("2026-11-10"), "Selasa, 10-11-2026");
  assert.equal(formatTanggal("2026-11-01"), "Minggu, 01-11-2026");
  assert.equal(formatTanggal("2026-01-01"), "Kamis, 01-01-2026");
  assert.equal(formatTanggal("2024-02-29"), "Kamis, 29-02-2024");
});

test("formatTanggal: nilai tak valid dikembalikan apa adanya", () => {
  assert.equal(formatTanggal("2026-02-30"), "2026-02-30");
  assert.equal(formatTanggal("2026-13-01"), "2026-13-01");
  assert.equal(formatTanggal("2025-02-29"), "2025-02-29");
  assert.equal(formatTanggal("10-11-2026"), "10-11-2026");
  assert.equal(formatTanggal(""), "");
  assert.equal(formatTanggal("x"), "x");
});

test("formatBulan: nama bulan Indonesia", () => {
  assert.equal(formatBulan("2026-11"), "November 2026");
  assert.equal(formatBulan("2026-01"), "Januari 2026");
  assert.equal(formatBulan("2026-12"), "Desember 2026");
});

test("formatBulan: nilai tak valid dikembalikan apa adanya", () => {
  assert.equal(formatBulan("2026-13"), "2026-13");
  assert.equal(formatBulan("2026-00"), "2026-00");
  assert.equal(formatBulan("2026-11-01"), "2026-11-01");
  assert.equal(formatBulan(""), "");
});
