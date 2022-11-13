<template>
  <div>
    <div>TESTING PAGE</div>
    <div>
      <a-space>
        <a-button @click="showManga" type="primary">Fetch</a-button>
        <a-button @click="$router.push('/')" type="secondary" status="warning"
          >Home</a-button
        >
      </a-space>
    </div>
    <div>
      {{ manga }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'

const manga = ref<manga.MangaHome>()
const showManga = () => {
  GetMangaHome(1, 5).then(res => {
    manga.value = res
  })
}
</script>

<style scoped></style>
