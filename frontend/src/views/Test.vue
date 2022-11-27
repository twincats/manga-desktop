<template>
  <div class="h-100vh" @contextmenu.prevent="openContextMenu">
    <div>TESTING PAGE</div>
    <div>
      <a-space>
        <a-button @click="showManga" type="primary">Fetch</a-button>
        <a-button @click="$router.push('/')" type="secondary" status="warning"
          >Home</a-button
        >
        <a-button type="secondary" status="warning">context</a-button>
      </a-space>
    </div>
    <div>
      {{ manga }}
    </div>
    <div>
      <ContextMenuApp ref="refMenu">
        <li @click="alerts('hello world')">alert</li>
        <li>testing</li>
        <li>Kneal</li>
        <div class="divider"></div>
        <li>Back</li>
      </ContextMenuApp>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'
import ContextMenuApp from '@/components/app/ContextMenu.vue'
import { UseContextMenu } from '@/composable/helper'

const manga = ref<manga.MangaHome>()
const showManga = () => {
  GetMangaHome(null, 1, 5).then(res => {
    manga.value = res
  })
}
const { refMenu, openContextMenu, closeContextMenu } = UseContextMenu()

const alerts = (m: string) => {
  alert(m)
  closeContextMenu()
}
</script>

<style scoped></style>
