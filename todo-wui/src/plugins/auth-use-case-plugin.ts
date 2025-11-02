import {useAuthFeature} from "@/features/auth/useAuthFeature";
import type AuthUseCase from "@/features/auth/usecases/AuthUseCase";

import type {App} from "vue";
import type {Router} from "vue-router";
import type HttpClient from "@/core/http/HttpClient";
import {decode} from "jsonwebtoken-esm";

export default function authUseCasePlugin(app: App, router: Router, httpClient: HttpClient) {
    const authUseCase: AuthUseCase = useAuthFeature(router, httpClient);

    app.provide("authUseCase", authUseCase);

    router.addRoute({
        path: "/login",
        name: "Login",
        component: () => import("@/features/auth/views/LoginView.vue"),
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
}
