import type TodoItem from "@/features/todo/entities/TodoItem";
import type {Ref} from "vue";

export default interface TodoUseCase {
    getReactiveCollection(): Ref<TodoItem[]>;
    add(todoText: string): void;
    retrieveAll(): void;
    modify(todoItem: TodoItem): void;
    remove(id: string): void;
}