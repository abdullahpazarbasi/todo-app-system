import {useTodoFeature} from "@/features/todo/useTodoFeature";
import type TodoUseCase from "@/features/todo/usecases/TodoUseCase";

import type {App} from "vue";
import type {Router} from "vue-router";
import type HttpClient from "@/core/http/HttpClient";

export default function todoUseCasePlugin(app: App, router: Router, httpClient: HttpClient) {
    const todoUseCase: TodoUseCase = useTodoFeature(httpClient);

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
