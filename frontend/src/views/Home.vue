<template>
  <div class="xl:(grid grid-cols-10) px-3 pt-3">
    <div id="homeid" class="col-span-8">
      <a-input
        id="homeinput"
        class="w-full mb-2"
        placeholder="Search Manga"
        v-model="searchManga"
        @change="loadManga(searchManga)"
        :input-attrs="{ class: 'text-center', id: 'searchManga' }"
        allow-clear
      >
        <template #append>
          <a-space size="mini">
            <a-button type="primary" @click="loadManga(searchManga)"
              ><icon-search
            /></a-button>
            <a-button><icon-eraser /></a-button>
          </a-space>
        </template>
      </a-input>

      <div :style="{ minHeight: minH }">
        <div class="grid grid-cols-5 gap-2 lg:grid-cols-8">
          <div
            v-for="(manga, i) in mangaHome?.manga"
            :key="i"
            class="box relative"
            :class="today(manga.download_time)"
            @contextmenu.prevent="openContextMenu($event, manga)"
          >
            <div class="cover" @click="$router.push(`chapter/${manga.id}`)">
              <div class="h-203px lg:h-200px">
                <img
                  :alt="manga.title"
                  :src="`file/${MangaTitleURL(manga.title)}/cover.webp`"
                  :onerror="errorLoadImage"
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
            <div
              v-if="
                DateApp.NewDate().Format('DD-MM-YYYY') ==
                DateApp.NewDate(manga.download_time.toString()).Format(
                  'DD-MM-YYYY'
                )
              "
              class="absolute text-xs top-2 right-0 px-1 rounded-l border-b border-dark-100 drop-shadow bg-blue-500"
            >
              New
            </div>
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
          :total="
            mangaHome?.pagination?.total ? mangaHome?.pagination?.total : 0
          "
        />
      </div>
      <teleport to="#app">
        <context-menu ref="refMenu" v-slot="{ item }: { item: manga.Manga }">
          <li @click="$router.push(`/addalter/${item.id}`)">Add Alternatif</li>
          <li @click="$router.push('/convert/' + item.id)">Convert Manga</li>
          <div class="divider"></div>
          <li class="text-red-500 font-serif font-bold">
            <a-popconfirm
              :content="
                'Are you sure want Delete? ' + item.title.substring(0, 15)
              "
              ok-text="OK"
              cancel-text="NO"
              type="error"
            >
              <div class="w-full">DELETE</div>
            </a-popconfirm>
          </li>
        </context-menu>
      </teleport>
    </div>
    <div v-if="lg" class="ml-3 col-span-2">
      <div id="panelDate">
        <a-date-picker
          :default-value="new Date()"
          :locale="enUS.datePicker"
          :show-now-btn="false"
          @select="testLog"
          :disabled-date="
            current =>
              current ? current.getTime() > new Date().getTime() : false
          "
          hide-trigger
          style="
            width: 100%;
            --color-bg-popup: #2a2a2a;
            --color-neutral-3: rgba(0, 0, 0, 0.75);
            --link-6: var(--orange-6);
          "
        />
      </div>
      <div class="mt-3">
        <a-typography-title class="text-center" :heading="5">
          Random Unread Manga
        </a-typography-title>
        <a-list :bordered="false" :max-height="500" :scrollbar="true">
          <a-list-item v-for="(manga, i) in mangaHome?.manga" :key="i">
            <a-list-item-meta>
              <template #avatar>
                <a-image
                  :preview="false"
                  width="64"
                  :src="`file/${MangaTitleURL(manga.title)}/cover.webp`"
                />
              </template>
              <template #title>
                <a-link
                  style="color: var(--color-text-1)"
                  @click="$router.push(`chapter/${manga.id}`)"
                  >{{ manga.title.substring(0, 100)
                  }}{{ manga.title.length >= 100 ? '...' : '' }}</a-link
                >
              </template>
              <template #description>
                <a-link
                  style="color: rgba(var(--orange-6), 0.75)"
                  @click="$router.push(`/page/${manga.chapter_id}`)"
                  >Chapter {{ manga.chapter }}</a-link
                >
              </template>
            </a-list-item-meta>
          </a-list-item>
        </a-list>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { IconSearch, IconEraser } from '@arco-design/web-vue/es/icon'
import {
  MangaTitleURL,
  GetBreakPoints,
  Sequence,
  UseContextMenu,
  DateApp,
} from '@/composable/helper'
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'
import imageFail from '@/assets/images/404.webp'
import enUS from '@arco-design/web-vue/es/locale/lang/en-us'
import { useMangaState } from '@/store'

/* INITIAL REACTIVE VARIABLE */
const { refMenu, openContextMenu } = UseContextMenu()
const { mangaHome, navHome: nav, searchManga } = useMangaState()
const { breakpoints } = GetBreakPoints()
const lg = breakpoints.greater('lg')
const minH = ref('574px')

/* INITIAL PRELOAD FUNCTION */
//load manga
const loadManga = (sarch: string | null = null) => {
  // console.log(nav, 'berore fetching')
  GetMangaHome(sarch, nav.page, nav.limit).then(res => {
    mangaHome.value = res
  })
}

const maxLimit = 24
const minLimit = 10
// ON_CREATED function loading Manga Data
// to make load nav accurate when sitch pages
if (lg.value) {
  nav.limit = maxLimit
  minH.value = '856px'
} else {
  nav.limit = minLimit
  minH.value = '574px'
}
if (mangaHome.value == null) {
  loadManga()
}

// Watching View and Pagination page change
watch([lg, () => nav.page], ([l, p], [_, op]) => {
  if (l) {
    nav.limit = maxLimit
    minH.value = '856px'
    if (op == p) {
      nav.page = Sequence(3, p)
    }
    loadManga(searchManga.value)
  } else {
    nav.limit = minLimit
    minH.value = '574px'
    if (op == p) {
      nav.page = (p - 1) * 3 + 1
    }
    loadManga(searchManga.value)
  }
})

//today class
const today = (date: Date): string => {
  if (
    DateApp.NewDate().Format('DD-MM-YYYY') ==
    DateApp.NewDate(date.toString()).Format('DD-MM-YYYY')
  ) {
    return 'today'
  } else {
    return ''
  }
}

const errorLoadImage = (e: Event) => {
  const img = <HTMLImageElement>e.target
  img.src = imageFail
}

const testLog = () => {
  console.log('berUbah')
}
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

@border: rgba(255, 117, 24, 1);
.today {
  border-color: @border !important;
  .cover {
    background-color: screen(@border, #cccccc);
    border-bottom: 1px solid @border;
    border-radius: 0.5rem 0.5rem 0 0;
  }
  .title {
    background-color: screen(@border, #cccccc);
    color: black;
  }
  .chapter {
    background-color: @border;
    button {
      &:hover {
        color: white;
        background-color: rgb(59, 130, 246);
      }
    }
  }
  &:hover {
    .cover,
    .title {
      background-color: tint(@border, 60%);
    }
  }
}
</style>

<style lang="less">
#panelDate {
  .arco-panel-month,
  .arco-panel-quarter,
  .arco-panel-year {
    box-sizing: border-box;
    width: 100%;
    .arco-picker-row {
      padding: 15px 0;
    }
  }
}
#homeinput {
  .arco-input-append {
    padding: 0;
    border: none;
    .arco-space {
      &:first-child {
        button {
          border-top-left-radius: 0;
          border-bottom-left-radius: 0;
          padding: 0 10px;
        }
      }
    }
  }
}
</style>
