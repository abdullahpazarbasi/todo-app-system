import type TodoUseCase from "@/features/todo/usecases/TodoUseCase";
import type TodoItem from "@/features/todo/entities/TodoItem";
import type TodoService from "@/features/todo/services/TodoService";
import ConcreteTodoService from "@/features/todo/services/ConcreteTodoService";
import ConcreteTodoUseCase from "@/features/todo/usecases/ConcreteTodoUseCase";

import {ref} from "vue";
import axios from "axios";

export function useTodoFeature(): TodoUseCase {
    const httpClient = axios.create({
        baseURL: import.meta.env.VITE_TODO_WBFF_BASE_URL + "/api/todos",
    });
    httpClient.interceptors.request.use(
        (config) => {
            const token = localStorage.getItem('token');
            if (token != null && token.length > 0) {
                config.headers.Authorization = `Bearer ${token}`;
            }

            return config;
        },
    );
    const todoService: TodoService = new ConcreteTodoService(
        httpClient,
    );

    return new ConcreteTodoUseCase(
        ref<TodoItem[]>([]),
        todoService,
    );
}