<script setup lang="ts">
import {RouterView} from 'vue-router';
import {loading} from "@/core/composables/loading";
</script>

<template>
  <RouterView/>
  <!-- loading -->
  <div v-if="loading" class="loading">
    <div class="spinner"></div>
  </div>
  <!-- /loading -->
</template>
