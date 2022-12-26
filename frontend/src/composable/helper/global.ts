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

// Sequence return n number to s sequence number + 1
// example sequence 3 => 1,1,1,2,2,2,3,3,3,4 (1..10)
export const Sequence = (s: number, n: number): number => {
  for (let i = 0; i < s; i++) {
    const modSec = i + 1 < s ? i + 1 : 0
    if (n % s == modSec) {
      return Math.floor((n - i) / s) + 1
    }
  }
  return NaN
}

export class DateApp {
  date: Date
  locale: string
  formatter: Intl.RelativeTimeFormat
  #DIVISIONS = [
    { amount: 60, name: 'seconds' },
    { amount: 60, name: 'minutes' },
    { amount: 24, name: 'hours' },
    { amount: 7, name: 'days' },
    { amount: 4.34524, name: 'weeks' },
    { amount: 12, name: 'months' },
    { amount: Number.POSITIVE_INFINITY, name: 'years' },
  ]
  constructor(s: string | number | null = null) {
    if (typeof s == 'string') {
      const parse = s ? Date.parse(s) : Date.now()
      this.date = new Date(parse)
    } else if (typeof s == 'number') {
      this.date = new Date(s)
    } else {
      this.date = new Date()
    }

    // console.log(s)

    this.locale = 'id-ID'
    this.formatter = new Intl.RelativeTimeFormat(this.locale, {
      localeMatcher: 'best fit',
      numeric: 'always',
    })
  }
  static NewDate(s: string | number | null = null) {
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
  formatTimeAgo() {
    let duration = (this.date.valueOf() - new Date().valueOf()) / 1000
    for (let i = 0; i <= this.#DIVISIONS.length; i++) {
      const division = this.#DIVISIONS[i]
      if (Math.abs(duration) < division.amount) {
        return this.formatter.format(
          Math.round(duration),
          <Intl.RelativeTimeFormatUnit>division.name
        )
      }
      duration /= division.amount
    }
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
  const SetPasteData = async (v: string) => {
    await navigator.clipboard.writeText(v)
  }
  return {
    pasteData,
    GetPasteData,
    SetPasteData,
  }
}

//check if string us URL
export const IsURL = (text: string) => {
  var URLObject: URL | null = null
  try {
    const PasteURL = new URL(text)
    URLObject = PasteURL
  } catch (e) {
    URLObject = null
  }
  return URLObject
}

export const UseParseDom = (data: string): Document => {
  const parser = new DOMParser()
  const val = parser.parseFromString(data, 'text/html')
  return val
}
