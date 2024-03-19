<script setup>
import TweetItem from './TweetItem.vue';

defineProps({
    tweets: Array
});

</script>

<template>
    <div class="space-y-2 p-4 bg-gray-800 rounded">
        <h2 class="text-white">Tweets</h2>
        <TweetItem v-for="tweet in tweets" :key="tweet.id" :tweet="tweet"></TweetItem>
    </div>
</template>