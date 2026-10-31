<script setup lang="ts">
import { computed, useId } from "vue";
import type { Size, Variant } from "../types/ui";

const props = withDefaults(
  defineProps<{
    modelValue?: boolean;
    variant?: Variant;
    size?: Size;
    label?: string;
    activeText?: string;
    inactiveText?: string;
    disabled?: boolean;
    required?: boolean;
  }>(),
  {
    modelValue: false,
    variant: "primary",
    size: "md",
    label: "",
    activeText: "",
    inactiveText: "",
    disabled: false,
    required: false,
  },
);

const emit = defineEmits<{
  "update:modelValue": [value: boolean];
  change: [checked: boolean];
}>();

const switchId = useId();

const switchClasses = computed(() => {
  const classes = ["toggle"];
  classes.push(`toggle-${props.variant}`);
  if (props.size !== "md") classes.push(`toggle-${props.size}`);
  return classes.join(" ");
});

function onChange(event: Event) {
  const checked = (event.target as HTMLInputElement).checked;

  emit("update:modelValue", checked);
  emit("change", checked);
}
</script>

<template>
  <div class="flex items-center gap-2">
    <input
      :id="switchId"
      type="checkbox"
      role="switch"
      :class="switchClasses"
      :checked="modelValue"
      :disabled="disabled"
      :required="required"
      @change="onChange"
    />

    <label v-if="label" :for="switchId" class="label cursor-pointer">
      <span class="label-text">{{ label }}</span>
    </label>

    <span v-if="activeText || inactiveText" class="label-text">
      {{ modelValue ? activeText : inactiveText }}
    </span>
  </div>
</template>
