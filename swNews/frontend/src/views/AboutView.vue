<template>
  <PageHeader title="More" img-src="/about/header.jpg"></PageHeader>
  <WidthWrapper class="mt-6">
    <h2>Latest news</h2>
    <n-grid x-gap="12" cols="2 800:4">
      <n-gi v-for="(item, index) in newsList " :key="index">
        <a :href="item.href" target="_blank">
          <n-card :title="item.title">
          <template #cover>
            <img :src="parseUrl(item.cover)">
          </template>
        </n-card>
        </a>
      </n-gi>

    </n-grid>
  </WidthWrapper>
  <WidthWrapper class="mt-28">
    <iframe width="100%" height="600" src="https://www.youtube.com/embed/sVrRt8KcBbA?si=MfRkhQfcKtwii_Y3" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
  </WidthWrapper>
  <div class="gardient-block">
    <WidthWrapper>
      <n-grid x-gap="12" y-gap="20" cols="1 800:2" class="min-h-40">
        <n-gi>
          <div class="flex items-center h-full">
            <h2 >Diving World Cup 2024</h2>
          </div>
        </n-gi>
        <n-gi>
          <div class="flex items-center h-full">
            <p>2024 Calendar: 🇨🇦 Montreal - 29 Feb / 3 Mar 🇩🇪 Berlin - 21/24 Mar 🇨🇳 Xi’an Super Final -19/21 Apr More info - https://www.worldaquatics.com/series/166/world-aquatics-diving-world-cup-2024</p>
          </div>
        </n-gi>
      </n-grid>
    </WidthWrapper>
  </div>

  <WidthWrapper class="mt-28">
    <iframe width="100%" height="600" src="https://www.youtube.com/embed/ArUlm6Xq9e4?si=ieK2cjopZr84xBNZ" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
  </WidthWrapper>
  <div class="gardient-block">
    <WidthWrapper>
      <n-grid x-gap="12" y-gap="20" cols="1 800:2" class="min-h-40">
        <n-gi>
          <div class="flex items-center h-full">
            <h2 >World Aquatics 2024 GlobalCalendar</h2>
          </div>
        </n-gi>
        <n-gi>
          <div class="flex items-center h-full">
            <p>🌎🗓 World Aquatics 2024 Global calendar looking on 🔥 Be ready to witness history acrossthe globe in this amazing year 🤩 #AQUA #unitedbywater ▶▶ SUBSCRIBE to our YouTubechannel: ▶▶ htps://www.youtube.com/channel/UCZko_COpTM4_d2kPrxZfzFg</p>
          </div>
        </n-gi>
      </n-grid>
    </WidthWrapper>
  </div>
</template>
<script setup lang="ts">
import PageHeader from '@/components/PageHeader.vue'
import WidthWrapper from '@/components/WidthWrapper.vue'
import { ref } from 'vue'
const newsList = ref([
  {
    title: 'Seeking Talented Scholar Athletes: Applications now open for… ',
    cover: 'https://resources.fina.org/photo-resources/2024/03/14/106becba-9649-4262-8bb7-04ee449d7de4/GettyImages-72699473.jpg?width=342&height=192',
    href: 'https://www.worldaquatics.com/news/3932906/scholarship-athletes-swimming-applications-open-hungary-university-programme-world-aquatics'
  },
  {
    title: 'Generations at the Pool: Cuban and Venezuelan family swimmers…',
    cover: 'https://resources.fina.org/photo-resources/2024/02/19/67c61e15-db52-454b-bad6-f65618640dd1/GettyImages-2017214252.jpg?width=342&height=192',
    href: 'https://www.worldaquatics.com/news/3902424/generations-at-the-pool-cuban-and-venezuelan-family-swimmers-pursue-olympic-dreams-and-heritage-world-aquatics-championships-doha-2024'
  },
  {
    title: 'United States wins eighth gold medal on final night in Doha;… ',
    cover: 'https://resources.fina.org/photo-resources/2024/02/18/62c9e51f-7764-4c66-a0bb-14c9f5cdb8e7/GettyImages-2020714874.jpg?width=342&height=192',
    href: 'https://www.worldaquatics.com/news/3902804/united-states-wins-eighth-gold-medal-on-final-night-in-doha-wiffen-sjostrom-scare-world-records-world-aquatics-championships-doha-2024'
  },
  {
    title: 'They Said | Day 17 (18 Feb) at the World Aquatics Championships -… ',
    cover: 'https://resources.fina.org/photo-resources/2024/02/18/40715f77-5663-4faf-bdf8-64f308c884e5/GettyImages-2019540021.jpg?width=342&height=192',
    href: 'https://www.worldaquatics.com/news/3902913/they-said-athlete-quotes-world-aquatics-championships-doha-2024'
  }
])
const parseUrl = (src: string) => {
  return src.replace(/^@/, '/src/')
}
</script>
<style lang="scss" scoped>
.gardient-block {
  color: white;
  background: linear-gradient(180deg,hsla(0,0%,100%,.3),hsla(0,0%,100%,0)),linear-gradient(180deg,#0d4fad,#0d4fad);
}
</style>
