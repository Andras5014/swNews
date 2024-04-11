<template>
    <PageHeader title="Message" img-src="/comment/header.jpg"></PageHeader>
    <WidthWrapper>
        <div class="max-w-3xl my-10">
            <n-input v-model:value="comment" type="textarea" placeholder="留下你的看法" />
            <div class="flex mt-4">
                <n-input v-model:value="name" type="text" placeholder="留下你的昵称" />
                <n-button class="ml-10" @click="sendComment" strong secondary round type="primary">
                提交
            </n-button>
            </div>
        </div>

        <div>
            <CommentItem v-for="(item, index) in commentList.list" 
            :key="index" :name="item.name" :content="item.content" :time="item.time"/>
            <n-pagination class="my-10" v-model:page="pageNo" 
            :on-update:page="handlePageChange"
            :page-count="commentList.totalPage" />
        </div>
    </WidthWrapper>
</template>

<script setup lang="ts">
import PageHeader from '@/components/PageHeader.vue'
import WidthWrapper from '@/components/WidthWrapper.vue'
import CommentItem from '@/components/CommentItem.vue'
import { useMessage } from 'naive-ui'
import { ref, onBeforeMount } from 'vue'
import {getCommentList, addComment, type CommentList} from '@/api/comment'
const comment = ref('')
const name = ref('')
const pageNo = ref(1)
const commentList = ref<CommentList>({list: [], totalPage: 0, pageSize: 10, pageNo: 1})

const loadComments = async (page?:number ) => {
  if (page) {
    pageNo.value = page
  } 
  const  res = await getCommentList(pageNo.value)  
  commentList.value = res
}
const handlePageChange = (pageNo:number) => {
  loadComments(pageNo)
}
onBeforeMount(() => {
  loadComments()
})
const message = useMessage()
const sendComment = async () => {
  if (comment.value === '' || name.value === '') {
    message.error('昵称和评论不能为空')
    return
  }
  await addComment({name: name.value, content: comment.value, time: 0})
  message.success('评论成功')
  comment.value = ''
  name.value = ''
  loadComments()
}
</script>