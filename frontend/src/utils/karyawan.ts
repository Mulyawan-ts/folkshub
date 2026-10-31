import type { OpsiInput } from "@/components/AppInput.vue";

// karyawanOptions mengubah daftar karyawan menjadi opsi <select> dengan value id,
// sehingga form tidak perlu meminta user menghafal ID.
export function karyawanOptions(
  daftar: { id: number; nama: string }[],
): OpsiInput[] {
  return daftar.map((karyawan) => ({
    value: karyawan.id,
    label: karyawan.nama,
  }));
}
