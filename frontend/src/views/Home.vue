<template>
  <div class="xl:(grid grid-cols-12) px-3 pt-3">
    <div id="homeid" class="col-span-10">
      <a-input
        id="homeinput"
        class="w-full mb-2"
        placeholder="Search Manga"
        :value="searchManga"
        @change="e => (searchManga = e)"
        :input-attrs="{
          class: 'text-center',
          id: 'searchManga',
          spellcheck: false,
          autocomplete: false,
        }"
        allow-clear
      >
        <template #append>
          <a-space size="mini">
            <a-button type="primary"><icon-search /></a-button>
            <a-button @click="clearFilter"><icon-eraser /></a-button>
          </a-space>
        </template>
      </a-input>

      <div :style="{ minHeight: minH }">
        <div class="grid grid-cols-5 gap-2 lg:grid-cols-10" draggable="false">
          <div
            v-for="(manga, i) in mangaView"
            :key="i"
            class="box relative"
            :class="
              isEqualDate(new Date(manga.download_time), new Date())
                ? 'today'
                : ''
            "
            @contextmenu.prevent="openContextMenu($event, manga)"
          >
            <div class="cover" @click="$router.push(`/chapter/${manga.id}`)">
              <div>
                <m-image
                  height="200px"
                  width="165px"
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
          v-model:page-size="nav.per_page"
          :hide-on-single-page="true"
          :auto-adjust="false"
          :total="totalManga"
        />
      </div>
      <teleport to="#app">
        <context-menu
          ref="refMenu"
          v-slot="{ item }: { item: manga.MangaHomeApi }"
        >
          <li @click="$router.push(`/addalter/${item.id}`)">Add Alternatif</li>
          <li @click="$router.push('/convert/' + item.id)">Convert Manga</li>
          <div class="divider"></div>
          <li
            @click="showModalDeleteConfirm(item)"
            class="text-red-500 font-serif font-bold"
          >
            DELETE
          </li>
        </context-menu>
      </teleport>
    </div>
    <div v-if="lg" class="ml-3 col-span-2 hidden lg:block">
      <div id="panelDate">
        <a-date-picker
          :locale="enUS.datePicker"
          :show-now-btn="false"
          value-format="Date"
          v-model="dateFilter"
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
          <a-list-item v-for="(manga, i) in randomManga" :key="i">
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
    <a-modal
      v-model:visible="Del.modal"
      ok-text="Delete"
      :mask-closable="false"
      :body-style="{ padding: '.5rem 1.5rem' }"
      :ok-button-props="{ status: 'danger' }"
      :simple="true"
      :width="500"
      :esc-to-close="false"
      @before-ok="handleDelete"
    >
      <template #title>
        <span class="select-none"
          ><icon-exclamation-circle-fill class="text-#ef4444" /> Confirmation
          Delete</span
        >
      </template>
      <div class="select-none text-center">
        You are going to Delete : <br />
        <strong class="text-#ef4444">{{ Del.Title }} </strong><br />
        with all Chapter and Images. <br />
        Are you sure?
      </div>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { IconSearch, IconEraser } from '@arco-design/web-vue/es/icon'
import { MangaTitleURL, UseContextMenu, DateApp } from '@/composable/helper'
import { GetRandomMangaHome, DeleteMangaWithFile } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'
import enUS from '@arco-design/web-vue/es/locale/lang/en-us'
import { useMangaState } from '@/store'
import { IconExclamationCircleFill } from '@arco-design/web-vue/es/icon'

/* INITIAL REACTIVE VARIABLE */
const { refMenu, openContextMenu } = UseContextMenu()
const {
  lg,
  useNavStorage,
  getLimit,
  loadManga,
  totalManga,
  mangaHome,
  dateFilter,
  searchManga,
} = useMangaState()

const mangaLimit = getLimit(10, 30)
const minH = getLimit('574px', '856px')
const nav = useNavStorage(mangaLimit)
const randomManga = ref<manga.MangaHomeApi[]>([])

//update per_page when created
nav.per_page = mangaLimit.value

/* FUNCTION & COMPUTED DATA */
// isEqualDate check date1 and Date1
const isEqualDate = (date1: Date, date2: Date) => {
  return date1.toLocaleDateString() === date2.toLocaleDateString()
}

// mangaView computed mangaHome for view
const mangaView = computed<manga.MangaHomeApi[]>(() => {
  // check mangaHome empty
  if (!mangaHome.value) {
    return []
  }

  var mangaList = mangaHome.value.manga
  const start = (nav.page - 1) * nav.per_page

  // mangaView filtered when doing search
  if (searchManga.value !== '') {
    dateFilter.value = new Date()
    const filteredManga = mangaList.filter(manga =>
      manga.title.toLowerCase().includes(searchManga.value.toLowerCase())
    )
    const end = Math.min(nav.page * nav.per_page, filteredManga.length)
    totalManga.value = filteredManga.length
    return filteredManga.slice(start, end)
  }

  // mangaView filtered with dateFilter
  if (!isEqualDate(dateFilter.value, new Date())) {
    const filteredMangaDate = mangaHome.value.manga.filter(manga =>
      isEqualDate(new Date(manga.download_time), dateFilter.value)
    )
    mangaList = filteredMangaDate
  }

  const end = Math.min(nav.page * nav.per_page, mangaList.length)
  totalManga.value = mangaList.length
  return mangaList.slice(start, end)
})

// showRandomUnread is getRandome manga with Limit
const showRandomUnread = async () => {
  randomManga.value = await GetRandomMangaHome(10)
}

// clear filter button function
const clearFilter = () => {
  dateFilter.value = new Date()
  searchManga.value = ''
  nav.page = 1
}

const Del = reactive({
  ID: 0,
  Title: '',
  modal: false,
})
// showModalDeleteConfirm
const showModalDeleteConfirm = (m: manga.MangaHomeApi) => {
  // console.log('showmodal ', m.title)
  Del.ID = m.id
  Del.Title = m.title
  Del.modal = true
  refMenu.value?.close()
}

// handleDelete
const handleDelete = async () => {
  if (Del.ID != 0) {
    try {
      await DeleteMangaWithFile(Del.ID)
      // delete mangastate
      if (isDefined(mangaHome)) {
        const i = mangaHome.value.manga.findIndex(it => it.id == Del.ID)
        i >= 0 ? mangaHome.value.manga.splice(i, 1) : ''
        // reset Del Data
        Del.ID = 0
        Del.Title = ''
      }
      return true
    } catch (error) {
      console.log(error)
      return false
    }
  }
}

/* INITIAL PRELOAD FUNCTION */
loadManga() //load manga
showRandomUnread() //display random
</script>

<style lang="less">
@mainColor: --color-bg-3;

.box {
  text-align: center;
  border: solid 2px var(@mainColor);
  border-radius: 0.5rem;
  margin-bottom: 0;

  .cover {
    padding: 0.25rem;
    border-bottom: 1px solid transparent;
    div {
      overflow: hidden;
      position: relative;
      border-radius: 0.3rem 0.3rem 0 0;
    }
  }
  .title {
    padding: 0.5rem 0;

    span {
      line-height: 1.5;
      overflow: hidden;
      display: -webkit-box;
      line-clamp: 1;
      -webkit-box-orient: vertical;
      -webkit-line-clamp: 1;
    }
  }
  .chapter {
    padding: 0.25rem 0;
    border-radius: 0 0 0.35rem 0.35rem;
    background-color: var(@mainColor);
    user-select: none;

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
