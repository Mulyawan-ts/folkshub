import type { Variant } from "../types/ui";

// statusVariant memetakan status domain ke varian AppBadge.
// Status yang tidak dikenal dikembalikan sebagai "neutral".
export function statusVariant(
  status: string,
  map: Record<string, Variant>,
): Variant {
  return map[status] ?? "neutral";
}

export const KARYAWAN_STATUS: Record<string, Variant> = {
  tetap: "success",
  outsourcing: "warning",
  magang: "info",
};

export const ABSENSI_STATUS: Record<string, Variant> = {
  hadir: "success",
  izin: "info",
  sakit: "warning",
  alfa: "error",
  lembur: "neutral",
};

export const CUTI_STATUS: Record<string, Variant> = {
  draft: "neutral",
  submitted: "warning",
  approved: "success",
  rejected: "error",
  cancelled: "neutral",
};

export const REIMBURSE_STATUS: Record<string, Variant> = {
  draft: "neutral",
  submitted: "warning",
  approved: "success",
  rejected: "error",
  paid: "success",
};

export const GAJI_STATUS: Record<string, Variant> = {
  draft: "neutral",
  diproses: "warning",
  dibayar: "success",
  ditolak: "error",
};

export const PELATIHAN_STATUS: Record<string, Variant> = {
  terjadwal: "neutral",
  berlangsung: "warning",
  selesai: "success",
  dibatalkan: "error",
};
