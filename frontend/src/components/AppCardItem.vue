<script setup lang="ts">
import AppCard from "./AppCard.vue";
import AppButton from "./AppButton.vue";
import AppBadge from "./AppBadge.vue";
import type { Variant } from "../types/ui";

withDefaults(
  defineProps<{
    judul: string;
    subjudul?: string;
    status?: string;
    statusVariant?: Variant;
    fields?: { label: string; value: string | number }[];
    hapusLabel?: string;
  }>(),
  {
    subjudul: "",
    status: "",
    statusVariant: "neutral",
    fields: () => [],
    hapusLabel: "Hapus",
  },
);

defineEmits<{
  hapus: [];
}>();
</script>

<template>
  <AppCard padding="4" class="flex flex-col gap-3">
    <div class="flex items-start justify-between gap-2">
      <div class="min-w-0">
        <p class="font-semibold truncate">{{ judul }}</p>
        <p v-if="subjudul" class="text-sm text-text-muted truncate">
          {{ subjudul }}
        </p>
      </div>
      <AppBadge v-if="status" :variant="statusVariant" size="sm">
        {{ status }}
      </AppBadge>
    </div>

    <dl v-if="fields.length" class="grid grid-cols-2 gap-x-3 gap-y-1 text-sm">
      <template v-for="f in fields" :key="f.label">
        <dt class="text-text-muted">{{ f.label }}</dt>
        <dd class="text-right">{{ f.value }}</dd>
      </template>
    </dl>

    <div class="mt-auto flex justify-end">
      <AppButton variant="error" size="sm" outline @click="$emit('hapus')">
        {{ hapusLabel }}
      </AppButton>
    </div>
  </AppCard>
</template>
