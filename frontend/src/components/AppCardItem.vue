<script setup lang="ts">
import AppCard from "./AppCard.vue";
import AppButton from "./AppButton.vue";
import AppBadge from "./AppBadge.vue";
import AppInput from "./AppInput.vue";
import type { Variant } from "../types/ui";

withDefaults(
  defineProps<{
    judul: string;
    subjudul?: string;
    status?: string;
    statusVariant?: Variant;
    fields?: { label: string; value: string | number; icon?: string }[];
    hapusLabel?: string;
    editable?: boolean;
  }>(),
  {
    subjudul: "",
    status: "",
    statusVariant: "neutral",
    fields: () => [],
    hapusLabel: "Hapus",
    editable: false,
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
        <p v-if="subjudul" class="text-sm text-base-content/60 truncate">
          {{ subjudul }}
        </p>
      </div>
      <AppBadge v-if="status" :variant="statusVariant" size="sm">
        {{ status }}
      </AppBadge>
    </div>

    <div v-if="editable" class="flex flex-col gap-2">
      <slot name="edit" />
    </div>

    <div v-else-if="fields.length" class="flex flex-col gap-2">
      <AppInput
        v-for="f in fields"
        :key="f.label"
        :label="f.label"
        :icon="f.icon"
        :model-value="f.value"
        inline
        readonly
      />
    </div>

    <div class="mt-auto flex justify-end gap-2">
      <slot name="aksi" />
      <AppButton variant="error" size="sm" outline @click="$emit('hapus')">
        {{ hapusLabel }}
      </AppButton>
    </div>
  </AppCard>
</template>
