<script setup lang="ts">
import type TodoCandidate from "@/features/todo/entities/TodoCandidate";

const props = defineProps<{
    candidate: TodoCandidate
}>();

const emit = defineEmits<{
    (e: 'candidate-value-set', todoText: string): void
}>();

function add(event: Event) {
    const value = (event.target as HTMLInputElement).value.trim();
    if (value) {
        emit('candidate-value-set', value);
        (event.target as HTMLInputElement).value = '';
    }
}
</script>

<template>
    <input type="text" placeholder="???" v-bind:value="props.candidate.value" v-on:keyup.enter="add" autofocus/>
</template>
