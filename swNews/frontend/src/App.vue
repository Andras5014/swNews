<template>
  <n-config-provider :theme-overrides="themeOverrides">
    <n-message-provider>
    <LayoutView ></LayoutView>
  </n-message-provider>
  </n-config-provider>
</template>
<script setup lang="ts">
import LayoutView from '@/layout/LayoutView.vue';
import { useTheme } from '@/stores/theme';
const { themeOverrides } = useTheme()
</script>

<style scoped>

</style>
