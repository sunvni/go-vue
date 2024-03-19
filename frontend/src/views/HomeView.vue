<script setup>
import TweetList from '@/components/TweetList.vue';
import axios from 'axios';
import { onMounted, ref } from 'vue';

const tweetContent = ref('');
const tweetTitle = ref('');
const tweets = ref([]);


function tweet() {
  axios.post('/api/v1/tweet', {
    content: tweetContent.value,
    title: tweetTitle.value
  })
    .then(() => {
      tweetContent.value = '';
      getAllTweets();
    });
}

onMounted(() => {
  getAllTweets()
});

function getAllTweets() {
  axios.get('/api/v1/tweets')
    .then(response => {
      tweets.value = response.data.tweets;
    });
}

</script>

<template>
  <main class="container mx-auto bg-gray-900 p-4 rounded-xl">

    <div class="bg-gray-800 rounded p-4">
      <div>
        <label for="title">Tweet Title</label>
        <input id="title" class="w-full bg-gray-700 rounded p-2 mt-1" v-model="tweetTitle" />
      </div>
      <div>
        <label for="tweet">Tweet Content</label>
        <textarea id="tweet" class="w-full bg-gray-700 rounded p-2 mt-1" rows="4" v-model="tweetContent"></textarea>
      </div>
      <div>
        <button class="bg-blue-600 hover:bg-blue-700 active:bg-blue-800 rounded px-3 py-2 mt-2"
          @click="tweet">Tweet</button>
      </div>
    </div>
    <div class="mt-4">
      <TweetList :tweets="tweets"></TweetList>
    </div>
  </main>
</template>
