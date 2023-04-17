<template>
  <div
    id="download"
    class="m-1"
    style="--primary-6: 255, 117, 24; --primary-5: 204, 92, 18"
  >
    <a-tabs v-model:active-key="tabsActive">
      <a-tab-pane :key="1" title="Download">
        <div class="px-2 my-2">
          <a-input
            v-model="urldata"
            :style="{ width: '100%' }"
            :input-attrs="{ class: 'text-center' }"
            placeholder="URL"
            allow-clear
          >
            <template #append>
              <a-space size="mini">
                <a-button type="secondary" class="px-2" @click="clearDownload">
                  <span class="icons">
                    <i-mdi:eraser />
                  </span>
                </a-button>
                <a-button
                  @click="goGetchDownload"
                  type="secondary"
                  class="px-2"
                >
                  <span class="icons" style="color: rgb(255, 117, 24)"
                    ><i-mdi:cloud-search-outline
                  /></span>
                </a-button>
              </a-space>
            </template>
          </a-input>
          <div class="mb-2 p-2" style="background-color: var(--color-bg-2)">
            <div id="servers" v-if="servers">
              <a-radio-group
                v-model="selectedServer"
                size="mini"
                class=""
                type="button"
              >
                <div class="grid grid-cols-8 w-full">
                  <a-radio
                    v-for="(item, i) in servers"
                    :key="i"
                    :value="item.id"
                  >
                    {{ item.name }}</a-radio
                  >
                </div>
              </a-radio-group>
            </div>
          </div>
        </div>
        <div
          v-if="chapterList"
          class="h-[calc(100vh-16rem)] px-2 overflow-auto"
        >
          <div
            class="min-h-220px mb-2"
            style="background-color: var(--color-bg-2)"
          >
            <div class="text-lg p-2 text-center">
              <div v-if="chapterList">{{ chapterList.manga }}</div>
            </div>
            <div class="flex">
              <div class="w-150px h-full px-2">
                <img
                  v-if="chapterList"
                  class="w-full"
                  :src="chapterList.cover_url"
                  alt="cover"
                />
                <div></div>
              </div>
              <div class="w-full flex flex-col" v-if="chapterList">
                <div class="h-full min-h-120px">
                  Server : {{ chapterList.server_name }} <br />
                  Total Chapter : {{ chapterList.chapter.length }} <br />
                  Download : <br />
                  Selected :
                  <span v-if="selected_chapter_url.length > 0">Chapter </span>
                  <span v-for="(item, i) in selected_chapter_url" :key="i">
                    {{ item.chapter }}
                    <i-twemoji-flag-for-flag-indonesia
                      v-if="item.language == 'Indonesia'"
                    />
                    <i-twemoji-flag-for-flag-united-kingdom v-else />
                    <span v-if="i + 1 < selected_chapter_url.length">, </span>
                  </span>
                  <br />
                </div>
                <div class="p-2 pl-0" v-if="selectedServer == 1">
                  <a-form-item
                    field="size"
                    style="margin-bottom: 0"
                    label="Mangadex Server"
                  >
                    <a-radio-group v-model="mdexPageServer" size="small">
                      <a-radio :value="false">Original</a-radio>
                      <a-radio :value="true">Data Saver</a-radio>
                    </a-radio-group>
                  </a-form-item>
                </div>
              </div>
              <div class="p-2 self-end" v-if="chapterList">
                <a-button
                  :disabled="chapterList.manga_id < 1"
                  status="warning"
                  @click="$router.push(`chapter/${chapterList?.manga_id}`)"
                  >Read</a-button
                >
                <a-button @click="progress_dl.modal_dl = !progress_dl.modal_dl"
                  >Test</a-button
                >
              </div>
            </div>
          </div>
          <div v-if="chapterList">
            <a-table
              :data="tableDownload"
              :pagination="{
                pageSize: tablePageSize,
                total: chapterList.total,
                size: 'mini',
              }"
              :loading="tableLoading"
              @page-change="tablePageChange"
              @change="(data, extra) => tableChange(<types.ChapterList[]>data, <TableChangeExtra>extra)"
              class="select-none"
              size="mini"
              :scroll="{ x: 1200 }"
              @row-click="(c)=>clickRowTable(<types.ChapterList>c)"
            >
              <template #columns>
                <a-table-column
                  fixed="left"
                  :width="100"
                  align="center"
                  title="Chapter"
                  data-index="chapter"
                />
                <a-table-column :width="50" title="Vol" data-index="volume" />
                <a-table-column title="Judul Chapter" data-index="title" />

                <a-table-column title="Group" data-index="group_name" />

                <a-table-column :width="135" align="center" title="Release">
                  <template #cell="{ record }: { record: types.ChapterList }">
                    {{
                      DateApp.NewDate(record.timestamp * 1000).formatTimeAgo()
                    }}
                  </template>
                </a-table-column>
                <a-table-column
                  fixed="right"
                  :width="80"
                  align="center"
                  title="Bahasa"
                  data-index="language"
                >
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <i-twemoji-flag-for-flag-indonesia
                      v-if="record.language == 'Indonesia'"
                    />
                    <i-twemoji-flag-for-flag-united-kingdom v-else />
                  </template>
                </a-table-column>
                <a-table-column
                  fixed="right"
                  :width="70"
                  align="center"
                  title="Status"
                >
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <span
                      class="bg-green-600 px-2 rounded-lg"
                      v-if="record.status"
                      >OK</span
                    >
                    <span v-else class="bg-red-600 px-2 rounded-lg">No</span>
                  </template>
                </a-table-column>
                <a-table-column
                  :fixed="'right'"
                  :width="50"
                  title="Check"
                  align="center"
                >
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <span
                      v-if="selected_chapter_url.indexOf(record) != -1"
                      class="text-green-600"
                      ><icon-check
                    /></span>
                    <span v-else class="text-red-600"><icon-minus /> </span>
                  </template>
                </a-table-column>
                <a-table-column
                  :fixed="'right'"
                  :width="110"
                  class="min-h-20px"
                  title="Download"
                  align="center"
                >
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <a-button
                      size="mini"
                      @click.stop="testDownload(record)"
                      :disabled="selected_chapter_url.length > 1"
                      >Download</a-button
                    >
                  </template>
                </a-table-column>
              </template>
              <template #th="{ column }">
                <custom-download-th
                  :active="selected_chapter_url.length > 1"
                  :column="column"
                  @click-download="testDownload()"
                />
              </template>
              <template #empty>
                <div
                  class="arco-empty flex justify-center items-center md:min-h-150px lg:min-h-420px"
                >
                  <div>
                    <div class="arco-empty-image">
                      <i-mdi-emoticon-sad-outline />
                    </div>
                    <div class="arco-empty-description">
                      <div class="my-2"></div>
                      Tidak ada data
                    </div>
                  </div>
                </div>
              </template>
            </a-table>
          </div>
        </div>
      </a-tab-pane>
      <a-tab-pane :key="2" title="Web">
        <download-web />
      </a-tab-pane>
    </a-tabs>
    <!-- modal -->
    <a-modal
      v-model:visible="progress_dl.modal_dl"
      :footer="false"
      :mask-closable="false"
      :body-style="{ padding: '.5rem 1.5rem' }"
    >
      <template #title> Download Progress </template>
      <div class="select-none">
        <div class="text-center mb-3">Chapter {{ progress_dl.chapter }}</div>
        <div id="mprogress" class="mt-1 text-center">
          <a-progress
            status="success"
            :show-text="false"
            :stroke-width="15"
            :percent="progress_dl.chap"
          />
          {{ progress_dl.index_chap }} to {{ progress_dl.total_chap }}
        </div>
        <div id="mprogress" class="mt-1 text-center">
          <a-progress
            status="normal"
            :show-text="false"
            :stroke-width="15"
            :percent="progress_dl.page"
          />
          {{ progress_dl.index_page }} to {{ progress_dl.total_page }}
        </div>
        <div class="text-center mt-3">
          Please Wait <a-spin style="--primary-6: 255, 117, 24" :size="14" />
        </div>
      </div>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { IconCheck, IconMinus } from '@arco-design/web-vue/es/icon'
import { useClipboardData, IsURL, DateApp, Sequence } from '@/composable/helper'
import { EventsOn } from '@wails/runtime/runtime'
import {
  GetChapter,
  GetChapterMdexPagination,
  GetPage,
} from '@wails/go/download/Download'
import { Message, TableChangeExtra } from '@arco-design/web-vue'
import '@arco-design/web-vue/es/message/style/index'
import { types } from '@wails/go/models'
import { UseTable, UseServer } from '@/composable/downloads/download'
import { useDownloadState } from '@/store/global'
import type { EventChap, EventPage } from '@/type/download'

//// STARTING CODE BOOTUP ////
const { tableDownload, tableLoading, tablePageSize } = UseTable()
const { servers, getSelectedServer } = UseServer()

// GLOBAL VALUE
const tabsActive = ref(1)
const { chapterList, urldata, selectedServer } = useDownloadState()
const progress_dl = ref({
  modal_dl: false,
  chapter: 0,
  index_chap: 0,
  index_page: 0,
  total_chap: 0,
  total_page: 0,
  chap: 0,
  page: 0,
})

// FOR NEXT PROGRESS
/*
  1. status check
  2. status download
  3. save db
  4. retry
*/

// GET CHAPTER DATA
const { GetPasteData } = useClipboardData()

const goGetchDownload = async () => {
  // auto fetch paste URL(only)
  const paste = await GetPasteData()
  const pasteURL = IsURL(paste)
  if (pasteURL != null) {
    urldata.value = pasteURL.href
    //here testing download chapter
  }

  //testing download chapter
  const server = getSelectedServer(selectedServer.value)
  if (server && urldata.value != '') {
    GetChapter({
      url: urldata.value,
      server_name: server.name,
    })
      .then(res => {
        //warning torefs expect reactive
        chapterList.value = res
        res.chapter.forEach(item => {
          tableDownload.push(item)
        })
        // console.log(res)
      })
      .catch(e => {
        console.log(e)
        Message.error(e)
      })
  }
}

// if chapter is avalable insert table
chapterList.value?.chapter.forEach(item => {
  tableDownload.push(item)
})

// const testUpdatedTable = () => {
//   const i = tableDownload.findIndex(
//     i => i.id == 'ff944168-9143-44f2-979c-a3df74d14075'
//   )
//   tableDownload[i].status = true
// }

///TESTING FOR TEMPORARY FILL DATA DOWNLOAD
// DOWNLOAD SINGLE OR MULTIPLE CHAPTER
const mdexPageServer = ref(false) // modelValue mdex datasaver server
const testDownload = (c: types.ChapterList | null = null) => {
  const server = getSelectedServer(selectedServer.value)
  if (server && urldata.value != '' && chapterList.value != null) {
    var oPage = new types.OptionPage()
    oPage.server_name = server.name
    oPage.datasaver = mdexPageServer.value
    oPage.manga_title = chapterList.value.manga
    if (c == null) {
      oPage.chapters = selected_chapter_url.value
    } else {
      oPage.chapters = [c]
    }

    // call
    progress_dl.value.modal_dl = true
    GetPage(oPage)
      .then(res => {
        console.log(res)
      })
      .catch(e => {
        console.log(e)
      })

    EventsOn('dl_eventchap', (p: EventChap) => {
      progress_dl.value.chapter = p.chapter
      progress_dl.value.index_chap = p.index_chap
      progress_dl.value.total_chap = p.total_chap
      progress_dl.value.total_page = p.total_page
      progress_dl.value.index_page = 0
      progress_dl.value.chap = p.index_chap / p.total_chap
    })
    EventsOn('dl_eventpage', (p: EventPage) => {
      progress_dl.value.index_page += 1
      progress_dl.value.page =
        progress_dl.value.index_page / progress_dl.value.total_page
    })
  }
}

// SET SELECTED CHAPTER TO DOWNLOAD
const selected_chapter_url = ref<types.ChapterList[]>([])
const clickRowTable = (c: types.ChapterList) => {
  const i = selected_chapter_url.value.findIndex(i => i.chapter == c.chapter)
  if (i != -1) {
    selected_chapter_url.value.splice(i, 1)
  } else {
    selected_chapter_url.value.push(c)
  }
  selected_chapter_url.value.sort(
    (a, b) => Number(a.chapter) - Number(b.chapter)
  )
}

///TESTING
const tablePageChange = (p: number) => {
  console.log(p)
}

const tableChange = (chap: types.ChapterList[], extra: TableChangeExtra) => {
  //basic limit = 30 ==> Offset calculation, offset = page * limit - limit
  const limit = 30
  const pageSize = extra.pageSize ? extra.pageSize : limit
  const page = extra.page ? extra.page : 1

  if (chap.length == 0) {
    tableLoading.value = true
    const seq = limit / pageSize
    const pSeq = Sequence(seq, page)
    const offset = pSeq * limit - limit
    // const startIndex = pSeq * limit
    // console.log(pSeq, offset, startIndex)
    GetChapterMdexPagination(urldata.value, limit, offset)
      .then(res => {
        console.log(res)
        if (res) {
          tableDownload.splice(offset, 0, ...res)
        }
        tableLoading.value = false
      })
      .catch(e => {
        console.log(e)
        tableLoading.value = false
      })
    // console.log(extra)
  }
}

//clearDownload
const clearDownload = () => {
  urldata.value = ''
  selectedServer.value = 1
  chapterList.value = null
  tableDownload.length = 0
}

//watch URL DATA => auto select servers
watchDebounced(
  [urldata, selectedServer],
  ([newURL]) => {
    const isURL = IsURL(newURL)
    if (isURL) {
      const res = servers.value?.find(item => item.url == isURL.host)
      if (res?.id) {
        selectedServer.value = res.id
      }
    }
  },
  { debounce: 50 }
)
</script>

<style>
#mprogress .arco-progress-line {
  border-radius: 0.25rem !important;
}
#mprogress .arco-progress-line-bar {
  border-radius: 0.25rem !important;
}
#mtable .arco-table-td {
  font-size: 0.75rem !important;
}
</style>
