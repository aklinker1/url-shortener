import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/@/ui",
            component: () => import("./pages/Home.vue")
        },
        {
            path: "/@/ui/login",
            component: () => import("./pages/Login.vue")
        }
    ],
})

export default router;
