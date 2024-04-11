<template>
    <n-carousel draggable autoplay :interval="3000">
        <img class="carousel-img" src="@/assets/carousel/1.jpg">
        <img class="carousel-img" src="@/assets/carousel/2.jpg">
        <img class="carousel-img" src="@/assets/carousel/3.jpg">
        <img class="carousel-img" src="@/assets/carousel/4.jpg">
    </n-carousel>
    <WidthWrapper class="py-6">
        <n-flex align="center" vertical>
            <h2 class="mt-6">世界游泳锦标赛</h2>
            <p class="gray mt-6">
                世界泳联锦标赛（英语：World Aquatics Championships，2022年前为国际泳联世界锦标赛，英语：FINA World
                Championships），又称世界游泳锦标赛，是由世界水上运动（前称国际游泳联合会）主办的最高级别国际性游泳赛事。1973年起开始举行，每2年举行一届，在1978年至1998年曾经每4年举行一届，在2001年开始恢复每2年举行一届。
            </p>
            <n-grid x-gap="20" cols="1 800:2" y-gap="20" class="mt-6">
                <n-gi>
                    <div class="info-img-wrap">
                        <img class="info-img" src="/home/1.png" alt="">
                    </div>
                </n-gi>
                <n-gi>
                    <div class="info-img-wrap">
                        <img class="info-img" src="/home/2.jpg" alt="">
                    </div>
                </n-gi>
            </n-grid>
        </n-flex>
    </WidthWrapper>
    <div class="bg-slate-100">
        <WidthWrapper class="py-6">
            <h2 class="mt-6">2024年世界游泳锦标赛</h2>
            <p class="gray mt-6">
                2024年世界游泳锦标赛，为第21届世界游泳锦标赛，于2024年2月2日至18日在卡塔尔多哈举行[1]。是届赛事原定于2023年秋季举行，因福冈世界游泳锦标赛延期至2023年，本届赛事亦因而顺延至2024年。这是多哈首次举办世界游泳锦标赛，亦是阿拉伯世界第一次有城市举办世界游泳锦标赛。本届赛事距离巴黎奥运开幕只有5个月，这也是世界游泳锦标赛首次在奥运年举行
            </p>
            <n-grid x-gap="20" cols="1 800:3" y-gap="20" class="mt-6">
                <n-gi>
                    <div class="h-64">
                        <img class="info-img" src="/home/3.jpg" alt="">
                    </div>
                </n-gi>
                <n-gi>
                    <div class="h-64">
                        <img class="info-img" src="/home/4.jpg" alt="">
                    </div>
                </n-gi>
                <n-gi>
                    <div class="h-64">
                        <img class="info-img" src="/home/5.jpg" alt="">
                    </div>
                </n-gi>
            </n-grid>
        </WidthWrapper>
    </div>
    <WidthWrapper class="py-6">
        <n-grid x-gap="20" cols="1 800:2" y-gap="20" class="mt-6">
            <n-gi>
                <div class="h-full">
                    <n-flex vertical class="h-full">
                        <h2>2024年世界游泳锦标赛跳水比赛</h2>
                        <n-flex align="center" class="flex-1">
                            <p class="gray">
                                2024年世界游泳锦标赛跳水比赛是2024年世界游泳锦标赛的其中一个比赛项目，于2024年2月2日至10日在卡塔尔多哈的哈马德水上运动中心举行。跳水比赛共产生13枚金牌[1][2]。此次比赛同时也是2024年巴黎奥运跳水项目的资格赛之一，个人小项提供余下未分配的运动员名额，直至入围奥运正赛总人数达到相应数量为止，每个双人小项则提供4个席位[3]。
                            </p>
                        </n-flex>
                    </n-flex>
                </div>
            </n-gi>
            <n-gi>
                <div class="h-80">
                    <img class="info-img" src="/home/6.jpg" alt="">
                </div>
            </n-gi>
        </n-grid>
    </WidthWrapper>

    <div class="bg-neutral-400">
        <WidthWrapper class="py-6">
            <h2 class="text-white text-center">合作伙伴</h2>
            <n-grid x-gap="12" y-gap="20" cols="2 800:4" class="mt-6">
                <n-gi>
                    <a href="https://www.myrthapools.com/zh/" target="_blank">
                        <div class="bg-white h-44">
                        <img class="info-img" src="/home/board/my.jpg" alt="">
                    </div>
                    </a>
                </n-gi>
                <n-gi>
                    <a href="https://en.nongfuspring.com/" target="_blank">
                        <div  class="bg-white h-44">
                        <img class="info-img" src="/home/board/ns.png" alt="">
                    </div>
                    </a>
                </n-gi>
                <n-gi>
                    <a href="https://www.omegawatches.com.hk/zh-hk/" target="_blank">
                        <div class="bg-white h-44" >
                        <img class="info-img" src="/home/board/om.png" alt="">
                    </div>
                    </a>
                </n-gi>
                <n-gi>
                    <a href="https://www.yakult.co.jp/english/" target="_blank">
                        <div  class="bg-white h-44">
                        <img class="info-img" src="/home/board/ya.png" alt="">
                    </div>
                    </a>
                </n-gi>
            </n-grid>
        </WidthWrapper>
    </div>
</template>
<script setup lang="ts">
import WidthWrapper from '@/components/WidthWrapper.vue'

</script>

<style lang="scss" scoped>
.carousel-img {
    width: 100vw;
    height: 600px;

    object-fit: cover;
}


.info-img-wrap {
    height: 400px;
    overflow: hidden;

}

.info-img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.gray {
    color: gray;
}
</style>