<template>
  <div>
    <menu-bar
      @clickMenu="collapse = !collapse"
      @clickToggleTheme="toggleTheme()"
    />
    <side-bar :collapse="collapse" />
    <div id="main">
      <router-view></router-view>
    </div>
    <footer-bar />
  </div>
</template>

<script lang="ts" setup>
import { EmitListenOnce, useTheme } from '@/composable/helper'

const router = useRouter()
const collapse = ref(true)
const { theme, toggleTheme } = useTheme()
EmitListenOnce('args', (res: string[]) => {
  if (res.includes('convert')) {
    router.push('/convert')
  }
})
</script>

<style>
#main {
  height: calc(100vh - 56px);
  overflow: auto;
  background-color: rgb(v-bind('theme.bg'));
  color-scheme: v-bind('theme.scheme');
}
</style>
