import type Todo from "@/features/todo/models/Todo";

export default interface TodoService {
    add(todoItem: Todo): Promise<Todo[]>;
    findAll(): Promise<Todo[]>;
    modify(todoItem: Todo): Promise<Todo[]>;
    remove(id: string): Promise<Todo[]>;
}