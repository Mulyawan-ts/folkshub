<script setup lang="ts">
import CariInput from "./CariInput.vue";
import AppCard from "./AppCard.vue";

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
  <AppCard scroll>
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
  </AppCard>
</template>
