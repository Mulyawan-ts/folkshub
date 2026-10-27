<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: (val) => ['primary', 'secondary', 'accent', 'neutral', 'info', 'success', 'warning', 'error'].includes(val)
  },
  size: {
    type: String,
    default: 'md',
    validator: (val) => ['xs', 'sm', 'md', 'lg'].includes(val)
  },
  outline: {
    type: Boolean,
    default: false
  },
  style: {
    type: String,
    default: 'filled',
    validator: (val) => ['filled', 'soft'].includes(val)
  }
})

const badgeClasses = computed(() => {
  const classes = ['badge']

  if (props.style === 'soft') classes.push('badge-soft')
  if (props.outline) classes.push('badge-outline')
  classes.push(`badge-${props.variant}`)

  if (props.size !== 'md') classes.push(`badge-${props.size}`)

  return classes.join(' ')
})
</script>

<template>
  <span :class="badgeClasses">
    <slot />
  </span>
</template>
