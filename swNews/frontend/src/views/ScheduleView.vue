<template>
  <PageHeader title="Schedule" img-src="/schedule/header.jpg"></PageHeader>
  <WidthWrapper class="mt-10">
    <n-grid y-gap="12" x-gap="12" cols="2 800:3">
      <n-gi>
        <div class="w-full flex justify-center">
          <n-date-picker v-model:value="pickedDate" clearable type="date" />
        </div>
      </n-gi>
      <n-gi>
        <div class="flex items-center w-full justify-center">
          <span class="text-gray-400 mr-2">Phase:</span>
          <n-select class="w-40" v-model:value="matchType" :options="typeOptions" />
        </div>
      </n-gi>
      <n-gi>
        <div class="flex justify-center w-full">
          <n-button quaternary type="primary" @click="resetFilter">
            Reset
          </n-button>
        </div>
      </n-gi>
    </n-grid>
    <template v-if="filteredScheduleListItems.length>0">
      <div class="mt-10"  v-for="(item, index ) in filteredScheduleListItems" :key=index>
      <p class="text-3xl mb-4">{{ new Date(item.date).toDateString() }} <span class="text-gray-400">January</span></p>
      <ScheduleItem v-for="(schedule, i) in item.schedules" :key="i" 
        :schedule="schedule" class="mb-4 shadow"></ScheduleItem>
    </div>
    </template>
    <n-empty class="my-10" v-else description="暂无数据">
    <template #extra>
    </template>
  </n-empty>
  </WidthWrapper>
</template>

<script setup lang="ts">
import PageHeader from '@/components/PageHeader.vue'
import WidthWrapper from '@/components/WidthWrapper.vue'
import ScheduleItem from '@/components/ScheduleItem.vue'
import { ref, onBeforeMount, watch } from 'vue'
import { getSchedule, type Schedule } from '@/api/schedule'

const pickedDate = ref(null)
const matchType = ref('All')
const typeOptions = [
  { label: 'All', value: 'All' },
  { label: 'Finals', value: 'Finals' },
]

watch([pickedDate, matchType], () => {
  doFilter()
})

interface ScheduleListItem {
  date: string

  schedules: Schedule[]
}
const scheduleListItems = ref<ScheduleListItem[]>([])
const filteredScheduleListItems = ref<ScheduleListItem[]>([])
const resetFilter = () => {
  matchType.value = 'All'
  pickedDate.value = null
  doFilter()
}

const doFilter = () => {
  filteredScheduleListItems.value = []
  scheduleListItems.value.forEach((item) => {
    if (pickedDate.value && new Date(item.date).toDateString() !== new Date(pickedDate.value).toDateString()) {
      return
    }
    filteredScheduleListItems.value.push({
      date: item.date,
      schedules: [...item.schedules]

    })
  })

  filteredScheduleListItems.value.forEach((item) => {
    item.schedules = item.schedules.filter((schedule) => {
      if (matchType.value != 'All' && schedule.phaseName !== matchType.value) {
        return false
      }
      return true
    })
  })

}
onBeforeMount(async () => {
  const schedules: Schedule[] = await getSchedule()
  schedules.forEach((schedule) => {
    const date = schedule.date
    const index = scheduleListItems.value.findIndex((item) => item.date === date)
    if (index === -1) {
      scheduleListItems.value.push({
        date,
        schedules: [schedule]
      })
    } else {
      scheduleListItems.value[index].schedules.push(schedule)
    }
  })
  scheduleListItems.value.sort((a, b) => {
    return new Date(a.date).getTime() - new Date(b.date).getTime()
  })
  scheduleListItems.value.forEach((item) => {
    item.schedules.sort((a, b) => {
      return a.time > b.time ? 1 : -1
    })
  })

  doFilter()
})

</script>