<script setup>
import { ref } from 'vue';
import axios from 'axios';

const count = ref(0);

function increment() {
    axios.post('/api/v1/counter/increment')
        .then(response => {
            count.value += response.data.count;
        });
}
function decrement() {
    axios.post('/api/v1/counter/decrement')
        .then(response => {
            count.value -= response.data.count;
        });
}

</script>

<template>
    <div>
        <h2>Server Counter</h2>
        <p>Count: {{ count }}</p>
        <button @click="increment"
            class="bg-green-600 hover:bg-green-700 active:bg-green-800 rounded px-3 py-2">Increment</button>
        <button @click="decrement"
            class="bg-red-600 hover:bg-red-700 active:bg-red-800 rounded px-3 py-2">Decrement</button>
    </div>
</template>