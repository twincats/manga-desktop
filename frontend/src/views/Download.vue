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
                    @dblclick="openBrowser(item.url)"
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
                  :disabled="chapterList.manga_id == null"
                  status="warning"
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
              }"
              :loading="tableLoading"
              @page-change="tablePageChange"
              @change="(data, extra) => tableChange(<types.ChapterList[]>data, extra)"
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
                    <span class="bg-red-600 px-2 rounded-lg">No</span>
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
                    <span v-else class="text-red-600"><icon-minus /></span>
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
        <div class="h-[calc(100vh-7.5rem)] px-2">
          <a-input-search
            allow-clear
            class="w-full mb-2"
            placeholder="URL"
            v-model="url"
            search-button
            @search="fetchBrowserData"
          />
          <div
            ref="web"
            class="h-[calc(100vh-11.2rem)] overflow-auto p-2"
            style="background-color: var(--color-bg-2)"
            v-html="browserData?.body"
          ></div>
        </div>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup lang="ts">
import { IconCheck, IconMinus } from '@arco-design/web-vue/es/icon'
import {
  useClipboardData,
  IsURL,
  DateApp,
  GetBreakPoints,
  UseParseDom,
  Sequence,
} from '@/composable/helper'
import { GetServer } from '@wails/go/manga/Manga'
import { WebBrowser } from '@wails/go/tool/Web'
import { EventsOn } from '@wails/runtime/runtime'

import {
  GetChapter,
  GetChapterMdexPagination,
  GetPage,
} from '@wails/go/download/Download'
import { Message, TableChangeExtra } from '@arco-design/web-vue'
import '@arco-design/web-vue/es/message/style/index'
import { manga, tool, types } from '@wails/go/models'

// const cover = ref('/file/Bar Flowers/cover.webp')
const chapterList = ref<types.Chapter | null>(null)
const tableDownload = reactive<types.ChapterList[]>([])
const tablePageSize = ref(5)
const tableLoading = ref(false)
const { lg } = GetBreakPoints()
const tabsActive = ref(1)
const mdexPageServer = ref(false)
const selected_chapter_url = ref<types.ChapterList[]>([])

if (lg.value) {
  tablePageSize.value = 5
} else {
  tablePageSize.value = 15
}

watch(lg, nlg => {
  if (nlg) {
    tablePageSize.value = 5
  } else {
    tablePageSize.value = 15
  }
})

//INTERNAL Function
const getSelectedServer = (id: number): manga.Server | undefined => {
  return servers.value?.find(item => item.id == id)
}

///TESTING FOR TEMPORARY FILL DATA DOWNLOAD
const testDownload = (c: types.ChapterList | null = null) => {
  if (c == null) {
    console.log('Download multiple')
  } else {
    console.log('download single')
    const server = getSelectedServer(selectedServer.value)
    if (server && urldata.value != '') {
      GetPage({
        url: c.id,
        server_name: server.name,
        datasaver: mdexPageServer.value,
      })
        .then(res => {
          console.log(res)
        })
        .catch(e => {
          console.log(e)
        })
      EventsOn('dlProgress', p => console.log(p))
    }
  }
}
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
  // activateMultiDownload.value = !activateMultiDownload.value
}

const openBrowser = (uri: string) => {
  console.log(url)
  tabsActive.value = 2
  url.value = 'https://' + uri
  fetchBrowserData()
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

//get server
const servers = ref<manga.Server[] | null>(null)
GetServer().then(res => {
  servers.value = res.filter(item => item.status_active === true)
})

const urldata = ref('')
const selectedServer = ref(1)
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

//clearDownload
const clearDownload = () => {
  urldata.value = ''
  selectedServer.value = 1
  chapterList.value = null
  tableDownload.length = 0
}

//watch URL => auto select servers
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

//TAB 2
const web = ref<HTMLElement | null>(null)
const url = ref('')
const browserData = ref<tool.Web | null>(null)
const fetchBrowserData = () => {
  if (url.value) {
    WebBrowser(url.value).then(res => {
      const htmlDoc = UseParseDom(res.body)
      const v = htmlDoc.getElementsByTagName('h1')
      console.log(v.item(0)?.innerText)
    })
  }
}
useEventListener(web, 'click', e => {
  e.preventDefault()
  if (e.target instanceof HTMLAnchorElement) {
    console.log(e.target.href)
  } else {
    const path1 = e.composedPath()[1]
    if (path1 instanceof HTMLAnchorElement) {
      console.log(path1.href)
    }
  }
})
</script>

<style lang="less">
#download {
  .arco-input-append {
    padding: 0 0 0 0.25rem;
  }
}

#servers {
  .arco-radio-group-button {
    display: flex;
  }

  .arco-radio-button:nth-of-type(8n + 1)::before {
    content: none;
  }
}

.icons {
  font-size: 1rem;
  line-height: 1rem;
}
</style>
