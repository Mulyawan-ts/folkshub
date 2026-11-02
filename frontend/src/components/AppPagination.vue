<script setup lang="ts">
const props = defineProps<{
  halaman: number
  totalHalaman: number
}>()

const emit = defineEmits<{ "update:halaman": [value: number] }>()

function ke(h: number) {
  if (h < 1 || h > props.totalHalaman) return
  emit("update:halaman", h)
}
</script>

<template>
  <div class="join">
    <button class="btn btn-sm join-item" :disabled="halaman <= 1" @click="ke(halaman - 1)">«</button>
    <button class="btn btn-sm join-item" disabled>{{ halaman }} / {{ totalHalaman }}</button>
    <button class="btn btn-sm join-item" :disabled="halaman >= totalHalaman" @click="ke(halaman + 1)">»</button>
  </div>
</template>
