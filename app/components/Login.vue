<template>
  <div>
    <p>
      {{props.message}}
      <br>
      {{props.count}}
      <button @click="count++">+</button>
    </p>
  </div>
</template>

<script lang="ts" setup>
import {onMounted} from 'vue'

const props = defineProps<{
  message: string;
  count: number;
}>()

onMounted(()=>{
  console.log('mounted')
})
</script>