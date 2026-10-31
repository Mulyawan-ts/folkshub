<script setup lang="ts">
import CariInput from "./CariInput.vue";

withDefaults(
  defineProps<{
    cari?: string;
    placeholder?: string;
  }>(),
  {
    cari: "",
    placeholder: "Cari...",
  },
);

defineEmits<{
  "update:cari": [value: string];
}>();
</script>

<template>
  <div class="card p-4 bg-base-100 shadow-sm overflow-x-auto">
    <CariInput
      :model-value="cari"
      :placeholder="placeholder"
      @update:model-value="$emit('update:cari', $event)"
    />
    <table class="table w-full text-left">
      <thead>
        <tr>
          <slot name="head" />
        </tr>
      </thead>
      <tbody>
        <slot />
      </tbody>
    </table>
  </div>
</template>
