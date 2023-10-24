import type TodoService from "@/features/todo/services/TodoService";
import type Todo from "@/features/todo/models/Todo";

import type {Axios} from "axios";

export default class ConcreteTodoService implements TodoService {

    protected client: Axios;

    constructor(client: Axios) {
        this.client = client;
    }

    add = async (todo: Todo): Promise<Todo[]> => {
        const formData = new URLSearchParams();
        formData.append('value', todo.value || "");
        const response = await this.client?.post(
            '/api/todos',
            formData,
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
            },
        );

        return response.data;
    };

    findAll = async (): Promise<Todo[]> => {
        const response = await this.client.get('/api/todos');

        if (response.status === 204) {
            return [];
        }

        return response.data;
    };

    modify = async (todo: Todo): Promise<Todo[]> => {
        const formData = new URLSearchParams();
        if (todo.value !== null) {
            formData.append('value', todo.value);
        }
        if (todo.completed !== null) {
            formData.append('completed', todo.completed ? "true" : "false");
        }
        const response = await this.client.patch(
            '/api/todos/' + todo.id,
            formData,
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
            },
        );

        return response.data;
    };

    remove = async (id: string): Promise<Todo[]> => {
        const response = await this.client.delete(
            '/api/todos/' + id,
        );

        if (response.status === 204) {
            return [];
        }

        return response.data;
    };

}
