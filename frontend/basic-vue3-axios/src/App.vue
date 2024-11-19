<template>
  <div id="app">
    <header>
      <h1>Welcome to the Basic Vue 3 Axios TypeScript Project</h1>
    </header>
    <DataFetcher />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import DataFetcher from './components/DataFetcher.vue';

export default defineComponent({
  name: 'App',
  components: {
    DataFetcher
  }
});
</script>

<style>
#app {
  text-align: center;
  margin-top: 50px;
  font-family: Arial, sans-serif;
  color: #333;
}
</style>
