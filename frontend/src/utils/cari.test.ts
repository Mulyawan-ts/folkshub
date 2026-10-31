import { test } from "node:test";
import assert from "node:assert/strict";
import { cocok } from "./cari.ts";

test("cocok: query kosong selalu cocok", () => {
  assert.equal(cocok("", "apa saja"), true);
  assert.equal(cocok("   ", "apa saja"), true);
});

test("cocok: cocok sebagian tanpa beda huruf besar", () => {
  assert.equal(cocok("bu", "Budi", "budi@contoh.com"), true);
  assert.equal(cocok("CONTOH", "Budi", "budi@contoh.com"), true);
  assert.equal(cocok("12", 1234), true);
});

test("cocok: nilai null/kosong dilewati", () => {
  assert.equal(cocok("x", null, undefined, "", "axb"), true);
  assert.equal(cocok("z", null, undefined, ""), false);
});
