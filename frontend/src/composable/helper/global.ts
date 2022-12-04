import { useBreakpoints, breakpointsTailwind } from '@vueuse/core'
import ContextMenuApp from '@/components/app/ContextMenu.vue'

export const MangaTitleURL = (url: string) => {
  const fixed = url.replaceAll(/[^A-Za-z0-9_\-[\]()' ~.,!@&]|\.+$/gm, '')
  return fixed.trim()
}

export const GetBreakPoints = () => {
  const breakpoints = useBreakpoints(breakpointsTailwind)

  const bp = {
    sm: breakpoints.smaller('sm'),
    sme: breakpoints.smallerOrEqual('sm'),
    md: breakpoints.between('sm', 'md'),
    lg: breakpoints.between('md', 'lg'),
    xl: breakpoints.between('lg', 'xl'),
    xxl: breakpoints.between('xl', '2xl'),
    xxxl: breakpoints['2xl'],
    breakpoints: breakpoints,
  }
  return bp
}

export const SEQUENCE3 = (n: number): number => {
  if (n % 3 == 1) {
    return Math.floor(n / 3) + 1
  }
  if (n % 3 == 2) {
    return Math.floor((n - 1) / 3) + 1
  }
  if (n % 3 == 0) {
    return Math.floor((n - 2) / 3) + 1
  }
  return NaN
}

export class DateApp {
  date: Date
  locale: string
  constructor(s: string | null = null) {
    const parse = s ? Date.parse(s) : Date.now()
    // console.log(s)
    this.date = new Date(parse)
    this.locale = 'id-ID'
  }
  static NewDate(s: string | null = null) {
    return new DateApp(s)
  }
  Format(f: string): string {
    const format = {
      D: this.date.toLocaleString(this.locale, { day: 'numeric' }),
      DD: this.date.toLocaleString(this.locale, { day: '2-digit' }),
      MM: this.date.toLocaleString(this.locale, { month: '2-digit' }),
      MMM: this.date.toLocaleString(this.locale, { month: 'short' }),
      MMMM: this.date.toLocaleString(this.locale, { month: 'long' }),
      YY: this.date.toLocaleString(this.locale, { year: '2-digit' }),
      YYYY: this.date.toLocaleString(this.locale, { year: 'numeric' }),
      HH: this.date.toLocaleString(this.locale, { hour: '2-digit' }),
      mm: (
        '0' + this.date.toLocaleString(this.locale, { minute: '2-digit' })
      ).slice(-2),
      ss: (
        '0' + this.date.toLocaleString(this.locale, { second: '2-digit' })
      ).slice(-2),
      ddd: this.date.toLocaleString(this.locale, { weekday: 'short' }),
      dddd: this.date.toLocaleString(this.locale, { weekday: 'long' }),
    }
    //loop trough format string
    for (const [key, value] of Object.entries(format)) {
      if (f.split(/\W/).includes(key)) {
        f = f.replace(key, value.toString())
      }
    }
    return f
  }
}

// context menu instance
export const UseContextMenu = () => {
  const refMenu = ref<InstanceType<typeof ContextMenuApp> | null>(null)
  const openContextMenu = (ev: MouseEvent, data: any = null) => {
    refMenu.value?.open(ev, data)
  }
  const closeContextMenu = () => {
    refMenu.value?.close()
  }
  return {
    refMenu,
    openContextMenu,
    closeContextMenu,
  }
}

export const useClipboardData = () => {
  const pasteData = ref('')
  const GetPasteData = async () => {
    pasteData.value = await navigator.clipboard.readText()
    return pasteData.value
  }
  return {
    pasteData,
    GetPasteData,
  }
}
