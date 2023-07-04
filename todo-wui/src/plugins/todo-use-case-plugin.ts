import {useTodoFeature} from "@/features/todo/useTodoFeature";
import type TodoUseCase from "@/features/todo/usecases/TodoUseCase";

import type {App} from "vue";
import type {Router} from "vue-router";

export default function todoUseCasePlugin(app: App, router: Router) {
    const todoUseCase: TodoUseCase = useTodoFeature();

    app.provide("todoUseCase", todoUseCase);

    router.addRoute({
        path: "/todo",
        name: "Todo",
        component: () => import("@/features/todo/views/TodoManagementView.vue"),
        meta: {
            requiresAuth: true,
        }
    });
}
