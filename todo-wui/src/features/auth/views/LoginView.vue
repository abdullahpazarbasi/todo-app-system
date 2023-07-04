<script setup lang="ts">
import {inject, ref} from "vue";

import TopNavBar from "@/core/components/TopNavBar.vue";
import type AuthUseCase from "@/features/auth/usecases/AuthUseCase";

const authUseCase: AuthUseCase | undefined = inject('authUseCase');
const username = ref("");
const password = ref("");

function login() {
  if (username.value.length > 0 && password.value.length > 0) {
    authUseCase!.login(username.value, password.value);
  }
}
</script>

<template>
  <TopNavBar>
    <router-link to="/">Home</router-link>
  </TopNavBar>
  <main>
    <form v-on:submit.prevent="login">
      <input type="text" v-model="username"/>
      <input type="password" v-model="password"/>
      <button type="submit">Login</button>
    </form>
  </main>
</template>
