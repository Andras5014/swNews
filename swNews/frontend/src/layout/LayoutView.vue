<template>
  <n-layout  >
    <n-layout-header bordered class="top-0 z-50"  position="absolute">
      <HeaderView  />
    </n-layout-header>
    <n-layout-content ref="layoutRef">
      <ContentView>
        
      </ContentView>
    </n-layout-content>
    <n-layout-footer>
      <FooterView></FooterView>
    </n-layout-footer>
  </n-layout>
  <n-back-top :right="100"  :listen-to="layoutRef" />
</template>

<script setup lang="ts">
import HeaderView from './header/HeaderView.vue'
import FooterView from './footer/FooterView.vue'
import ContentView from './content/ContentView.vue'
import { ref } from 'vue'
const layoutRef = ref(null)
</script>