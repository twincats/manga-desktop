<template>
  <div id="homeid" class="px-3 pt-3">
    <a-input-search
      style="--primary-6: 255, 117, 24"
      class="w-full mb-2"
      placeholder="Search Manga"
      :input-attrs="{ class: 'text-center' }"
      search-button
    />
    <div class="grid grid-cols-5 gap-2 lg:grid-cols-10 h-572px lg:h-856px">
      <div v-for="(manga, i) in mangaHome?.manga" :key="i" class="box">
        <div class="cover" @click="$router.push(`chapter/${manga.id}`)">
          <div class="h-203px">
            <img
              :alt="manga.title"
              :src="`file/${MangaTitleURL(manga.title)}/cover.webp`"
            />
          </div>
        </div>
        <div class="title select-none">
          <span>{{ manga.title }}</span>
        </div>
        <div class="chapter">
          <button @click.stop="$router.push(`/page/${manga.chapter_id}`)">
            Chapter {{ manga.chapter }}
          </button>
        </div>
      </div>
    </div>
    <div class="mt-2">
      <a-pagination
        style="
          --primary-6: 255, 117, 24;
          --color-primary-light-1: rgba(255, 117, 24, 0.2);
        "
        class="justify-center"
        v-model:current="nav.page"
        v-model:page-size="nav.limit"
        :auto-adjust="false"
        :total="mangaHome?.pagination?.total ? mangaHome?.pagination?.total : 0"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { MangaTitleURL, GetBreakPoints, SEQUENCE3 } from '@/composable/helper'
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'

/* INTERFACE */
interface Nav {
  page: number
  limit: number
}

/* INITIAL REACTIVE VARIABLE */
const mangaHome = ref<manga.MangaHome | null>(null)
const nav = reactive<Nav>({ page: 1, limit: 10 })
const { breakpoints } = GetBreakPoints()
const lg = breakpoints.greater('lg')

/* INITIAL PRELOAD FUNCTION */
//load manga
const loadManga = (n: Nav) => {
  // console.log(nav, 'berore fetching')
  GetMangaHome(n.page, n.limit).then(res => {
    mangaHome.value = res
  })
}

// ON_CREATED function loading Manga Data
// to make load nav accurate when sitch pages
if (lg.value) {
  nav.limit = 30
} else {
  nav.limit = 10
}
loadManga(nav)

// Watching View and Pagination page change
watch([lg, () => nav.page], ([l, p], [_, op]) => {
  if (l) {
    nav.limit = 30
    if (op == p) {
      nav.page = SEQUENCE3(p)
    }
    loadManga(nav)
  } else {
    nav.limit = 10
    if (op == p) {
      nav.page = (p - 1) * 3 + 1
    }
    loadManga(nav)
  }
})
</script>

<style lang="less" scoped>
// #homeid {
//   font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
// }

@mainColor: --color-bg-3;

.box {
  text-align: center;
  border: solid 2px var(@mainColor);
  border-radius: 0.5rem;
  margin-bottom: 0.25rem;
  .cover {
    padding: 0.25rem;
    div {
      overflow: hidden;
      position: relative;
      border-radius: 0.3rem 0.3rem 0 0;
      img {
        height: 120%;
        position: absolute;
        margin: auto;
        left: -999px;
        right: -999px;
      }
    }
    &:hover {
      cursor: pointer;
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
  }
}
</style>
