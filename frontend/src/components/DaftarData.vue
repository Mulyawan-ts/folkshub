<script setup lang="ts">
import CariInput from "./CariInput.vue";
import AppCard from "./AppCard.vue";

withDefaults(
  defineProps<{
    cari?: string;
    placeholder?: string;
    kosong?: string;
    kosongCari?: string;
    jumlah?: number;
  }>(),
  {
    cari: "",
    placeholder: "Cari...",
    kosong: "Belum ada data",
    kosongCari: "Tidak ada hasil yang cocok",
    jumlah: 0,
  },
);

defineEmits<{
  "update:cari": [value: string];
}>();
</script>

<template>
  <AppCard>
    <CariInput
      :model-value="cari"
      :placeholder="placeholder"
      @update:model-value="$emit('update:cari', $event)"
    />

    <div
      v-if="jumlah > 0"
      class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4"
    >
      <slot />
    </div>

    <p v-else class="text-center text-text-muted py-8">
      {{ cari ? kosongCari : kosong }}
    </p>
  </AppCard>
</template>
