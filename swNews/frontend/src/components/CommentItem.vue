<template>
<n-card :title="'昵称: ' + props.name" size="small">
    <div class="text-lg">
        {{ props.content }}
    </div>
    <template #footer>
      <div class="text-right">
        {{ new Date(props.time * 1000).toLocaleString() }}
      </div>
    </template>
  </n-card>
</template>
<script setup lang="ts">
const props = defineProps({
  name: String,
  content: String,
  time: {
    type: Number,
    required: true,
  },
})
</script>