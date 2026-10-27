<script setup>
import { computed, useId } from 'vue'

const props = defineProps({
  modelValue: {
    type: [Boolean, Array],
    default: false
  },
  value: {
    type: [String, Number, Boolean, Object],
    default: undefined
  },
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
  label: {
    type: String,
    default: ''
  },
  disabled: {
    type: Boolean,
    default: false
  },
  required: {
    type: Boolean,
    default: false
  },
  indeterminate: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

const checkboxId = useId()

const isArrayModel = computed(() => Array.isArray(props.modelValue))

const checked = computed(() => {
  if (isArrayModel.value) {
    if (props.value === undefined) return false
    return props.modelValue.some((item) => item === props.value)
  }
  return props.modelValue === true
})

const checkboxClasses = computed(() => {
  const classes = ['checkbox']
  classes.push(`checkbox-${props.variant}`)
  if (props.size !== 'md') classes.push(`checkbox-${props.size}`)
  return classes.join(' ')
})

function onChange(event) {
  const target = event.target

  if (isArrayModel.value) {
    let next
    if (target.checked) {
      next = props.value === undefined ? [...props.modelValue, target.value] : [...props.modelValue, props.value]
    } else {
      next = props.value === undefined
        ? props.modelValue.filter((item) => item !== target.value)
        : props.modelValue.filter((item) => item !== props.value)
    }
    emit('update:modelValue', next)
  } else {
    emit('update:modelValue', target.checked)
  }

  emit('change', target.checked)
}
</script>

<template>
  <div class="flex items-start gap-2">
    <input
      :id="checkboxId"
      type="checkbox"
      :class="checkboxClasses"
      :checked="checked"
      :disabled="disabled"
      :required="required"
      :indeterminate="indeterminate"
      @change="onChange"
    />

    <label v-if="label" :for="checkboxId" class="label cursor-pointer">
      <span class="label-text">{{ label }}</span>
    </label>
  </div>
</template>
