<script setup lang="ts">
import { computed } from "vue";
import type { Padding } from "../types/ui";

const props = withDefaults(
  defineProps<{
    as?: string;
    padding?: Padding;
    scroll?: boolean;
  }>(),
  {
    as: "div",
    padding: "4",
    scroll: false,
  },
);

const cardClasses = computed(() => {
  const classes = ["card", "bg-base-100", "shadow-sm"];

  if (props.padding !== "0") classes.push(`p-${props.padding}`);
  if (props.scroll) classes.push("overflow-x-auto");

  return classes.join(" ");
});
</script>

<template>
  <component :is="as" :class="cardClasses">
    <slot />
  </component>
</template>
