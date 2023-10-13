<template>
  <div class="px-2">
    <div>
      <div :style="{ minHeight: minH }">
        <div class="grid grid-cols-7 gap-2 lg:grid-cols-14">
          <div class="box" v-for="(manga, i) in mview" :key="i">
            <div class="cover" @click="$router.push(`/chapter/${manga.id}`)">
              <m-image
                height="150px"
                fit="cover"
                lazy
                :src="`file/${MangaTitleURL(manga.title)}/cover.webp`"
              />
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
const minH = getLimit('574px', '856px')

const mview = computed(() => {
  if (!mangaHome.value) {
    return []
  }
  return mangaHome.value.manga.sort((a, b) => a.title.localeCompare(b.title))
})

loadManga()
</script>

<style lang="less" scoped>
@mainColor: --color-bg-3;

.box {
  text-align: center;
  border: solid 2px var(@mainColor);
  border-radius: 0.5rem;
  margin-bottom: 0.25rem;
  .cover {
    padding: 0.25rem;
    border-bottom: 1px solid transparent;
    div {
      overflow: hidden;
      position: relative;
      border-radius: 0.3rem 0.3rem 0 0;
      // img {
      //   height: 120%;
      //   position: absolute;
      //   margin: auto;
      //   left: -999px;
      //   right: -999px;
      // } // penyebab error observer
    }
  }
  .title {
    padding: 0.5rem 0;

    span {
      line-height: 1.5;
      overflow: hidden;
      display: -webkit-box;
      -webkit-box-orient: vertical;
      -webkit-line-clamp: 1;
    }
  }
  .chapter {
    padding: 0.25rem 0;
    border-radius: 0 0 0.35rem 0.35rem;
    background-color: var(@mainColor);

    button {
      background-color: unset;
      border: unset;
      &:hover {
        color: var(--app-main);
        background-color: var(--app-main-fade);
        border-radius: 0.25rem;
        cursor: pointer;
      }
    }
  }
  &:hover {
    background-color: var(--color-fill-4);
    .cover,
    .title {
      cursor: pointer;
    }
  }
}
</style>
