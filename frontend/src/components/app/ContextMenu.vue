<template>
  <div
    ref="cmenu"
    class="context-menu"
    v-if="data.show"
    :style="{
      left: data.x + 'px',
      top: data.y - 10 + 'px',
      width: prop.width + 'px',
    }"
  >
    <ul class="">
      <slot :item="extData"></slot>
    </ul>
  </div>
</template>

<script setup lang="ts">
// setup props
interface Props {
  width?: number
}

// defineProps with default value need reactivityTransform: true
const prop = withDefaults(defineProps<Props>(), {
  width: 250,
})

// setup emit
// const emit = defineEmits<{
//   (e: 'close'): void
// }>()
const { width: winWidth, height: winHeight } = useWindowSize()
const cmenu = ref<HTMLElement | null>(null)
const extData = ref<any>(null)
const data = reactive({
  show: false,
  x: 0,
  y: 0,
})
const el = document.getElementById('main')
const scroll = useScroll(el)
const close = () => (data.show = false)
onClickOutside(cmenu, () => close())
watch(
  () => scroll.isScrolling.value,
  scrollStat => {
    if (scrollStat) {
      close()
    }
  }
)

const open = (event: MouseEvent, dataExt: any = null) => {
  const limitX = winWidth.value - event.x
  const limitY = winHeight.value - event.y
  data.show = true
  data.x = event.x
  data.y = event.y
  extData.value = dataExt
  if (prop.width > limitX) {
    data.x = winWidth.value - prop.width
  }
  if (cmenu.value) {
    const height = cmenu.value.offsetHeight
    if (height > limitY) {
      data.y = winHeight.value - height
    }
  }
}
// expose tobe available in $refs instance
defineExpose({
  open,
  close,
})
</script>

<style lang="less">
.context-menu {
  background-color: var(--app-context);
  position: fixed;
  border-radius: 0.5rem;
  box-shadow: 1px 1px 1px 0px rgba(0, 0, 0, 0.5);
  z-index: 50;

  ul {
    list-style: none;
    margin: 0;
    padding: 0;

    li {
      user-select: none;
      margin: 0.25rem;
      padding: 0.25rem 0.5rem;
      display: flex;
      border-radius: 0.375rem;
      font-size: 0.875rem;
      line-height: 1.25rem;
      font-weight: lighter;
      &:hover {
        background-color: var(--app-context-light);
      }
    }
    .divider {
      height: 1px;
      background-color: var(--app-context-light);
    }
  }
}
</style>
