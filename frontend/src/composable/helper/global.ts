import { useBreakpoints, breakpointsTailwind } from '@vueuse/core'

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
