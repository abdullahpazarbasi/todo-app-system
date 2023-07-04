import type TodoService from "@/features/todo/services/TodoService";
import type TodoItem from "@/features/todo/entities/TodoItem";
import type TodoUseCase from "@/features/todo/usecases/TodoUseCase";
import type Todo from "@/features/todo/models/Todo";

import type {Ref} from "vue";

export default class ConcreteTodoUseCase implements TodoUseCase {

    protected reactiveCollection: Ref<TodoItem[]>;
    protected todoService: TodoService;

    constructor(collection: Ref<TodoItem[]>, todoService: TodoService) {
        this.reactiveCollection = collection;
        this.todoService = todoService;
    }

    getReactiveCollection = (): Ref<TodoItem[]> => this.reactiveCollection;

    add = async (todoText: string): Promise<void> => {
        const itemList = await this.todoService.add({
            id: null,
            value: todoText,
            completed: false,
        });
        this.reactiveCollection.value = [];
        itemList.forEach((item) => {
            this.reactiveCollection.value.push(<TodoItem>{
                id: item.id,
                value: item.value,
                completed: item.completed,
            });
        });
    };

    retrieveAll = async (): Promise<void> => {
        const itemList = await this.todoService.findAll();
        this.reactiveCollection.value = [];
        itemList.forEach((item) => {
            this.reactiveCollection.value.push(<TodoItem>{
                id: item.id,
                value: item.value,
                completed: item.completed,
            });
        });
    };

    modify = async (todoItem: TodoItem): Promise<void> => {
        const itemList = await this.todoService.modify(<Todo>{
            id: todoItem.id,
            value: todoItem.value,
            completed: todoItem.completed,
        });
        this.reactiveCollection.value = [];
        itemList.forEach((item) => {
            this.reactiveCollection.value.push(<TodoItem>{
                id: item.id,
                value: item.value,
                completed: item.completed,
            });
        });
    };

    remove = async (id: string): Promise<void> => {
        const itemList = await this.todoService.remove(id);
        this.reactiveCollection.value = [];
        itemList.forEach((item) => {
            this.reactiveCollection.value.push(<TodoItem>{
                id: item.id,
                value: item.value,
                completed: item.completed,
            });
        });
    };

}