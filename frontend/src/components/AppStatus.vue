<script setup lang="ts">
import { computed } from "vue";
import AppBadge from "./AppBadge.vue";
import { statusVariant } from "@/utils/status";
import type { Size, Variant } from "@/types/ui";

const props = withDefaults(
  defineProps<{
    status: string;
    map?: Record<string, Variant>;
    size?: Size;
    label?: string;
  }>(),
  {
    map: () => ({}),
    size: "sm",
    label: "",
  },
);

const variant = computed(() =>
  props.map && Object.keys(props.map).length
    ? statusVariant(props.status, props.map)
    : "neutral",
);
</script>

<template>
  <AppBadge :variant="variant" :size="size">
    {{ label || status }}
  </AppBadge>
</template>
