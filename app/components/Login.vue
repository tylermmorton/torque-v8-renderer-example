<template>
  <div>
    <h1>{{ props.message }}</h1>
    <input type="text" placeholder="Enter your email" />
    <input type="password" placeholder="Enter your password" />
    <button>Login</button>
  </div>
</template>

<script lang="ts" setup>
import { defineProps } from 'vue'

const props = defineProps<{
  message: string
}>()
</script>