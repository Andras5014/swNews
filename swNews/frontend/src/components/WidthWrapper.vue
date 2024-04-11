<template>
    <div :style="widgetStyle">
        <slot></slot>
    </div>
</template>
<script lang="ts" setup>
import { useMediaWidth } from '@/stores/mediaWidth'
const { widgetStyle } = useMediaWidth()
</script>