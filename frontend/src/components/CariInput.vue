<script setup lang="ts">
import AppInput from "./AppInput.vue";

const props = withDefaults(
  defineProps<{
    modelValue: string;
    placeholder?: string;
    label?: string;
  }>(),
  {
    placeholder: "Cari...",
    label: "",
  },
);

const emit = defineEmits<{
  "update:modelValue": [value: string];
}>();

function onUpdate(value: string | number) {
  emit("update:modelValue", String(value));
}

function bersihkan() {
  emit("update:modelValue", "");
}
</script>

<template>
  <div class="flex items-end gap-2 mb-4">
    <AppInput
      :model-value="modelValue"
      :placeholder="placeholder"
      :label="label"
      @update:model-value="onUpdate"
    />
    <button
      v-if="modelValue"
      type="button"
      class="btn btn-ghost btn-sm"
      @click="bersihkan"
    >
      Bersihkan
    </button>
  </div>
</template>
