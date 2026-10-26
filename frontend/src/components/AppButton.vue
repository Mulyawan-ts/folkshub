<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: (val) => ['primary', 'secondary', 'accent', 'ghost', 'error', 'success', 'warning'].includes(val)
  },
  size: {
    type: String,
    default: 'md',
    validator: (val) => ['xs', 'sm', 'md', 'lg'].includes(val)
  },
  type: {
    type: String,
    default: 'button'
  },
  outline: {
    type: Boolean,
    default: false
  },
  loading: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

defineEmits(['click'])

const buttonClasses = computed(() => {
  const classes = ['btn']

  // Variant & Outline
  if (props.outline) classes.push('btn-outline')
  if (props.variant) classes.push(`btn-${props.variant}`)

  // Ukuran
  if (props.size !== 'md') classes.push(`btn-${props.size}`)

  return classes.join(' ')
})
</script>

<template>
  <button
    :type="type"
    :class="buttonClasses"
    :disabled="disabled || loading"
    @click="$emit('click', $event)"
  >
    <span v-if="loading" class="loading loading-spinner loading-xs"></span>
    <slot />
  </button>
</template>
