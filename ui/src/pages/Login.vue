<template>
  <div>
    <div id="firebase-login-container" :class="{ hidden: isLoggedIn }" />
    <div v-if="isLoggedIn">
      <button @click="logout">Logout</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, watch } from "vue";
import * as firebaseui from "firebaseui";
import firebase from "firebase/app";
import { useAuth } from "../composition/auth";

export default defineComponent({
  setup() {
    const { signInSuccess, isLoggedIn, logout } = useAuth();

    onMounted(() => {
      const ui = new firebaseui.auth.AuthUI(firebase.auth());
      ui.start("#firebase-login-container", {
        signInOptions: [
          {
            provider: firebase.auth.EmailAuthProvider.PROVIDER_ID,
            requireDisplayName: false,
          },
        ],
        callbacks: {
          signInSuccessWithAuthResult: signInSuccess,
        },
      });
    });
    return {
      isLoggedIn,
      logout,
    };
  },
});
</script>
