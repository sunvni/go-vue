<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="h-screen flex flex-col text-white">
    <header class="">
      <div class="container mx-auto">
        <h1 class="text-3xl bg-gray-900 px-10 py-4 rounded">Hello Go - Vue</h1>
        <nav class="space-x-2 mt-4">
          <RouterLink class="bg-gray-700 hover:bg-gray-800 px-4 py-2 rounded" to="/">Home</RouterLink>
          <RouterLink class="bg-gray-700 hover:bg-gray-800 px-4 py-2 rounded" to="/about">About</RouterLink>
        </nav>
      </div>
    </header>

    <RouterView class="flex-1 mt-10" />
  </div>
</template>