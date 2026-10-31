<script setup lang="ts">
import { computed } from "vue";
import type { Variant } from "@/types/ui";

const props = withDefaults(
  defineProps<{
    variant?: Variant;
    judul?: string;
    pesan?: string;
    soft?: boolean;
    dismissible?: boolean;
  }>(),
  {
    variant: "info",
    judul: "",
    pesan: "",
    soft: false,
    dismissible: false,
  },
);

const emit = defineEmits<{
  dismiss: [];
}>();

const alertClasses = computed(() => {
  const classes = ["alert"];
  classes.push(
    props.soft ? `alert-soft alert-${props.variant}` : `alert-${props.variant}`,
  );
  return classes.join(" ");
});
</script>

<template>
  <div :class="alertClasses" role="alert">
    <div class="min-w-0 flex-1">
      <p v-if="judul" class="font-semibold">{{ judul }}</p>
      <span v-if="pesan" :class="judul ? 'text-sm opacity-80' : ''">
        {{ pesan }}
      </span>
      <slot v-else />
    </div>

    <button
      v-if="dismissible"
      type="button"
      class="btn btn-ghost btn-xs btn-circle shrink-0"
      aria-label="Tutup"
      @click="emit('dismiss')"
    >
      <svg
        class="size-3.5"
        fill="none"
        stroke-width="2"
        stroke="currentColor"
        viewBox="0 0 24 24"
        aria-hidden="true"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M6 18 18 6M6 6l12 12"
        />
      </svg>
    </button>
  </div>
</template>
