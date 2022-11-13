import { useColorMode, usePreferredDark, useCycleList } from '@vueuse/core'
import { computed, reactive } from 'vue'

const darkThemeVar = reactive({
  bg: [24, 24, 24],
  text: [255, 255, 255],
  scheme: 'dark',
})
const lightThemeVar = reactive({
  bg: [250, 250, 250],
  text: [5, 5, 5],
  scheme: 'light',
})

export const useTheme = () => {
  const isdark = usePreferredDark()
  const mode = useColorMode({
    selector: 'body',
    attribute: 'arco-theme',
  })

  //toggleTheme only work where theme is apply
  const { next: toggleTheme } = useCycleList(['dark', 'light'], {
    initialValue: mode,
  })

  // model auto change to dark or light
  if (mode.value == 'auto') {
    mode.value = isdark ? 'dark' : 'light'
  }

  const theme = computed(() => {
    if (mode.value == 'dark') {
      return darkThemeVar
    } else {
      return lightThemeVar
    }
  })

  return { mode, isdark, theme, toggleTheme }
}
