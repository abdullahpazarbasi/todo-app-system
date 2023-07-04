<script setup lang="ts">
import TodoItemComponent from "@/features/todo/components/TodoItemComponent.vue";
import type TodoItem from "@/features/todo/entities/TodoItem";

const props = defineProps<{
  collection: TodoItem[],
}>();

const emit = defineEmits<{
  (e: 'item-set-as-completed', id: string): void
  (e: 'item-set-as-not-completed', id: string): void
  (e: 'item-removed', id: string): void
}>();

function setItemAsCompleted(id: string) {
  if (id.length > 0) {
    emit('item-set-as-completed', id);
  }
}

function setItemAsNotCompleted(id: string) {
  if (id.length > 0) {
    emit('item-set-as-not-completed', id);
  }
}

function removeItem(id: string) {
  if (id.length > 0) {
    emit('item-removed', id);
  }
}
</script>

<template>
  <ul>
    <li v-for="item in props.collection" v-bind:key="item.id!">
      <TodoItemComponent v-bind:item="item"
                         v-on:item-set-as-completed="setItemAsCompleted"
                         v-on:item-set-as-not-completed="setItemAsNotCompleted"
                         v-on:item-removed="removeItem"/>
    </li>
  </ul>
</template>
