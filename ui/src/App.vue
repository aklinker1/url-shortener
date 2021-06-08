<template>
  <div class="relative z-10 max-w-screen-lg mx-auto">
    <div class="h-80 mb-8">
      <h1
        class="
          pt-12
          h-32
          font-mono
          text-7xl
          align-baseline
          text-white
          m-0
          text-center
        "
      >
        apk.rip
      </h1>
      <h2
        class="
          h-8
          mb-12
          font-mono
          text-2xl text-white
          font-light
          text-opacity-70
          m-0
          text-center
          uppercase
        "
      >
        Url Shortener
      </h2>
      <AddUrl class="max-w-screen-md mx-auto" @created="reload" />
    </div>
    <UrlEntriesTable ref="table" class="mb-16" />
  </div>
  <div class="absolute top-0 h-80 w-full bg-red-500 z-0 overflow-hidden">
    <img src="/header.png" class="mx-auto custom-transform" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
// @ts-ignore: Using setup script, no default export
import AddUrl from "./components/AddUrl.vue";
import UrlEntriesTable from "./components/UrlEntriesTable.vue";

export default defineComponent({
  name: "App",
  components: { AddUrl, UrlEntriesTable },
  setup() {
    const table = ref<typeof UrlEntriesTable | null>(null);
    const reload = () => table.value?.getUrlEntries?.();

    return {
      table,
      reload,
    };
  },
});
</script>

<style scoped>
.custom-transform {
  transform: translate(-40%, -72px);
}
</style>
