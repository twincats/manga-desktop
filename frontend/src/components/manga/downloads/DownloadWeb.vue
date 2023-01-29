<template>
  <div>
    <div class="h-[calc(100vh-7.5rem)] px-2">
      <a-input-search
        allow-clear
        class="w-full mb-2"
        placeholder="URL"
        v-model="url"
        search-button
        @search="fetchBrowserData"
      />
      <div
        ref="web"
        class="h-[calc(100vh-11.2rem)] overflow-auto p-2"
        style="background-color: var(--color-bg-2)"
        v-html="browserData?.body"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { WebBrowser } from '@wails/go/tool/Web'
import type { tool } from '@wails/go/models'
import { UseParseDom } from '@/composable/helper'

// const openBrowser = (uri: string) => {
//   console.log(url)
//   tabsActive.value = 2
//   url.value = 'https://' + uri
//   fetchBrowserData()
// }

//TAB 2
const web = ref<HTMLElement | null>(null)
const url = ref('')
const browserData = ref<tool.Web | null>(null)
const fetchBrowserData = () => {
  if (url.value) {
    WebBrowser(url.value).then(res => {
      const htmlDoc = UseParseDom(res.body)
      const v = htmlDoc.getElementsByTagName('h1')
      console.log(v.item(0)?.innerText)
    })
  }
}
useEventListener(web, 'click', e => {
  e.preventDefault()
  if (e.target instanceof HTMLAnchorElement) {
    console.log(e.target.href)
  } else {
    const path1 = e.composedPath()[1]
    if (path1 instanceof HTMLAnchorElement) {
      console.log(path1.href)
    }
  }
})
</script>

<style scoped></style>
