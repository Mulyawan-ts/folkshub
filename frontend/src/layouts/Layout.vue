<script setup lang="ts">
withDefaults(
  defineProps<{
    lebarMaksimal?: string;
  }>(),
  {
    lebarMaksimal: "max-w-6xl",
  },
);
</script>

<template>
  <div class="min-h-screen bg-base-100 text-base-content">
    <header class="border-b border-base-300">
      <div
        :class="['mx-auto flex items-center gap-4 px-6 py-4', lebarMaksimal]"
      >
        <slot name="header" />
      </div>
    </header>

    <nav
      v-if="$slots.navigasi"
      class="border-b border-base-300 bg-base-200"
      aria-label="Navigasi utama"
    >
      <div :class="['mx-auto px-6', lebarMaksimal]">
        <slot name="navigasi" />
      </div>
    </nav>

    <main :class="['mx-auto px-6 py-6', lebarMaksimal]">
      <slot />
    </main>

    <slot name="footer" />
  </div>
</template>
