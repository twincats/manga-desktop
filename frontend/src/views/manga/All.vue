<template>
  <div class="p-2">
    <div>
      <div>
        <div class="grid grid-cols-7 gap-2 lg:grid-cols-14">
          <div class="box" v-for="(manga, i) in mview" :key="i">
            <div class="cover" @click="$router.push(`/chapter/${manga.id}`)">
              <div>
                <m-image
                  height="148px"
                  width="120px"
                  fit="cover"
                  :src="`file/${MangaTitleURL(manga.title)}/cover.webp`"
                />
              </div>
            </div>
            <div
              class="title select-none"
              @click="$router.push(`/chapter/${manga.id}`)"
            >
              <span>{{ manga.title }}</span>
            </div>
            <div class="chapter">
              <button @click.stop="$router.push(`/page/${manga.chapter_id}`)">
                Chapter {{ manga.chapter }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { manga } from '@wails/go/models'
import { MangaTitleURL } from '@/composable/helper'
import { useMangaState } from '@/store'

const { loadManga, mangaHome, getLimit } = useMangaState()
// const imgHigh = getLimit('133px', '148px')

const mview = computed(() => {
  if (!mangaHome.value) {
    return []
  }
  const mangaview = ref([...mangaHome.value.manga])
  return mangaview.value.sort((a, b) => a.title.localeCompare(b.title))
})

loadManga()
</script>

<style lang="less" scoped></style>
