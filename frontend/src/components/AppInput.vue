<script setup lang="ts">
import { computed, useId } from 'vue'
import type { Size } from '../types/ui'

const props = withDefaults(
  defineProps<{
    modelValue?: string | number
    type?: string
    variant?: 'default' | 'filled' | 'ghost'
    size?: Size
    placeholder?: string
    label?: string
    hint?: string
    error?: string
    disabled?: boolean
    readonly?: boolean
    required?: boolean
    maxlength?: number
  }>(),
  {
    modelValue: '',
    type: 'text',
    variant: 'default',
    size: 'md',
    placeholder: '',
    label: '',
    hint: '',
    error: '',
    disabled: false,
    readonly: false,
    required: false,
    maxlength: undefined
  }
)

const emit = defineEmits<{
  'update:modelValue': [value: string | number]
  enter: [value: string | number]
}>()

const inputId = useId()

const hasError = computed(() => props.error !== '')

const inputClasses = computed(() => {
  const classes = ['input']

  if (props.variant === 'filled') classes.push('input-filled')
  if (props.variant === 'ghost') classes.push('input-ghost')
  if (props.size !== 'md') classes.push(`input-${props.size}`)
  if (hasError.value) classes.push('input-error')

  return classes.join(' ')
})

function onInput(event: Event) {
  emit('update:modelValue', (event.target as HTMLInputElement).value)
}

function onKeydown(event: KeyboardEvent) {
  if (event.key === 'Enter') emit('enter', (event.target as HTMLInputElement).value)
}
</script>

<template>
  <div class="form-control w-full">
    <label v-if="label" :for="inputId" class="label">
      <span class="label-text">{{ label }}</span>
      <span v-if="required" class="text-error">*</span>
    </label>

    <input
      :id="inputId"
      :type="type"
      :class="inputClasses"
      :value="modelValue"
      :placeholder="placeholder"
      :disabled="disabled"
      :readonly="readonly"
      :required="required"
      :maxlength="maxlength"
      @input="onInput"
      @keydown="onKeydown"
    />

    <div v-if="error || hint" class="label">
      <span v-if="error" class="label-text text-error">{{ error }}</span>
      <span v-else-if="hint" class="label-text">{{ hint }}</span>
    </div>
  </div>
</template>
