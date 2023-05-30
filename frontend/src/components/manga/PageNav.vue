<template>
  <div>
    <div id="pnav">
      <a-pagination
        v-model:current="curPage"
        :total="chaps.length"
        :default-page-size="1"
        simple
      />
    </div>
    <div id="tool">
      <a-button
        style="
          --color-secondary: rgba(var(--gray-1), 1);
          --color-secondary-hover: rgba(var(--gray-1), 1);
          --color-secondary-active: rgba(var(--gray-2), 1);
          --border-radius-small: 6px;
          --color-text-2: rgba(var(--orange-6), 1);
        "
        size="small"
        @click="$router.push('/')"
        ><icon-home
      /></a-button>
    </div>
    <div id="pscroll">
      <a-button
        v-if="!scroll.arrivedState.top"
        shape="circle"
        type="primary"
        status="success"
        :style="{ opacity: scroll.arrivedState.bottom ? 1 : 0.3 }"
        @click="scroll.y.value = 0"
      >
        <i-mdi:format-vertical-align-top />
      </a-button>
      <a-button
        v-if="!scroll.arrivedState.bottom"
        shape="circle"
        type="primary"
        status="danger"
        :style="{ opacity: scroll.arrivedState.top ? 1 : 0.3 }"
        @click="scroll.y.value = scroolHeigh ? scroolHeigh : 0"
      >
        <i-mdi:format-vertical-align-bottom />
      </a-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { IconHome } from '@arco-design/web-vue/es/icon'

const props = defineProps<{
  chaps: number[]
}>()

const emit = defineEmits<{
  (e: 'change', cid: number): void
}>()

const route = useRoute()
const curPage = ref(1)

// fetching element with scroll
const el = document.getElementById('main')
const scroolHeigh = 1000 * 1000
const scroll = useScroll(el, { behavior: 'smooth' })

//fetching url params
const paramCid = ref<number>(Number(route.params.cid))
const initIndex = props.chaps.findIndex(item => item == Number(paramCid.value))
curPage.value = initIndex + 1

// watching curPage to trigger event change
watch(
  () => curPage.value,
  ncid => {
    if (props.chaps) {
      paramCid.value = props.chaps[ncid - 1]
      emit('change', props.chaps[ncid - 1])
    }
  }
)
</script>

<style lang="less" scoped>
#pnav {
  position: absolute;
  bottom: 35px;
  width: 165px;
  background-color: var(--color-bg-1);
  left: 50%;
  transform: translateX(-50%);
  opacity: 0.3;
  border-radius: 0.5rem;
  &:hover {
    opacity: 1;
  }
}

#pscroll {
  position: absolute;
  bottom: 35px;
  right: 20px;
  width: 32px;

  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  button {
    margin-top: 5px;
    opacity: 0.3;
    &:hover {
      opacity: 1;
    }
  }
}

#tool {
  position: absolute;
  bottom: 35px;

  left: 80%;
  opacity: 0.3;
  &:hover {
    opacity: 0.8;
  }
}
</style>
