<script setup>

defineProps({
    tweet: Object
});


</script>

<template>
    <div class="bg-gray-700 rounded p-4 flex flex-col">
        <div class="flex-1">
            <p class="text-semibold">{{ tweet.Title }}</p>
            <p class="text-red-400">{{ tweet.Content }}</p>
        </div>
        <div class="w-full flex justify-end">
            <p class="text-xs text-gray-400">{{ tweet.CreatedAt }}</p>
        </div>
    </div>
</template>