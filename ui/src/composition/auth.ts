import { computed, inject, provide, Ref, ref } from 'vue';
import firebase from 'firebase/app';
import { useRouter } from 'vue-router';

export function useAuthSetup() {
    const user = ref<any | undefined>();
    provide("user", user);
    firebase.auth().onAuthStateChanged((newUser) => {
        console.log("Auth state changed:", {newUser})
        user.value = newUser
    })
}

export function useAuth() {
    const user = inject<any | undefined>("user", undefined);

    const isLoggedIn = computed(() => user.value != null);
    const logout = () => {
        firebase.auth().signOut();
    }
    const router = useRouter();
    const signInSuccess = (): false => {
        router.push("/@/ui");
        return false;
    }

    return {
        logout,
        isLoggedIn,
        signInSuccess,
    }
}
