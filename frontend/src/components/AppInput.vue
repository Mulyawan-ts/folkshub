<script setup lang="ts">
import { computed, useId } from "vue";
import type { Size } from "../types/ui";

export type TipeInput = "text" | "select" | "textarea";

export interface OpsiInput {
  value: string | number;
  label: string;
}

const props = withDefaults(
  defineProps<{
    modelValue?: string | number;
    tipe?: TipeInput;
    type?: string;
    variant?: "default" | "filled" | "ghost";
    size?: Size;
    placeholder?: string;
    label?: string;
    hint?: string;
    error?: string;
    icon?: string;
    options?: OpsiInput[];
    inline?: boolean;
    rows?: number;
    disabled?: boolean;
    readonly?: boolean;
    required?: boolean;
    maxlength?: number;
  }>(),
  {
    modelValue: "",
    tipe: "text",
    type: "text",
    variant: "default",
    size: "md",
    placeholder: "",
    label: "",
    hint: "",
    error: "",
    icon: "",
    options: () => [],
    inline: false,
    rows: 3,
    disabled: false,
    readonly: false,
    required: false,
    maxlength: undefined,
  },
);

const emit = defineEmits<{
  "update:modelValue": [value: string | number];
  enter: [value: string | number];
}>();

const inputId = useId();

const hasError = computed(() => props.error !== "");

const controlClasses = computed(() => {
  const base = props.tipe === "select" ? "select" : "input";

  const classes =
    props.tipe === "textarea" ? ["textarea", "w-full"] : [base, "w-full"];

  if (props.variant === "filled") classes.push(`${base}-filled`);
  if (props.variant === "ghost") classes.push(`${base}-ghost`);
  if (props.size !== "md") classes.push(`${base}-${props.size}`);
  if (hasError.value) classes.push(`${base}-error`);

  return classes.join(" ");
});

const wrapperClasses = computed(() =>
  props.inline ? "flex items-center gap-3 w-full" : "form-control w-full",
);

const labelClasses = computed(() =>
  props.inline ? "w-32 shrink-0 text-sm text-base-content/60" : "label",
);

function onInput(event: Event) {
  emit("update:modelValue", (event.target as HTMLInputElement).value);
}

function onKeydown(event: KeyboardEvent) {
  if (event.key === "Enter")
    emit("enter", (event.target as HTMLInputElement).value);
}
</script>

<template>
  <div :class="wrapperClasses">
    <label v-if="label" :for="inputId" :class="labelClasses">
      <span class="label-text">{{ label }}</span>
      <span v-if="required" class="text-error">*</span>
    </label>

    <div class="flex w-full items-center gap-2">
      <svg
        v-if="icon"
        class="size-4 shrink-0 text-base-content/60"
        fill="none"
        stroke-width="1.5"
        stroke="currentColor"
        viewBox="0 0 24 24"
        aria-hidden="true"
      >
        <path stroke-linecap="round" stroke-linejoin="round" :d="icon" />
      </svg>

      <select
        v-if="tipe === 'select'"
        :id="inputId"
        :class="controlClasses"
        :value="modelValue"
        :disabled="disabled"
        :required="required"
        @change="onInput"
      >
        <option value="" disabled v-if="placeholder">{{ placeholder }}</option>
        <option v-for="opt in options" :key="opt.value" :value="opt.value">
          {{ opt.label }}
        </option>
      </select>

      <textarea
        v-else-if="tipe === 'textarea'"
        :id="inputId"
        :class="controlClasses"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readonly"
        :required="required"
        :maxlength="maxlength"
        :rows="rows"
        @input="onInput"
      ></textarea>

      <input
        v-else
        :id="inputId"
        :type="type"
        :class="controlClasses"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readonly"
        :required="required"
        :maxlength="maxlength"
        @input="onInput"
        @keydown="onKeydown"
      />
    </div>

    <div v-if="error || hint" class="label">
      <span v-if="error" class="label-text text-error">{{ error }}</span>
      <span v-else-if="hint" class="label-text">{{ hint }}</span>
    </div>
  </div>
</template>
