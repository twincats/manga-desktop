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
            id="urlinput"
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
                <div class="grid md:grid-cols-8 lg:grid-cols-16 w-full">
                  <a-radio
                    v-for="(item, i) in servers"
                    :key="i"
                    :value="item.id + 1"
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
                <div class="h-full min-h-120px leading-5">
                  Server : {{ chapterList.server_name }} <br />
                  Total Chapter : {{ chapterList.chapter.length }} <br />
                  Download : {{ chapterList.manga_id }} <br />
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
                  :disabled="chapterList.manga_id == 0"
                  status="warning"
                  @click="$router.push(`chapter/${chapterList?.manga_id}`)"
                  >Read</a-button
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
                current: tablePageCurrent,
              }"
              :loading="tableLoading"
              @page-change="tablePageChange"
              class="select-none"
              size="mini"
              @row-click="(c)=>clickRowTable(<types.ChapterList>c)"
            >
              <template #columns>
                <a-table-column
                  :width="75"
                  align="center"
                  title="Chapter"
                  cell-class="chap"
                  data-index="chapter"
                />
                <a-table-column :width="50" title="Vol" data-index="volume" />
                <a-table-column title="Judul Chapter" data-index="title">
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <a-typography-paragraph
                      style="margin: 0"
                      :ellipsis="{
                        rows: 1,
                        showTooltip: true,
                      }"
                    >
                      {{ record.title }}
                    </a-typography-paragraph>
                  </template>
                </a-table-column>

                <a-table-column title="Group" data-index="group_name">
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <a-typography-paragraph
                      style="margin: 0"
                      :ellipsis="{
                        rows: 1,
                        showTooltip: true,
                      }"
                    >
                      {{ record.group_name }}
                    </a-typography-paragraph>
                  </template>
                </a-table-column>

                <a-table-column :width="135" align="center" title="Release">
                  <template #cell="{ record }: { record: types.ChapterList }">
                    {{
                      DateApp.NewDate(record.timestamp * 1000).formatTimeAgo()
                    }}
                  </template>
                </a-table-column>
                <a-table-column
                  :width="75"
                  align="center"
                  title="Bahasa"
                  data-index="language"
                  :sortable="{ sortDirections: ['ascend', 'descend'] }"
                >
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <i-twemoji-flag-for-flag-indonesia
                      class="id"
                      v-if="record.language == 'Indonesia'"
                    />
                    <i-twemoji-flag-for-flag-united-kingdom v-else />
                  </template>
                </a-table-column>
                <a-table-column :width="70" align="center" title="Status">
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <span
                      class="bg-green-600 px-2 rounded-lg"
                      v-if="record.status"
                      >OK</span
                    >
                    <span v-else class="bg-red-600 px-2 rounded-lg">No</span>
                  </template>
                </a-table-column>
                <a-table-column :width="50" title="Check" align="center">
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
                  :width="110"
                  class="min-h-20px"
                  title="Download"
                  align="center"
                >
                  <template #cell="{ record }: { record: types.ChapterList }">
                    <a-button
                      size="mini"
                      @click.stop="downloadChapter(record)"
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
                  @click-download="downloadChapter()"
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
  GetChapterMdexPagination,
  CheckChapterDB,
} from '@wails/go/download/Download'
import { Message, Modal } from '@arco-design/web-vue'
import '@arco-design/web-vue/es/message/style/index'
import { types } from '@wails/go/models'
import { UseTable, UseServer } from '@/composable/downloads/download'
import { useDownloadState } from '@/store/global'
import { promiseTimeout, toValue } from '@vueuse/core'
import { Download } from '@/composable/downloads/wrapper'
import type { EventChap, EventPage } from '@/type/download'

//// STARTING CODE BOOTUP ////
const { tableDownload, tableLoading, tablePageSize, tablePageCurrent } =
  UseTable()
const { servers, getSelectedServer } = UseServer()

// GLOBAL VALUE
const tabsActive = ref(1)
const { chapterList, urldata, selectedServer } = useDownloadState()

const inititalProgressDL = {
  modal_dl: false,
  chapter: 0,
  index_chap: 0,
  index_page: 0,
  total_chap: 0,
  total_page: 0,
  chap: 0,
  page: 0,
}

const progress_dl = reactive({ ...inititalProgressDL })

// FOR NEXT PROGRESS
/*
OK  1. status check
OK  2. status download
OK  3. save db
next  4. retry
*/

// GET CHAPTER DATA
const { GetPasteData } = useClipboardData()
const dl = new Download()

let afterGetChapURL = ''
const goGetchDownload = useDebounceFn(async () => {
  // auto fetch paste URL(only)
  const paste = await GetPasteData()
  const pasteURL = IsURL(paste)
  if (pasteURL != null) {
    urldata.value = pasteURL.href
    //here testing download chapter
  }

  // wait watch serverID
  await promiseTimeout(26)

  //testing download chapter
  const server = getSelectedServer(selectedServer.value - 1)
  if (server && urldata.value != '') {
    const o = new types.Option()
    o.url = urldata.value
    o.server_name = server.name
    dl.GetChap(o, server)
      .then(res => {
        // reset table
        if (afterGetChapURL != '' && urldata.value != afterGetChapURL) {
          tableDownload.length = 0
          selected_chapter_url.value = []
          tablePageCurrent.value = 1
        }

        // top empty table if new manga
        if (chapterList.value?.manga != res.manga) {
          tableDownload.length = 0
        }
        //warning torefs expect reactive
        chapterList.value = res
        res.chapter.forEach(item => {
          tableDownload.push(item)
        })
        if (res.total > res.chapter.length) {
          var dummyChap = new types.ChapterList()
          dummyChap.chapter = ''
          dummyChap.id = ''
          dummyChap.group_name = ''
          dummyChap.language = ''
          dummyChap.timestamp = 0
          dummyChap.title = ''
          dummyChap.volume = ''

          const dumArr = Array.from(
            { length: res.total - res.chapter.length },
            (_, i) => dummyChap
          )
          tableDownload.splice(res.chapter.length, 0, ...dumArr)
          chapterList.value.chapter.splice(res.chapter.length, 0, ...dumArr)
        }
      })
      .catch(err => {
        console.log(err)
        Message.error(err)
      })
      .finally(() => {
        afterGetChapURL = toValue(urldata)
      })
  }
}, 500)

// if chapter is avalable insert table
chapterList.value?.chapter.forEach(item => {
  tableDownload.push(item)
})

// DOWNLOAD SINGLE OR MULTIPLE CHAPTER
const mdexPageServer = ref(false) // modelValue mdex datasaver server
const downloadChapter = (c: types.ChapterList | null = null) => {
  const server = getSelectedServer(selectedServer.value - 1)
  if (server && urldata.value != '' && chapterList.value != null) {
    var param = new types.Chapter(chapterList.value)
    param.datasaver = mdexPageServer.value

    if (c == null) {
      param.chapter = JSON.parse(JSON.stringify(selected_chapter_url.value))
    } else {
      param.chapter = [JSON.parse(JSON.stringify(c))]
    }

    // call
    Object.assign(progress_dl, inititalProgressDL)
    progress_dl.modal_dl = true
    progress_dl.index_page = 0

    dl.GetPage(param, server)
      .then(res => {
        console.log(res)
        if (res.status_dl == false || res.fail_chap.length > 0) {
          Modal.error({
            title: 'Error Some Download',
            okText: 'OK',
            content: () => [h('div', [res.manga, h('br'), res.error])],
          })
        } else {
          progress_dl.modal_dl = false
          Message.success('Berhasil Download')
          if (chapterList.value) {
            if (chapterList.value.manga_id == 0) {
              chapterList.value.manga_id = res.manga_id
            }
          }
        }
      })
      .catch(e => {
        console.log(e)
      })
    // end
  }
}

// Listening Download Event : dl_eventchap
EventsOn('dl_eventchap', (p: EventChap) => {
  progress_dl.chapter = p.chapter
  progress_dl.index_chap = p.index_chap
  progress_dl.total_chap = p.total_chap
  progress_dl.total_page = p.total_page
  progress_dl.index_page = 0
  progress_dl.chap = p.index_chap / p.total_chap
  console.log(
    '%cEventChap : %o',
    'color:orange',
    JSON.parse(JSON.stringify(progress_dl))
  )
})

// Listening Download Event : dl_eventpage
EventsOn('dl_eventpage', (p: EventPage) => {
  progress_dl.index_page += 1
  progress_dl.page = progress_dl.index_page / progress_dl.total_page
  console.log(
    '%cEventPage : index : %o = total %o == %i%',
    'color:Tomato',
    progress_dl.index_page,
    progress_dl.total_page,
    100 * progress_dl.page
  )

  if (progress_dl.page == 1) {
    const i = tableDownload.findIndex(i => i.chapter == p.chapter.toString())
    if (chapterList.value) {
      chapterList.value.chapter[i].status = true
      chapterList.value.chapter[i].check = false
    }
    tableDownload.forEach((item, i) => {
      if (item.chapter == p.chapter.toString()) {
        tableDownload[i].status = true
        tableDownload[i].check = false
      }
    })
    //clear selected
    const ix = selected_chapter_url.value.findIndex(
      item => item.chapter == p.chapter.toString()
    )
    selected_chapter_url.value.splice(ix, 1)
  }
})

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

/// tablePageChange dynamic load rest of chapter download for mangadex
const tablePageChange = (p: number) => {
  tablePageCurrent.value = p
  if (chapterList.value) {
    const limit = 30
    const pagesize = tablePageSize.value //5

    const seq = limit / pagesize
    const pSeq = Sequence(seq, p)
    const offset = pSeq * limit - limit
    // const offset = (p - 1) * (limit / pagesize) + 1;
    const tbl = tableDownload[offset]
    if (tbl.id == '') {
      tableLoading.value = true

      GetChapterMdexPagination(urldata.value, limit, offset)
        .then(res => {
          if (chapterList.value) {
            var cp = new types.Chapter()
            cp.manga = chapterList.value.manga
            cp.manga_id = chapterList.value.manga_id
            cp.chapter = JSON.parse(JSON.stringify(res))
            CheckChapterDB(cp)
              .then(cpp => {
                tableDownload.splice(offset, cpp.chapter.length, ...cpp.chapter)
                if (chapterList.value) {
                  chapterList.value.chapter.splice(
                    offset,
                    cpp.chapter.length,
                    ...cpp.chapter
                  )
                }
                tableLoading.value = false
              })
              .catch(() => {
                tableDownload.splice(offset, res.length, ...res)
                if (chapterList.value) {
                  chapterList.value.chapter.splice(offset, res.length, ...res)
                }
                tableLoading.value = false
              })
          }
        })
        .catch(e => {
          console.log(e)
          tableLoading.value = false
        })
    }
  } //chaplist
}

//clearDownload
const clearDownload = () => {
  urldata.value = ''
  selectedServer.value = 1 + 1 // temporary fix + 1 => not use id 1
  chapterList.value = null
  tableDownload.length = 0
  tablePageCurrent.value = 1
  selected_chapter_url.value = []
}

//watch URL DATA => auto select servers
watchDebounced(
  [urldata, selectedServer],
  ([newURL]) => {
    const isURL = IsURL(newURL)
    if (isURL) {
      const res = servers.value?.find(item => item.url == isURL.host)
      if (res) {
        // res.id + 1 // temporary fix warning
        selectedServer.value = res.id + 1
      }
    }
  },
  { debounce: 25 }
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
#urlinput .arco-input-append {
  padding: 0;
}
</style>

<style lang="less" scoped>
:deep(tr:has(.id)) {
  td {
    background-color: rgba(var(--black-6), 0.2) !important;
  }
  .chap {
    color: rgb(var(--primary-6));
  }
}
:deep(tr:hover) {
  td {
    background-color: rgba(var(--primary-6), 0.1) !important;
  }
}
</style>
