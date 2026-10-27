<script setup lang="ts">
import { computed } from 'vue'
import type { Size, Variant } from '../types/ui'

const props = withDefaults(
  defineProps<{
    variant?: Variant
    size?: Size
    outline?: boolean
    style?: 'filled' | 'soft'
  }>(),
  {
    variant: 'primary',
    size: 'md',
    outline: false,
    style: 'filled'
  }
)

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
