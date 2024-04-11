import {defineStore} from 'pinia'
import { ref, computed } from 'vue'
import { lighten } from '@/utils/colorFormat'
import type { GlobalThemeOverrides } from 'naive-ui'
export const useTheme = defineStore('theme', () => {
  const themeColor = ref('#0282c6')

  const themeOverrides = computed(() => {
    const lightenedColor = lighten(themeColor.value, 6)

    const overrides: GlobalThemeOverrides = {
      common: {
        primaryColor: themeColor.value,
        primaryColorHover: lightenedColor,
        primaryColorPressed: lightenedColor,
        primaryColorSuppl: themeColor.value,
      },
      LoadingBar: {
        colorLoading: themeColor,
      },
    }
    return overrides
  })

  return { themeColor, themeOverrides }
})