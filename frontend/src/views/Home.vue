<template>
  <div id="homeid" class="px-3 pt-3">
    <a-input-search
      v-model="searchManga"
      style="--primary-6: 255, 117, 24"
      @change="loadManga(searchManga)"
      allow-clear
      class="w-full mb-2"
      placeholder="Search Manga"
      :input-attrs="{ class: 'text-center' }"
      search-button
    />
    <div class="grid grid-cols-5 gap-2 lg:grid-cols-10">
      <div
        v-for="(manga, i) in mangaHome?.manga"
        :key="i"
        class="box"
        @contextmenu.prevent="openContextMenu($event, manga)"
      >
        <div class="cover" @click="$router.push(`chapter/${manga.id}`)">
          <div class="h-203px lg:h-200px">
            <img
              :alt="manga.title"
              :src="`file/${MangaTitleURL(manga.title)}/cover.webp`"
            />
          </div>
        </div>
        <div
          class="title select-none"
          @click="$router.push(`chapter/${manga.id}`)"
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
    <div class="mt-2">
      <a-pagination
        class="justify-center"
        v-model:current="nav.page"
        v-model:page-size="nav.limit"
        :hide-on-single-page="true"
        :auto-adjust="false"
        :total="mangaHome?.pagination?.total ? mangaHome?.pagination?.total : 0"
      />
    </div>
    <teleport to="#main">
      <context-menu ref="refMenu" v-slot="{ item }: { item: manga.Manga }">
        <li>Add Alternatif</li>
        <li @click="$router.push('/convert/' + item.id)">Convert Manga</li>
        <div class="divider"></div>
        <li class="text-red-500 font-serif font-bold">DELETE</li>
      </context-menu>
    </teleport>
  </div>
</template>

<script setup lang="ts">
import {
  MangaTitleURL,
  GetBreakPoints,
  SEQUENCE3,
  UseContextMenu,
} from '@/composable/helper'
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'

/* INTERFACE */
interface Nav {
  page: number
  limit: number
}

/* INITIAL REACTIVE VARIABLE */
const { refMenu, openContextMenu } = UseContextMenu()
const mangaHome = ref<manga.MangaHome | null>(null)
const nav = reactive<Nav>({ page: 1, limit: 10 })
const { breakpoints } = GetBreakPoints()
const lg = breakpoints.greater('lg')
const searchManga = ref('')

/* INITIAL PRELOAD FUNCTION */
//load manga
const loadManga = (sarch: string | null = null) => {
  // console.log(nav, 'berore fetching')
  GetMangaHome(sarch, nav.page, nav.limit).then(res => {
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
loadManga()

// Watching View and Pagination page change
watch([lg, () => nav.page], ([l, p], [_, op]) => {
  if (l) {
    nav.limit = 30
    if (op == p) {
      nav.page = SEQUENCE3(p)
    }
    loadManga(searchManga.value)
  } else {
    nav.limit = 10
    if (op == p) {
      nav.page = (p - 1) * 3 + 1
    }
    loadManga(searchManga.value)
  }
})
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
