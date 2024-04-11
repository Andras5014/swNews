<template>
    <footer class="footer">
        <WidthWrapper class="footer-wrap">
            <div class="left">
                <h2>Headquarters</h2>

                <p>Chemin de Bellevue 24a/24b</p>
                <p>CH - 1005 Lausanne</p>
                <p>Switzerland</p>

                <div class="icons">
                    <div class="icon-wrap">
                        <n-icon size="24">
                            <LogoWeibo />
                        </n-icon>
                    </div>
                    <div class="icon-wrap">
                        <n-icon size="24">
                            <LogoWechat />
                        </n-icon>
                    </div>
                    <div class="icon-wrap">
                        <n-icon size="24">
                            <LogoTwitter />
                        </n-icon>
                    </div>
                </div>
            </div>
            <div class="right">
                <LogoView :height="100" />
                <p>Copyright 2024 by 222100408 && 222100409</p>
            </div>
        </WidthWrapper>
        <!-- <div  :style="widgetStyle">
            
        </div> -->
    </footer>
</template>
<script setup lang="ts">
import WidthWrapper from '@/components/WidthWrapper.vue'
import { Weibo as LogoWeibo } from '@vicons/fa'
import { LogoWechat } from '@vicons/carbon'
import { LogoTwitter } from '@vicons/ionicons5'
import LogoView from '../logo/LogoView.vue'
</script>
<style lang="scss" scoped>
.footer {
    height: 300px;
    background-color: black;
}
.footer-wrap {
    height: 100%;
    color: white;
    display: flex;
    align-items: center;
    justify-content: space-between;
    .left {
        h2 {
        margin-bottom: 30px;
    }
    .icons {
        display: flex;
        margin-top: 30px;
        .icon-wrap {
            height: 40px;
            width: 40px;
            border: 1px solid white;
            border-radius: 50%;
            display: flex;
            justify-content: center;
            align-items: center;
            margin-right: 14px;
        }
    }
    }
    .right {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        p {
            margin-top: 30px;
        }
    }
}
</style>