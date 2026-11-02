<script setup lang="ts">
const props = defineProps<{
  tabs: { label: string; value: string }[]
  modelValue: string
}>()

const emit = defineEmits<{ "update:modelValue": [value: string] }>()
</script>

<template>
  <div role="tablist" class="tabs tabs-boxed">
    <button
      v-for="t in tabs"
      :key="t.value"
      role="tab"
      class="tab"
      :class="{ 'tab-active': t.value === modelValue }"
      @click="emit('update:modelValue', t.value)"
    >
      {{ t.label }}
    </button>
  </div>
</template>
