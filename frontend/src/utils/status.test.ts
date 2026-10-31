import { test } from "node:test";
import assert from "node:assert/strict";
import { statusVariant, ABSENSI_STATUS, GAJI_STATUS } from "./status.ts";

test("statusVariant: memetakan status ke varian", () => {
  assert.equal(statusVariant("hadir", ABSENSI_STATUS), "success");
  assert.equal(statusVariant("alfa", ABSENSI_STATUS), "error");
  assert.equal(statusVariant("dibayar", GAJI_STATUS), "success");
});

test("statusVariant: status tak dikenal jadi neutral", () => {
  assert.equal(statusVariant("entah", ABSENSI_STATUS), "neutral");
  assert.equal(statusVariant("", GAJI_STATUS), "neutral");
});
