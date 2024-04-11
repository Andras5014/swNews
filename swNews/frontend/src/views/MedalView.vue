<template>
    <PageHeader title="Medal" img-src="/medal/header.jpg"></PageHeader>
    <WidthWrapper class="my-10">
        <n-data-table striped class="mt-5" 
        row-class-name="medal-row"
        :columns="columns" :data="data"  size="large" :bordered="false" />

        <WorldMap :medalArr="data">

        </WorldMap>
    </WidthWrapper>
</template>

<script setup lang="ts">
import PageHeader from '@/components/PageHeader.vue'
import WidthWrapper from '@/components/WidthWrapper.vue'
import WorldMap from '@/components/WorldMap.vue'
import type { DataTableBaseColumn } from 'naive-ui'
import { getMedal, type Medal } from '@/api/medal'
import { ref, onBeforeMount, h } from 'vue'
const genFlagUrl = (countryCode: string) => {
  return `https://www.worldaquatics.com/resources/v2.11.7/i/elements/flags/${countryCode.toLowerCase()}.png`
}
const columns:DataTableBaseColumn<Medal>[] = [
  {
    title: 'Over Rank',
    key: 'rank',
    width: '140px'
  },
  {
    title: 'Country',
    key: 'country',
    render: (row) => {
      return h('div', {class: 'flex items-center', style: 'padding-left:calc(50% - 80px)'}, [
        h('img', {src: genFlagUrl(row.countryCode), alt: row.countryCode, class: 'w-6 h-6 mr-2'}),
        h('span', {class: 'ml-2'}, row.country)
      ])
    }
  },
  {
    title: 'Gold',
    key: 'gold',
    className: 'gold-header',
    maxWidth: '200px'
  },
  {
    title: 'Silver',
    key: 'silver',
    className: 'sliver-header',
    maxWidth: '200px'
  },
  {
    title: 'Bronze',
    key: 'bronze',
    className: 'bronze-header',
    maxWidth: '200px'

  },
  {
    title: 'Total',
    key: 'totalCount',
  }
]
const data = ref<Medal[]>([])

onBeforeMount( async () => {
  const medalList: Medal[] = await getMedal()
  data.value = medalList
})
</script>
<style lang="scss" scoped>
:deep .n-data-table-th {
  text-align: center;
}

:deep .n-data-table-td {
  text-align: center;
}
:deep(.medal-row) {
  height: 80px;
}
:deep(.gold-header) {

  
  background-image: linear-gradient(108deg,#fedb77,#ffb819);
  
}
:deep(.sliver-header) {
  background-image: linear-gradient(108deg,#d9d9d9,#c9c9c9);
}
:deep(.bronze-header) {
  background-image: linear-gradient(108deg,#d8c28d,#a58844);
}
</style>
