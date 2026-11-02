<script setup lang="ts">
const props = withDefaults(defineProps<{
  label?: string
  modelValue: string
  options: { label: string; value: string }[]
  inline?: boolean
}>(), {
  label: "",
  inline: false,
})

const emit = defineEmits<{ "update:modelValue": [value: string] }>()
</script>

<template>
  <div :class="inline ? 'flex items-center gap-3' : 'form-control w-full'">
    <label v-if="label" class="label" :class="inline ? 'w-32 shrink-0' : ''">
      <span class="label-text font-medium">{{ label }}</span>
    </label>
    <select class="select select-bordered w-full" :value="modelValue" @change="emit('update:modelValue', ($event.target as HTMLSelectElement).value)">
      <option v-for="o in options" :key="o.value" :value="o.value">{{ o.label }}</option>
    </select>
  </div>
</template>
