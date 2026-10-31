<script setup lang="ts">
defineProps<{
  colspan: number;
  kosong?: string;
  kosongCari?: string;
  mencari?: boolean;
}>();
</script>

<template>
  <tr>
    <td :colspan="colspan" class="text-center text-base-content/60 py-4">
      <slot>
        {{ mencari ? (kosongCari ?? "Tidak ada hasil yang cocok") : kosong }}
      </slot>
    </td>
  </tr>
</template>
