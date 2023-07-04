<script setup lang="ts">
import type TodoUseCase from "@/features/todo/usecases/TodoUseCase";
import TodoCandidateComponent from "@/features/todo/components/TodoCandidateComponent.vue";
import TodoListComponent from "@/features/todo/components/TodoListComponent.vue";
import TopNavBar from "@/core/components/TopNavBar.vue";

import {inject, onMounted, reactive} from "vue";

const todoUseCase: TodoUseCase | undefined = inject("todoUseCase");
const todoCollection = todoUseCase!.getReactiveCollection();
const todoCandidate = reactive({value: ""});

function setTodoAsCompleted(id: string) {
  todoUseCase!.modify({
    id: id,
    value: null,
    completed: true,
  });
}

function setTodoAsNotCompleted(id: string) {
  todoUseCase!.modify({
    id: id,
    value: null,
    completed: false,
  });
}

function removeTodo(id: string) {
  todoUseCase!.remove(id);
}

onMounted(todoUseCase!.retrieveAll);
</script>

<template>
  <TopNavBar>
    <router-link to="/">Home</router-link>
  </TopNavBar>
  <main>
    <TodoCandidateComponent v-bind:candidate="todoCandidate"
                            v-on:candidate-value-set="todoUseCase!.add"/>
    <TodoListComponent v-bind:collection="todoCollection"
                       v-on:item-set-as-completed="setTodoAsCompleted"
                       v-on:item-set-as-not-completed="setTodoAsNotCompleted"
                       v-on:item-removed="removeTodo"/>
  </main>
</template>