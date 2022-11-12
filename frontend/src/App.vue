<template>
  <div id="main">
    {{ response }}
    <router-view></router-view>
  </div>
</template>

<script lang="ts" setup>
import { EventsOnce, EventsEmit } from '@wails/runtime/runtime'

const router = useRouter()
const response = ref()
EventsEmit('args')
EventsOnce('args', (res: string[]) => {
  response.value = res
  console.log(res)
  if (res.includes('convert')) {
    router.push('/test')
  }
})
</script>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
#main {
  background-color: var(--color-bg-1);
  height: 100vh;
}
</style>
