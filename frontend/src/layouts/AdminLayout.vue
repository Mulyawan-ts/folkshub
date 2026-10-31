<script setup lang="ts">
import { computed, ref, watchEffect } from "vue";
import { storeToRefs } from "pinia";
import Layout from "./Layout.vue";
import AppSwitch from "@/components/AppSwitch.vue";
import AppBreadcrumb from "@/components/AppBreadcrumb.vue";
import AppToast from "@/components/AppToast.vue";
import { useToastStore } from "@/stores/toast";

export interface Halaman {
  id: string;
  label: string;
}

const props = defineProps<{
  halaman: Halaman[];
  aktif: string;
}>();

const emit = defineEmits<{
  "update:aktif": [id: string];
}>();

const toastStore = useToastStore();
const { items } = storeToRefs(toastStore);

const labelAktif = computed(
  () =>
    props.halaman.find((h) => h.id === props.aktif)?.label ||
    props.halaman[0]?.label ||
    "",
);

const THEME_KEY = "folkshub-theme";
const gelap = ref(localStorage.getItem(THEME_KEY) === "business");

watchEffect(() => {
  const tema = gelap.value ? "business" : "corporate";
  document.documentElement.dataset.theme = tema;
  localStorage.setItem(THEME_KEY, tema);
});

function pilihHalaman(id: string) {
  emit("update:aktif", id);
}
</script>

<template>
  <Layout>
    <template #header>
      <h1 class="text-2xl font-semibold">Folkshub</h1>
      <div class="ml-auto">
        <AppSwitch v-model="gelap" activeText="Dark" inactiveText="Light" />
      </div>
    </template>

    <template #navigasi>
      <div class="tabs tabs-bordered">
        <div class="tab" v-for="h in halaman" :key="h.id">
          <button
            :class="{ 'tab-active': aktif === h.id }"
            @click="pilihHalaman(h.id)"
          >
            {{ h.label }}
          </button>
        </div>
      </div>
    </template>

    <AppBreadcrumb
      :items="[{ label: 'Folkshub' }, { label: labelAktif }]"
      class="mb-4"
    />

    <slot />

    <template #footer>
      <AppToast :items="items" @dismiss="toastStore.dismiss" />
    </template>
  </Layout>
</template>
