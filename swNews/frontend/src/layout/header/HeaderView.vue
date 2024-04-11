<template>
<WidthWrapper class="header">
  <LogoView :height="40" :color="'black'"></LogoView>
    <div class="menu-warp">
      <n-menu 
      :style="{'font-size': '16px', 'font-weight': 'bold'}"
      :value="route.name"
      :options="menuOptions" mode="horizontal">
        
    </n-menu>
    </div>
</WidthWrapper>

</template>
<script setup lang="ts">
import LogoView from '../logo/LogoView.vue'
import WidthWrapper from '@/components/WidthWrapper.vue'
import { h } from 'vue'
import type { MenuOption } from 'naive-ui'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()
const renderR = (lable: string, params: string) => {
  return () => h(RouterLink, { to: params }, { default: () => lable })
}
const menuOptions: MenuOption[] = [
  {
    label: renderR('首页', '/'),
    key: 'home',
  },
  {
    label: renderR('运动员', '/athletes'),
    key: 'athletes',
  },
  {
    label: renderR('每日赛程', '/schedule'),
    key: 'schedule',
  },
  {
    label: renderR('详细赛况', '/results'),
    key: 'results',
  },
  {
    label: renderR('奖牌榜', '/medal'),
    key: 'medal',
  },
  {
    label: renderR('了解更多', '/about'),
    key: 'about',
  },
  {
    label: renderR('留言', '/comment'),
    key: 'comment',
  },

]
</script>


<style lang="scss" scoped>
.header {
  height: 80px;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

</style>