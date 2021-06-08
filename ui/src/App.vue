<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useAuthSetup } from './composition/auth';

export default defineComponent({
  setup() {
    useAuthSetup();
  },
})
</script>
