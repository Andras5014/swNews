<template>
  <n-card content-class="card" size="huge">
    <div class="flex items-center text-2xl">
      <span class="mr-4">{{ props.schedule.time }}</span>
      <n-divider vertical />
      <span class="ml-4">{{ props.schedule.disciplineName }}</span>
      <span class="mx-2 text-gray-400">•</span>
      <span :class="props.schedule.phaseName == 'Finals' ? ['text-amber-400'] : ['text-gray-400']">{{ props.schedule.phaseName }}</span>
      <n-icon v-if="props.schedule.phaseName == 'Finals'" class="text-amber-400 ml-2">
        <Medal />
      </n-icon>
      <RouterLink class="flex items-center hover-trigger text-gray-400 ml-auto" :to="`/results/${props.schedule.id}?name=${props.schedule.name}`">
        <span class="hoverd text-base">EVENT DETAILS</span>
        <n-icon class="text-gray-400 ml-2 ">
          <ArrowForwardOutline />
        </n-icon>
      </RouterLink>
    </div>

  </n-card>
</template>
<script setup lang="ts">
import { Medal, ArrowForwardOutline } from '@vicons/ionicons5'
import { useTheme } from '@/stores/theme'
import type { Schedule } from '@/api/schedule'
import type { PropType } from 'vue'
const { themeColor } = useTheme()
const props = defineProps({
  schedule: {
    type: Object as PropType<Schedule>,
    required: true
  }
})
</script>
<style lang="scss" scoped>
:deep(.card) {
  position: relative;

  &::after {
    content: '';
    position: absolute;
    left: 0;
    top: calc(50% - 15px);
    height: 30px;
    width: 4px;
    background-color: v-bind('themeColor');
  }
}

.hoverd {
  background-image: linear-gradient(rgba(13, 79, 173, .1), rgba(13, 79, 173, .1));
  background-repeat: no-repeat;
  background-size: 0 38%;
  background-position: 0 100%;
  transition: background-size .35s ease;
}

.hover-trigger:hover {
  .hoverd {
    background-size: 100% 38%;
  }
}
</style>