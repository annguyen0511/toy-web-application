<template>
  <pre>{{ meals }}</pre>
</template>

<script setup>

import store from '@/store';
import { computed } from 'vue';

const meals =computed(() => store.state.meals)

</script>