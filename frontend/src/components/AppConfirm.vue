<script setup lang="ts">
import { onUnmounted, watch } from "vue";
import AppButton from "./AppButton.vue";
import type { Variant } from "@/types/ui";

const props = withDefaults(
  defineProps<{
    terbuka: boolean;
    judul?: string;
    pesan?: string;
    confirmLabel?: string;
    batalLabel?: string;
    variant?: Variant;
  }>(),
  {
    judul: "Konfirmasi",
    pesan: "",
    confirmLabel: "Ya",
    batalLabel: "Batal",
    variant: "primary",
  },
);

const emit = defineEmits<{
  konfirmasi: [];
  batal: [];
}>();

function tutup() {
  emit("batal");
}

function onKeydown(event: KeyboardEvent) {
  if (event.key === "Escape") tutup();
}

watch(
  () => props.terbuka,
  (buka) => {
    if (buka) window.addEventListener("keydown", onKeydown);
    else window.removeEventListener("keydown", onKeydown);
  },
);

onUnmounted(() => window.removeEventListener("keydown", onKeydown));
</script>

<template>
  <div v-if="terbuka" class="modal modal-open" role="dialog" aria-modal="true">
    <div class="modal-box">
      <h3 class="font-bold text-lg">{{ judul }}</h3>

      <p v-if="pesan" class="py-4">{{ pesan }}</p>
      <div v-else-if="$slots.default" class="py-4">
        <slot />
      </div>

      <div class="modal-action">
        <AppButton variant="ghost" @click="tutup">{{ batalLabel }}</AppButton>
        <AppButton :variant="variant" @click="emit('konfirmasi')">
          {{ confirmLabel }}
        </AppButton>
      </div>
    </div>

    <div class="modal-backdrop" @click="tutup"></div>
  </div>
</template>
