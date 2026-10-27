<script setup>
import { computed, useId } from 'vue'

const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: ''
  },
  type: {
    type: String,
    default: 'text'
  },
  variant: {
    type: String,
    default: 'default',
    validator: (val) => ['default', 'filled', 'ghost'].includes(val)
  },
  size: {
    type: String,
    default: 'md',
    validator: (val) => ['xs', 'sm', 'md', 'lg'].includes(val)
  },
  placeholder: {
    type: String,
    default: ''
  },
  label: {
    type: String,
    default: ''
  },
  hint: {
    type: String,
    default: ''
  },
  error: {
    type: String,
    default: ''
  },
  disabled: {
    type: Boolean,
    default: false
  },
  readonly: {
    type: Boolean,
    default: false
  },
  required: {
    type: Boolean,
    default: false
  },
  maxlength: {
    type: Number,
    default: undefined
  }
})

const emit = defineEmits(['update:modelValue', 'enter'])

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

const wrapperClasses = computed(() => {
  const classes = ['form-control', 'w-full']
  if (props.label) classes.push('w-full')
  return classes.join(' ')
})

function onInput(event) {
  emit('update:modelValue', event.target.value)
}

function onKeydown(event) {
  if (event.key === 'Enter') emit('enter', event.target.value)
}
</script>

<template>
  <div :class="wrapperClasses">
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
