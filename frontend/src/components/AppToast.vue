<script setup lang="ts">
import type { ToastItem, ToastType } from "../stores/toast";

withDefaults(
  defineProps<{
    items?: ToastItem[];
    position?: "top-right" | "top-left" | "bottom-right" | "bottom-left";
  }>(),
  {
    items: () => [],
    position: "bottom-right",
  },
);

const emit = defineEmits<{
  dismiss: [id: number];
}>();

const variantClasses: Record<ToastType, string> = {
  success: "alert-success",
  error: "alert-error",
  warning: "alert-warning",
  info: "alert-info",
};

const iconPaths: Record<ToastType, string> = {
  success: "M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z",
  error:
    "M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z",
  warning:
    "M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z",
  info: "M11.25 11.25l.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z",
};

const positionClasses: Record<string, string> = {
  "top-right": "top-4 right-4",
  "top-left": "top-4 left-4",
  "bottom-right": "bottom-4 right-4",
  "bottom-left": "bottom-4 left-4",
};
</script>

<template>
  <div
    class="pointer-events-none fixed z-50 flex flex-col gap-2"
    :class="positionClasses[position]"
    role="region"
    aria-label="Notifications"
  >
    <TransitionGroup name="toast">
      <div
        v-for="item in items"
        :key="item.id"
        class="pointer-events-auto flex w-80 max-w-[calc(100vw-2rem)] items-start gap-3 rounded-lg border border-base-300 bg-base-100 p-4 shadow-lg"
        role="alert"
      >
        <svg
          class="mt-0.5 size-5 shrink-0"
          :class="variantClasses[item.type]"
          fill="none"
          stroke-width="1.5"
          stroke="currentColor"
          viewBox="0 0 24 24"
          aria-hidden="true"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            :d="iconPaths[item.type]"
          />
        </svg>

        <div class="min-w-0 flex-1">
          <p v-if="item.title" class="font-semibold">{{ item.title }}</p>
          <p
            class="text-sm wrap-break-word"
            :class="item.title ? 'opacity-80' : ''"
          >
            {{ item.message }}
          </p>

          <button
            v-if="item.action"
            type="button"
            class="btn btn-ghost btn-xs mt-2 -ml-2"
            @click="emit('dismiss', item.id)"
          >
            {{ item.action.label }}
          </button>
        </div>

        <button
          type="button"
          class="btn btn-ghost btn-xs btn-circle -mt-1 -mr-1 shrink-0"
          aria-label="Dismiss notification"
          @click="emit('dismiss', item.id)"
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
    </TransitionGroup>
  </div>
</template>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition:
    opacity 0.2s ease,
    transform 0.2s ease;
}

.toast-enter-from,
.toast-leave-to {
  opacity: 0;
  transform: translateY(0.5rem);
}
</style>
