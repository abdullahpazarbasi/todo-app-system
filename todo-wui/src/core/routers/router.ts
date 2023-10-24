import {createRouter, createWebHistory} from "vue-router";
import {loading} from "@/core/composables/loading";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "Home",
            component: () => import("@/core/views/HomeView.vue"),
        },
    ],
});

router.beforeEach(() => {
    loading.value = true;
    console.log("loading");
});

router.afterEach(() => {
    loading.value = false;
    console.log("loaded");
});

export default router;
