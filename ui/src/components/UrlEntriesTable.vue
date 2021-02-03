<template>
  <PaginatedTable class="PaginatedTable">
    <template v-slot:header>
      <th class="w-12" />
      <!-- Shortened -->
      <th class="px-4 h-12 w-48 text-left">URL</th>
      <!-- Arrow -->
      <th class="w-8" />
      <!-- Real URL -->
      <th />
      <!-- Action buttons -->
      <th class="w-48 px-4 h-12">Actions</th>
    </template>
    <template v-slot:items>
      <EntryRow
        v-for="(item, index) of items"
        :key="item.shortened"
        :entry="item"
        :class="index % 2 === 0 ? 'bg-white' : 'bg-gray-100'"
        @reload="getUrlEntries"
      />
    </template>
    <template v-slot:footer v-if="showingFooter">
      <td class="bg-white p-4" colspan="5">
        <div class="w-full h-full text-center font-bold">
          <span v-if="isLoading">Loading...</span>
          <span v-else-if="isError">{{ errorMessage || "No error message" }}</span>
          <span v-else-if="isEmpty">No URLs</span>
          <span v-else>Unknown state</span>
        </div>
      </td>
    </template>
  </PaginatedTable>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from "vue";
import api, { UrlEntry } from "../api";
import PaginatedTable from "./PaginatedTable.vue";
import EntryRow from "./EntryRow.vue";

export default defineComponent({
  components: { PaginatedTable, EntryRow },
  setup() {
    const items = ref<UrlEntry[]>([]);
    const isEmpty = computed(() => items.value.length === 0);

    const isLoading = ref(false);
    const errorMessage = ref<undefined | string>(undefined);
    const isError = computed(() => !!errorMessage.value);

    const showingFooter = computed(() => isEmpty.value || isError.value);

    const page = ref(0);
    const size = ref(20);
    const getUrlEntries = async (): Promise<void> => {
      try {
        isLoading.value = true;
        errorMessage.value = undefined;
        items.value = await api.listUrlEntries(page.value, size.value);
      } catch (err) {
        errorMessage.value = err.message;
      } finally {
        isLoading.value = false;
      }
    };
    onMounted(getUrlEntries);

    return {
      items,
      isLoading,
      isEmpty,
      isError,
      showingFooter,
      errorMessage,
      getUrlEntries,
    };
  },
});
</script>
