const HARI = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu']

// formatTanggal menampilkan tanggal simpanan (YYYY-MM-DD) sebagai
// "Hari, tanggal-bulan-tahun", contoh: "2026-11-10" -> "Selasa, 10-11-2026".
// Nilai yang kosong atau bukan tanggal kalender nyata dikembalikan apa adanya
// (backend menolak yang mustahil).
export function formatTanggal(iso: string): string {
  const m = /^(\d{4})-(\d{2})-(\d{2})$/.exec(iso)
  if (!m) return iso
  const [, y, mo, d] = m
  const date = new Date(Date.UTC(Number(y), Number(mo) - 1, Number(d)))
  if (Number.isNaN(date.getTime())) return iso
  if (date.getUTCMonth() !== Number(mo) - 1 || date.getUTCDate() !== Number(d)) return iso
  return `${HARI[date.getUTCDay()]}, ${d}-${mo}-${y}`
}
