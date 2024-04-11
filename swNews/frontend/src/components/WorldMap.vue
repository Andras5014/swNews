<template>
  <div class="flex flex-col items-center">
    <h2 class="mt-10">奖牌环形图</h2>
    <div  id="world-map"></div>
  </div>
  
</template>

<script lang="ts" setup>
import { nextTick, watch } from 'vue'
import type { Medal } from '@/api/medal'
import * as echarts from 'echarts'
const props = defineProps<{
  medalArr: Medal[]
}>() 
const getOption = () => {
  return {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      top: '5%',
      left: 'center'
    },
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        padAngle: 5,
        itemStyle: {
          borderRadius: 10
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 40,
            fontWeight: 'bold'
          }
        },
        labelLine: {
          show: false
        },
        data: props.medalArr.map((item) => {
          return {
            value: item.totalCount,
            name: item.country
          }
        })
      }
    ]
  }
}
watch(() => props.medalArr, () => {
  if (props.medalArr.length === 0) {
    return
  }
  nextTick(() => {
    const chart = echarts.init(document.getElementById('world-map') as HTMLDivElement)
    chart.setOption(getOption())
  })
}, {immediate: true} )

</script>
<style lang="scss" scoped>
#world-map {
  height:500px;
  width: 100%;
}
</style>