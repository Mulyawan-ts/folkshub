<script setup lang="ts">
// Tanpa prop `untuk`, label dirender sebagai <span> karena id elemen kontrol
// berada di luar komponen dan tidak bisa dihubungkan secara otomatis.
withDefaults(
  defineProps<{
    text: string;
    required?: boolean;
    untuk?: string;
    hint?: string;
  }>(),
  {
    required: false,
    untuk: "",
    hint: "",
  },
);
</script>

<template>
  <component
    :is="untuk ? 'label' : 'span'"
    :for="untuk || undefined"
    class="label"
  >
    <span class="label-text">
      {{ text }}<span v-if="required" class="text-error">*</span>
    </span>
    <span v-if="hint" class="label-text-alt">{{ hint }}</span>
  </component>
</template>
