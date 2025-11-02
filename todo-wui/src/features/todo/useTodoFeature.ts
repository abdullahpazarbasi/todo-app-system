import type TodoUseCase from "@/features/todo/usecases/TodoUseCase";
import type TodoItem from "@/features/todo/entities/TodoItem";
import type TodoService from "@/features/todo/services/TodoService";
import ConcreteTodoService from "@/features/todo/services/ConcreteTodoService";
import ConcreteTodoUseCase from "@/features/todo/usecases/ConcreteTodoUseCase";

import {ref} from "vue";
import type HttpClient from "@/core/http/HttpClient";

export function useTodoFeature(httpClient: HttpClient): TodoUseCase {
    const todoService: TodoService = new ConcreteTodoService(
        httpClient,
    );

    return new ConcreteTodoUseCase(
        ref<TodoItem[]>([]),
        todoService,
    );
}