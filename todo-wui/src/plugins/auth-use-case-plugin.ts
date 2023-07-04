import {useAuthFeature} from "@/features/auth/useAuthFeature";
import type AuthUseCase from "@/features/auth/usecases/AuthUseCase";

import type {App} from "vue";
import type {Router} from "vue-router";

export default function authUseCasePlugin(app: App, router: Router) {
    const authUseCase: AuthUseCase = useAuthFeature(router);

    app.provide("authUseCase", authUseCase);

    router.addRoute({
        path: "/login",
        name: "Login",
        component: () => import("@/features/auth/views/LoginView.vue"),
    });
}
