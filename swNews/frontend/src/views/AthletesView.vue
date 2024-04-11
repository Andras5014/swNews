<template>
  <PageHeader title="Athletes" img-src="/athletes/header.jpg"></PageHeader>
  <WidthWrapper class="my-10">
    <div class="flex">
      <div class="flex items-center w-40 ml-10">
        <span class="text-gray-400 mr-2">Gender:</span>
        <n-select v-model:value="gender" :options="genderOptions" />
      </div>
      <div class=" flex items-center w-40 ml-20">
        <span class="text-gray-400 mr-2">Country:</span>
        <n-select v-model:value="country" :options="countryOptions" />
      </div>
    </div>
    <n-data-table striped class="mt-5" :columns="columns" :data="data" ref="tableRef" size="large"
      row-class-name="athlete-row" :bordered="false" />
  </WidthWrapper>
</template>

<script setup lang="ts">
import PageHeader from '@/components/PageHeader.vue'
import WidthWrapper from '@/components/WidthWrapper.vue'
import { useTheme } from '@/stores/theme';
import type { DataTableBaseColumn, NDataTable } from 'naive-ui'
import { getAthletes, type Athlete } from '@/api/athletes'
import { ref, watch, onBeforeMount, h } from 'vue'
const { themeColor } = useTheme()
const gender = ref('All')
const country = ref('All')
type DataTable = InstanceType<typeof NDataTable>
const tableRef = ref<DataTable | null>(null)
const genderOptions = [
  { label: 'All', value: 'All' },
  { label: 'Men', value: 'Male' },
  { label: 'Women', value: 'Female' }
]
const countryOptions = ref([
  {
    label: 'All',
    value: 'All'
  }
])
const data = ref<Athlete[]>([])
const genFlagUrl = (countryCode: string) => {
  return `https://www.worldaquatics.com/resources/v2.11.7/i/elements/flags/${countryCode.toLowerCase()}.png`
}
onBeforeMount(async () => {
  const athleteList: Athlete[] = await getAthletes()
  const countrySet = new Set<string>()
  athleteList.forEach((athlete, index) => {
    athlete.dob = new Date(athlete.dob).toLocaleDateString()
    if (!countrySet.has(athlete.countryCode)) {
      countrySet.add(athlete.countryCode)
      countryOptions.value.push({
        label: athlete.countryCode,
        value: athlete.countryCode
      })
    }
  })
  data.value = athleteList
})
const columns: DataTableBaseColumn<Athlete>[] = [
  {
    title: 'Country',
    key: 'country',
    width: '20%',
    render: (row: Athlete) => {
      return h('div', { class: 'flex items-center justify-center' }, [
        h('img', { src: genFlagUrl(row.countryCode), alt: row.countryCode, class: 'w-6 h-6 mr-2' }),
        h('span', { class: 'ml-2' }, row.countryCode)
      ])
    },
    filter: (v, row) => {
      return row.countryCode === v
    }
  },
  {
    title: 'Athlete',
    key: 'athlete',
    render: (row: Athlete) => {
      return h('div', { class: 'flex', style: 'padding-left:calc(50% - 80px)' }, [
        h('img', {

          src: row.gender=='Male' ?
            'https://resources.fina.org/photo-resources/2023/09/15/3e56c238-f35c-476e-a277-c7f5ee26536d/d19ef7aa-7847-4f1d-9c4c-dbdc6a6b0425?width=80':
            'https://resources.fina.org/photo-resources/2024/01/24/49f3ba3b-6b8d-4278-acca-c39bb1a38f05/ff555f81-f026-4493-a1bd-78e84ce7bcd8?width=80',
          alt: 'athlete',
          class: 'w-12 h-12 rounded-full mr-4'
        }),
        h('div', { class: 'flex flex-col justify-center items-start' }, [
          h('div', row.firstName),
          h('div', { class: 'font-semibold' }, row.lastName)
        ])
      ])
    },
  },
  {
    title: 'Gender',
    key: 'gender',
    width: '20%',
    filter: (v, row) => {
      return row.gender === v
    }
  },
  {
    title: 'DOB',
    key: 'dob',
    width: '20%'
  }

]

watch([gender, country], () => {

  if (gender.value === 'All' && country.value === 'All') {
    tableRef.value?.clearFilter()
  }
  else {
    
    tableRef.value?.filter({
      country: country.value !== 'All' ? [country.value] : undefined,
      gender: gender.value!=='All'?[gender.value]:undefined
    })
  
  }
})
</script>

<style lang="scss" scoped>
:deep .n-data-table-th {
  text-align: center;
}

:deep .n-data-table-td {
  text-align: center;
}

:deep(.athlete-row) {
  height: 80px;
  position: relative;

  &::after {
    content: '';
    position: absolute;
    top: 25px;
    left: 0;
    width: 4px;
    height: 30px;
    background-color: v-bind('themeColor');
  }
}
</style>