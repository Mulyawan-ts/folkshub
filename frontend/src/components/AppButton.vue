<script setup lang="ts">
import { computed } from "vue";
import type { Size, Variant } from "@/types/ui";

type ButtonVariant = Variant | "ghost";

const props = withDefaults(
  defineProps<{
    variant?: ButtonVariant;
    size?: Size;
    type?: "button" | "submit" | "reset";
    outline?: boolean;
    loading?: boolean;
    disabled?: boolean;
  }>(),
  {
    variant: "primary",
    size: "md",
    type: "button",
    outline: false,
    loading: false,
    disabled: false,
  },
);

const emit = defineEmits<{
  click: [event: MouseEvent];
}>();

const buttonClasses = computed(() => {
  const classes = ["btn"];

  if (props.outline) classes.push("btn-outline");
  classes.push(`btn-${props.variant}`);

  if (props.size !== "md") classes.push(`btn-${props.size}`);

  return classes.join(" ");
});
</script>

<template>
  <button
    :type="type"
    :class="buttonClasses"
    :disabled="disabled || loading"
    @click="emit('click', $event)"
  >
    <span v-if="loading" class="loading loading-spinner loading-xs"></span>
    <slot />
  </button>
</template>
