<script setup lang="ts">
import { computed, useId } from "vue";
import type { Size, Variant } from "../types/ui";

const props = withDefaults(
  defineProps<{
    modelValue?: boolean;
    variant?: Variant;
    size?: Size;
    label?: string;
    disabled?: boolean;
    required?: boolean;
    indeterminate?: boolean;
  }>(),
  {
    modelValue: false,
    variant: "primary",
    size: "md",
    label: "",
    disabled: false,
    required: false,
    indeterminate: false,
  },
);

const emit = defineEmits<{
  "update:modelValue": [value: boolean];
  change: [checked: boolean];
}>();

const checkboxId = useId();

const checkboxClasses = computed(() => {
  const classes = ["checkbox"];
  classes.push(`checkbox-${props.variant}`);
  if (props.size !== "md") classes.push(`checkbox-${props.size}`);
  return classes.join(" ");
});

function onChange(event: Event) {
  const checked = (event.target as HTMLInputElement).checked;

  emit("update:modelValue", checked);
  emit("change", checked);
}
</script>

<template>
  <div class="flex items-start gap-2">
    <input
      :id="checkboxId"
      type="checkbox"
      :class="checkboxClasses"
      :checked="modelValue"
      :disabled="disabled"
      :required="required"
      :indeterminate="indeterminate"
      @change="onChange"
    />

    <label v-if="label" :for="checkboxId" class="label cursor-pointer">
      <span class="label-text">{{ label }}</span>
    </label>
  </div>
</template>
