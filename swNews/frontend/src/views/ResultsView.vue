<template>
    <PageHeader title="Results" img-src="/results/header.jpg"></PageHeader>
    <WidthWrapper class="mt-20" >
        <h2 class="text-4xl font-thin text-slate-500 mb-5">Diving</h2>
        <ResultItem v-for="(item, index) in resultList" 
        class="mb-4"
        @openDetail="switchToDetail(index)"
        :key="index"
        :id="item.id"
        :ref="(el) => setItemRef(el)"
        :discipline-name="item.disciplineName"></ResultItem>
    </WidthWrapper>
</template>

<script setup lang="ts">
import PageHeader from '@/components/PageHeader.vue'
import WidthWrapper from '@/components/WidthWrapper.vue'
import ResultItem from '@/components/ResultItem.vue'
import { ref, onBeforeMount, watch, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { getResults, type Result } from '@/api/results'
const resultList = ref<Result[]>([])
const route = useRoute()

onBeforeMount(() => {
  loadResults()
})

type LoadedCB = () => void
const loadedCallback:LoadedCB[] = []
const loadResults = async () => {
  const resultData: Result[] = await getResults()
  resultList.value = resultData
  loadedCallback.forEach((cb) => {
    cb()
  })
}

type ResultItemRef = typeof ResultItem
const resultItemRefs:(ResultItemRef)[] = []
const setItemRef = (el:any) => {
  if (el) {
    resultItemRefs.push(el)
  }
}

const switchToDetail = (index:number) => {
  resultItemRefs.forEach((item, i) => {
    if (i === index) {
      item.openDetail()
    } else {
      item.closeDetail()
    }
  })
}

const exclusiveOpen = (id:any) => {
  resultItemRefs.forEach((item, index) => {
    if (resultList.value[index].id == id) {
      item.openDetail()
    } else {
      item.closeDetail()
    }
  })
}

watch(() => route.params.id, () => {
  if (!route.params.id) {
    resultItemRefs.forEach((item) => {
      item.closeDetail()
    })
    return
  }
  if (resultItemRefs.length>0) {
    exclusiveOpen(route.params.id)
  } else {
    loadedCallback.push(() => {
      nextTick(() => {
        exclusiveOpen(route.params.id)
      })
    })
  }
  

}, {immediate: true} )
</script>