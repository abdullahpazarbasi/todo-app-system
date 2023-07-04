<script setup lang="ts">
import type TodoItem from "@/features/todo/entities/TodoItem";

const props = defineProps<{
  item: TodoItem,
}>();

const emit = defineEmits<{
  (e: 'item-set-as-completed', id: string): void
  (e: 'item-set-as-not-completed', id: string): void
  (e: 'item-removed', id: string): void
}>();

function toggle() {
  if (props.item.id !== null && props.item.id.length > 0) {
    if (props.item.completed) {
      emit('item-set-as-not-completed', props.item.id);
    } else {
      emit('item-set-as-completed', props.item.id);
    }
  }
}

function remove() {
  if (props.item.id !== null && props.item.id.length > 0) {
    emit('item-removed', props.item.id);
  }
}
</script>

<template>
  <div>
    <span v-bind:class="{ completed: props.item.completed! }">{{ props.item.value! }}</span>
    <button class="fa fa-check" v-on:click.prevent="toggle"></button>
    <button class="fa fa-remove" v-on:click.prevent="remove"></button>
  </div>
</template>
