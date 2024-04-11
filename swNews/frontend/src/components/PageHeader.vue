<template>
    <div class="page-header" :style="{backgroundImage: genBgImgUrl(props.imgSrc)}">
        <p>{{ props.title }}</p>
    </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
const props = defineProps({
  title: String,
  imgSrc: {
    type: String,
    default: ''
  },
})
const genBgImgUrl = (imgSrc: string) => {
  const newSrc = imgSrc.replace(/^@/, '/src/')
  return `url("${newSrc}")`
}
</script>

<style lang="scss" scoped>
.page-header {
    width: 100%;
    height: 600px;
    background-size: cover;
    background-position: center;
    display: flex;
    justify-content: center;
    align-items: center;
    p {
        font-size: 58px;
        color: white;
        text-shadow: 0 0 8px black;
    }
}
</style>