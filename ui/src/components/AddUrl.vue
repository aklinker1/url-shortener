<template>
  <form @submit.prevent="createNewUrlEntry">
    <div class="flex flex-row-reverse h-12">
      <div
        class="overflow-hidden transition-all"
        :class="{ 'max-w-0': !isShowingNew, 'max-w-xl': !!isShowingNew }"
      >
        <button
          type="submit"
          class="h-full w-full bg-red-800 text-white box-border px-4 uppercase font-medium rounded-r-md hover:bg-red-700 focus:bg-red-700"
          :disabled="isLoading"
        >
          Shorten
        </button>
      </div>
      <input
        class="px-4 flex-1 bg-white rounded-l-md outline-none ring-0 ring-red-100 ring-opacity-50 focus:ring"
        :class="{
          'rounded-r-md': !isShowingNew,
          'text-red-500': isLoggedIn && !isValidUrl,
        }"
        placeholder="Enter a url..."
        v-model="url"
        :disabled="isLoading"
      />
    </div>
    <div class="flex flex-row mt-3 text-white">
      <template v-if="isError">
        <ErrorIcon />
        <p class="ml-4">{{ errorMessage }}</p>
      </template>
    </div>
  </form>
</template>

<script lang="ts" setup>
import { computed, defineEmit, ref } from "vue";
import ErrorIcon from "../assets/ErrorIcon.vue";
import api from "../api";
import { useAuth } from '../composition/auth';

const emit = defineEmit(["created"]);

const url = ref("");
const isValidUrl = computed(
  () =>
    !!url.value.match(
      /^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$/
    )
);

const errorMessage = ref<undefined | string>(undefined);
const isError = computed(() => !!errorMessage.value);

const isLoading = ref(false);
async function createNewUrlEntry() {
  const value = url.value.trim();
  if (value == "") return;

  try {
    isLoading.value = true;
    errorMessage.value = undefined;
    const newEntry = await api.createUrlEntry(value);
    emit("created", newEntry);
    url.value = "";
  } catch (err) {
    errorMessage.value = err.message;
  } finally {
    isLoading.value = false;
  }
}

const { isLoggedIn } = useAuth();
const isShowingNew = computed(() => isLoggedIn.value && isValidUrl.value)
</script>
