<template>
  <div class="select-none cursor-pointer" ref="selfRef">
    <n-card @click="switchDetail" size="huge" class="shadow">
      <n-flex justify="space-between" align="center">
        <span class="text-xl font-medium">{{ props.disciplineName }}</span>
        <n-icon :class="[spread ? 'rotate-90' : '-rotate-90', 'transition duration-300']">
          <ArrowBackIosFilled />
        </n-icon>
      </n-flex>
    </n-card>
    <div v-if="spread" class="flex justify-center">
      <div class="flex items-center border border-solid blue border-gray-200">
        <div :class="['name-item w-60 h-20', activeName == item.name ? 'active' : '']"
          v-for="(item, index) in detailData" :key="index" @click="changeActiveName(item.name)">
          <div class="flex flex-col justify-center items-center">
            <span class="px-2 rounded-b-lg green">OFFICAL</span>
            <span class="text-lg font-semibold">{{ item.name }}</span>
            <span class="text-gray-400">{{ item.utcTime }}</span>
          </div>
        </div>
      </div>
    </div>
    <div v-show="spread">
      <n-data-table striped class="mt-5" row-class-name="athlete-row" :columns="columns" :data="getResList(activeName)"
        ref="tableRef" size="large" :bordered="false" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, h, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { DataTableBaseColumn } from 'naive-ui'
import { ArrowBackIosFilled } from '@vicons/material'
import { getResultDetail, type ResultDetail, type AthleteGroup, type Athlete } from '@/api/results'
const spread = ref(false)
const route = useRoute()
const router = useRouter()
const props = defineProps({
  disciplineName: {
    type: String,
    required: true
  },
  id: {
    type: String,
    required: true
  }
})
const emit = defineEmits(['closeDetail', 'openDetail'])
const genFlagUrl = (countryCode: string) => {
  return `https://www.worldaquatics.com/resources/v2.11.7/i/elements/flags/${countryCode.toLowerCase()}.png`
}

const renderAthleteName = (row: Athlete) => {
  return h('div', { class: 'flex', style: 'padding-left:calc(50% - 80px)' }, [
    h('img', {
      src: 'https://resources.fina.org/photo-resources/2023/09/15/3e56c238-f35c-476e-a277-c7f5ee26536d/d19ef7aa-7847-4f1d-9c4c-dbdc6a6b0425?width=80',
      alt: 'athlete',
      class: 'w-12 h-12 rounded-full mr-4'}),
    h('div', { class: 'flex flex-col justify-center items-start' }, [
      h('div', row.firstName),
      h('div', { class: 'font-semibold' }, row.lastName)
    ])
  ])
}
const columns: DataTableBaseColumn<AthleteGroup>[] = [
  {
    title: 'Overall Rank',
    key: 'rank',
    width: '140px'
  },
  {
    title: 'Country',
    key: 'countryName',
    width: '20%',
    render: (row) => {
      return h('div', { class: 'flex items-center justify-center' }, [
        h('img', { src: genFlagUrl(row.countryName), alt: row.countryName, class: 'w-6 h-6 mr-2' }),
        h('span', { class: 'ml-2' }, row.countryName)
      ])
    }
  },
  {
    title: 'Athlete',
    key: 'athlete',
    render: (row) => {
      if (row.athletes.length === 2) {
        return h('div', {}, [
          renderAthleteName(row.athletes[0]),
          h('div', {class:'h-5'}),
          renderAthleteName(row.athletes[1])
        ])
      } else {
        return renderAthleteName(row.athletes[0])
      }
    }
  },
  {
    title: 'Age',
    key: 'age',
    render: (row) => {
      if (row.athletes.length === 2) {
        return h('div', {}, [
          h('div', {}, row.athletes[0].age),
          h('div', {class:'h-4'}),
          h('div', {}, row.athletes[1].age)
        ])
      } else {
        return h('div', {}, row.athletes[0].age)
      }
    }
  },
  {
    title: 'Points',
    key: 'points',
  },
  {
    title: 'Pts Behind',
    key: 'pointBehind',
  }

]
const detailData = ref<ResultDetail[]>([])


const getResList = (name: string) => {
  return detailData.value.find(item => item.name === name)?.athleteGroups || []
}
const activeName = ref('')
const changeActiveName = (name: string) => {
  router.replace(`/results/${props.id}?name=${name}`)
  activeName.value = name
}
const selfRef = ref<HTMLElement | null>(null)
const closeDetail = () => {
  if (!spread.value) return
  emit('closeDetail')
  spread.value = false
}
const openDetail = () => {
  if (spread.value) return
  nextTick(() => {
    selfRef.value?.scrollIntoView({ behavior: 'smooth', block: 'center'})
  })
  spread.value = true
  if (detailData.value.length === 0) {
    getResultDetail(props.id).then(res => {
      detailData.value = res
      if (route.query.name) {
        activeName.value = route.query.name as string
      } else {
        activeName.value = res[0].name
      }
    })
  }
  emit('openDetail')

}
const switchDetail = () => {
  if (spread.value) {
    router.replace('/results')
  } else {
    router.replace(`/results/${props.id}?name=${activeName.value}`)
  }
}
defineExpose({
  closeDetail,
  openDetail
})
</script>

<style lang="scss" scoped>
:deep .n-data-table-th {
  text-align: center;
}

:deep .n-data-table-td {
  text-align: center;
}

.green {
  color: #408428;
  background-color: rgba(64, 132, 40, .2);
}

.blue {
  background-color: #ebeef057;
}

.name-item {
  cursor: pointer;
  position: relative;

  &::after {
    content: "";
    height: 0.4rem;
    position: absolute;
    right: 1.6rem;
    bottom: 0;
    left: 1.6rem;
    background-color: #0d4fad;
    transform: scaleY(0);
    transform-origin: bottom;
    transition: transform .35s ease;
  }

  &:hover::after {
    transform: scaleY(1);
  }
}

.active {
  &::after {
    transform: scaleY(1);
  }
}
</style>