import {createRouter, createWebHistory} from "vue-router";
import {decode} from "jsonwebtoken-esm";

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

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem("token");
    if (to.matched.some(record => record.meta.requiresAuth) && isTokenNotExistOrExpired(token)) {
        next({name: "Login"});
    } else {
        next();
    }
});

function isTokenNotExistOrExpired(token: string | null) {
    if (token == null || token.length == 0) {
        return true;
    }
    try {
        const decodedToken = decode(token, {json: true});
        if (decodedToken == null) {
            return true;
        }
        if (typeof decodedToken.exp != "number") {
            return true;
        }

        const expirationDate = new Date(decodedToken.exp * 1000);

        return expirationDate < new Date();
    } catch (error) {
        console.error("An error occurred while trying to check token:", error);

        return true;
    }
}

export default router
