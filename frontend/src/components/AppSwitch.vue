<script setup>
import { computed, useId } from 'vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
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
  activeText: {
    type: String,
    default: ''
  },
  inactiveText: {
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
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

const switchId = useId()

const switchClasses = computed(() => {
  const classes = ['toggle']
  classes.push(`toggle-${props.variant}`)
  if (props.size !== 'md') classes.push(`toggle-${props.size}`)
  return classes.join(' ')
})

function onChange(event) {
  emit('update:modelValue', event.target.checked)
  emit('change', event.target.checked)
}
</script>

<template>
  <div class="flex items-center gap-2">
    <input
      :id="switchId"
      type="checkbox"
      role="switch"
      :class="switchClasses"
      :checked="modelValue"
      :disabled="disabled"
      :required="required"
      @change="onChange"
    />

    <label v-if="label" :for="switchId" class="label cursor-pointer">
      <span class="label-text">{{ label }}</span>
    </label>

    <span v-if="activeText || inactiveText" class="label-text">
      {{ modelValue ? activeText : inactiveText }}
    </span>
  </div>
</template>
