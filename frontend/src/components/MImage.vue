<template>
  <img
    ref="images"
    :alt="alt"
    :style="{ ...sizeStyle, ...fitStyle }"
    :src="srcData"
    :onerror="onErrorLoad"
  />
</template>

<script setup lang="ts">
import type { CSSProperties } from 'vue'
import imageFail from '@/assets/images/404.webp'

interface Props {
  src: string
  width?: string | number
  height?: string | number
  fit?: 'contain' | 'cover' | 'fill' | 'none' | 'scale-down'
  preview?: boolean
  lazy?: boolean
  alt?: string
  thresold?: number
}

const props = withDefaults(defineProps<Props>(), {
  preview: false,
  lazy: false,
  thresold: 0,
})

const srcData = ref(props.src)
if (props.lazy) {
  srcData.value = ''
}

const updateImage = () => {
  srcData.value = props.src
}

const fitStyle = computed<CSSProperties>(() => {
  if (props.fit) {
    return { objectFit: props.fit }
  }
  return {}
})

const opt = Object.prototype.toString

function isNumber(obj: any): obj is number {
  return opt.call(obj) === '[object Number]' && obj === obj // eslint-disable-line
}

function isUndefined(obj: any): obj is undefined {
  return obj === undefined
}

const normalizeImageSizeProp = (size?: string | number) => {
  if (isUndefined(size)) return undefined
  if (!isNumber(size) && /^\d+(%)$/.test(size)) return size

  const num = parseInt(size as string, 10)

  return isNumber(num) ? `${num}px` : undefined
}

const sizeStyle = computed(() => ({
  width: normalizeImageSizeProp(props.width),
  height: normalizeImageSizeProp(props.height),
}))

const images = ref<HTMLImageElement | null>(null)

const onErrorLoad = (e: Event) => {
  srcData.value = imageFail
}

onMounted(() => {
  const observer = new IntersectionObserver(
    (entries: IntersectionObserverEntry[]) => {
      entries.forEach((entry: IntersectionObserverEntry) => {
        if (entry.isIntersecting) {
          updateImage()
          observer.unobserve(entry.target)
        }
      })
    },
    { threshold: props.thresold } // Adjust the threshold as needed for testing
  )

  if (images.value) {
    observer.observe(images.value)
  }
})
</script>
