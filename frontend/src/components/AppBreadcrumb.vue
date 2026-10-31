<script setup lang="ts">
export interface ItemBreadcrumb {
  label: string;
  href?: string;
}

withDefaults(
  defineProps<{
    items: ItemBreadcrumb[];
  }>(),
  {},
);
</script>

<template>
  <nav aria-label="Breadcrumb">
    <ul class="breadcrumbs text-sm">
      <li v-for="(item, i) in items" :key="item.label">
        <a v-if="item.href && i < items.length - 1" :href="item.href">
          {{ item.label }}
        </a>
        <span
          v-else
          :aria-current="i === items.length - 1 ? 'page' : undefined"
        >
          {{ item.label }}
        </span>
      </li>
    </ul>
  </nav>
</template>
