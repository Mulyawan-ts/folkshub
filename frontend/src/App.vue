<script setup lang="ts">
import { storeToRefs } from 'pinia'
import AppToast from './components/AppToast.vue'
import { useToastStore } from './stores/toast'

const toastStore = useToastStore()
const { items } = storeToRefs(toastStore)
</script>

<template>
  <main class="p-8">
    <h1 class="text-2xl font-semibold">Folkshub</h1>
  </main>

  <AppToast :items="items" @dismiss="toastStore.dismiss" />
</template>
