// cocokkan mengembalikan true bila salah satu nilai mengandung query,
// tanpa membedakan huruf besar/kecil. Nilai kosong/null dilewati.
export function cocok(
  query: string,
  ...nilai: (string | number | null | undefined)[]
): boolean {
  const q = query.trim().toLowerCase();
  if (!q) return true;
  return nilai.some((v) => v != null && String(v).toLowerCase().includes(q));
}
