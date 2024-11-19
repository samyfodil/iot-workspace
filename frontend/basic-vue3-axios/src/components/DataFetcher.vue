<template>
  <div>
    <h2>Data Fetcher</h2>
    <button @click="fetchData">Fetch Data</button>
    <div v-if="loading">Loading...</div>
    <ul v-else>
      <li v-for="item in data" :key="item.id">
        {{ item.title }}
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import axios from 'axios';

interface Post {
  id: number;
  title: string;
}

export default defineComponent({
  name: 'DataFetcher',
  setup() {
    const data = ref<Post[]>([]);
    const loading = ref(false);

    const fetchData = async () => {
      loading.value = true;
      try {
        const response = await axios.get<Post[]>('https://jsonplaceholder.typicode.com/posts');
        data.value = response.data.slice(0, 5); // Fetch only the first 5 items
      } catch (error) {
        console.error('Error fetching data:', error);
      } finally {
        loading.value = false;
      }
    };

    return {
      data,
      loading,
      fetchData,
    };
  }
});
</script>

<style scoped>
button {
  padding: 10px 20px;
  margin: 10px 0;
  font-size: 16px;
  background-color: #42b983;
  color: white;
  border: none;
  cursor: pointer;
}
button:hover {
  background-color: #358a6d;
}
ul {
  list-style: none;
  padding: 0;
}
li {
  background: #f0f0f0;
  margin: 5px 0;
  padding: 10px;
  border-radius: 5px;
}
</style>