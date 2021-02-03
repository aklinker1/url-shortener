<template>
  <tr
    class="h-16"
    @mouseenter="toggleHover(true)"
    @mouseleave="toggleHover(false)"
  >
    <td>
      <p
        class="w-10 h-10 bg-yellow-500 rounded-full text-yellow-100 text-xl font-mono text-center box-border pt-1.5 ml-3"
      >
        {{ entry.visits }}
      </p>
    </td>
    <td class="px-4">
      <a class="text-yellow-500" :href="entry.shortened" target="_blank">{{
        entry.shortened
      }}</a>
    </td>
    <td class="px-4">
      <ArrowIcon />
    </td>
    <td>
      <input
        ref="input"
        class="bg-transparent h-10 px-4 border-none w-full rounded-md ring-yellow-500"
        :class="{ 'ring-2': isEditing, 'ring-0': !isEditing }"
        :readonly="!isEditing"
        @keyup.enter="updateUrl"
        v-model="url"
      />
    </td>
    <td class="h-16 px-4 flex flex-row justify-end items-center space-x-2">
      <template v-if="areActionsVisible">
        <div
          v-if="isEditing"
          class="w-10 h-10 box-border p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
          title="Edit URL"
          @click="updateUrl"
        >
          <CheckIcon color="#808080" />
        </div>
        <div
          v-else
          class="w-10 h-10 box-border p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
          title="Edit URL"
          @click="toggleEditing(true)"
        >
          <EditIcon color="#808080" />
        </div>
        <div
          class="w-10 h-10 box-border p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
          title="Share Shortened URL"
          @click="copyShortened"
        >
          <CheckIcon v-if="isCopied" color="green" />
          <ShareIcon v-else color="#808080" />
        </div>
        <div
          class="w-10 h-10 box-border p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
          title="Delete"
          @click="deleteUrl"
        >
          <DeleteIcon color="#808080" />
        </div>
      </template>
    </td>
  </tr>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref } from "vue";
import api, { UrlEntry } from "../api";
import ArrowIcon from "../assets/ArrowIcon.vue";
import CheckIcon from "../assets/CheckIcon.vue";
import DeleteIcon from "../assets/DeleteIcon.vue";
import EditIcon from "../assets/EditIcon.vue";
import ShareIcon from "../assets/ShareIcon.vue";

export default defineComponent({
  components: { ArrowIcon, DeleteIcon, EditIcon, ShareIcon, CheckIcon },
  props: {
    entry: {
      type: Object as PropType<UrlEntry>,
      required: true,
    },
  },
  emits: ["reload"],
  setup(props, { emit }) {
    const input = ref<HTMLInputElement | null>(null);

    const isHovered = ref(false);
    const toggleHover = (isVisible: boolean) => {
      isHovered.value = isVisible;
    };

    const url = ref(props.entry.url);
    const isEditing = ref(false);
    const toggleEditing = (newIsEditing: boolean) => {
      setTimeout(() => {
        input.value?.focus();
        input.value?.select();
      }, 0);
      isEditing.value = newIsEditing;
    };
    const updateUrl = async () => {
      toggleEditing(false);
      const newUrl = url.value.trim();
      if (newUrl === props.entry.url) return;

      try {
        await api.updateUrlEntry(props.entry.id, newUrl);
      } catch (err) {
        console.warn(err);
        url.value = props.entry.url;
      }
    };

    const deleteUrl = async () => {
      try {
        await api.deleteUrlEntry(props.entry.id);
        emit("reload");
      } catch (err) {
        console.warn(err);
      }
    };

    const isCopied = ref(false);
    let copyTimeout: any | undefined;
    const copyShortened = () => {
      input.value?.select();
      document.execCommand("copy");
      input.value?.blur();
      isCopied.value = true;
      if (copyTimeout != null) {
        clearTimeout(copyTimeout);
      }
      copyTimeout = setTimeout(() => {
        isCopied.value = false;
      }, 3000);
    };

    const areActionsVisible = computed(
      () => isHovered.value || isEditing.value || isCopied.value
    );

    return {
      input,
      toggleHover,

      url,
      isEditing,
      toggleEditing,
      updateUrl,

      areActionsVisible,

      isCopied,
      copyShortened,

      deleteUrl,
    };
  },
});
</script>
