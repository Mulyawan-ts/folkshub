import { reactive, ref } from "vue";

interface OpsiForm<T extends Record<string, string>> {
  awal: T;
  validasi: (nilai: T) => string;
  aksi: (nilai: T) => Promise<unknown>;
  pesanSukses: string;
  toast: {
    error(message: string): unknown;
    success(message: string): unknown;
  };
}

// useForm menyatukan alur form halaman: simpan, validasi, notifikasi, dan reset.
// Semua field disimpan sebagai string karena berasal dari input; konversi
// tipe (Number, dll.) dilakukan di opsi aksi.
export function useForm<T extends Record<string, string>>(opsi: OpsiForm<T>) {
  const nilai = reactive({ ...opsi.awal }) as T;
  const sedangSimpan = ref(false);

  function reset() {
    Object.assign(nilai, opsi.awal);
  }

  async function simpan() {
    const pesan = opsi.validasi(nilai);
    if (pesan) {
      opsi.toast.error(pesan);
      return;
    }

    sedangSimpan.value = true;
    try {
      await opsi.aksi(nilai);
      opsi.toast.success(opsi.pesanSukses);
      reset();
    } catch {
      // error sudah disetel oleh store
    } finally {
      sedangSimpan.value = false;
    }
  }

  return { nilai, sedangSimpan, simpan, reset };
}
