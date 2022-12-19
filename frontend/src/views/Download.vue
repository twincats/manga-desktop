<template>
  <div
    id="download"
    class="m-1"
    style="--primary-6: 255, 117, 24; --primary-5: 204, 92, 18"
  >
    <a-tabs default-active-key="1">
      <a-tab-pane key="1" title="Download">
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
        <div class="h-[calc(100vh-16rem)] px-2 overflow-auto">
          <div
            class="min-h-240px mb-2"
            style="background-color: var(--color-bg-2)"
          >
            <div class="text-lg p-2 text-center">
              Uragirareta S Rank Boukensha no Ore wa, Aisuru Dorei no Kanojora
              to Tomoni Dorei dake no Harem Guild o Tsukuru
            </div>
            <div class="flex">
              <div class="w-150px px-2">
                <img class="w-full" :src="cover" alt="cover" />
              </div>
              <div class="w-full flex flex-col">
                <div class="h-full">
                  Server : Mangadex <br />
                  Total Chapter : 42 <br />
                  Download : <br />
                  Selected : <br />
                </div>
                <div class="p-2 pl-0" v-if="selectedServer == 1">
                  Mangadex Server :
                  <a-radio-group size="small">
                    <a-radio value="A">Original</a-radio>
                    <a-radio value="B">Data Saver</a-radio>
                  </a-radio-group>
                </div>
              </div>
              <div class="p-2 self-end">
                <a-button status="warning">Read</a-button>
              </div>
            </div>
          </div>
          <div>
            <a-table
              :data="tableDownload"
              :pagination="{ pageSize: tablePageSize, size: 'mini' }"
              size="mini"
              :scroll="{ x: 1200 }"
              @row-click="(c)=>clickRowTable(<TableDownload>c)"
            >
              <template #columns>
                <a-table-column
                  fixed="left"
                  :width="100"
                  align="center"
                  title="Chapter"
                  data-index="chapter"
                />
                <a-table-column :width="50" title="Vol" data-index="vol" />
                <a-table-column title="Judul Chapter" data-index="title" />
                <a-table-column
                  :width="120"
                  align="center"
                  title="Bahasa"
                  data-index="bahasa"
                />
                <a-table-column title="Group" data-index="group" />
                <a-table-column
                  fixed="right"
                  :width="135"
                  align="center"
                  title="Release"
                  data-index="release"
                />
                <a-table-column
                  fixed="right"
                  :width="70"
                  align="center"
                  title="Status"
                  data-index="status"
                />
                <a-table-column
                  :fixed="'right'"
                  :width="50"
                  title="Check"
                  align="center"
                >
                  <template #cell>
                    <span v-if="true" class="text-green-600"
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
                  <template #cell="{ record }: { record: TableDownload }">
                    <a-button
                      size="mini"
                      @click="testDownload"
                      :disabled="activateMultiDownload"
                      >Download</a-button
                    >
                  </template>
                </a-table-column>
              </template>
              <template #th="{ column }">
                <custom-download-th
                  :active="activateMultiDownload"
                  :column="column"
                />
              </template>
            </a-table>
          </div>
        </div>
      </a-tab-pane>
      <a-tab-pane key="2" title="Web">
        <div class="h-[calc(100vh-7.5rem)] px-2">
          <a-input-search
            allow-clear
            class="w-full mb-2"
            placeholder="URL"
            search-button
          />
          <div
            class="h-[calc(100vh-11.2rem)] overflow-auto p-2"
            style="background-color: var(--color-bg-2)"
          >
            <div id="webID"></div>
          </div>
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
} from '@/composable/helper'
import { GetServer } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'

interface TableDownload {
  chapter: number
  vol: number
  title: string
  bahasa: string
  group: string
  release: string
  status: boolean
}

function makeid(length: number) {
  var result = ''
  var characters =
    'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  var charactersLength = characters.length
  for (var i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength))
  }
  return result
}

const cover = ref('/file/Bar Flowers/cover.webp')
const tableDownload = reactive<TableDownload[]>([])
const tablePageSize = ref(5)
const { lg } = GetBreakPoints()

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

///TESTING FOR TEMPORARY FILL DATA DOWNLOAD
function randomDate(start: Date, end: Date) {
  return new Date(
    start.getTime() + Math.random() * (end.getTime() - start.getTime())
  )
}

for (let i = 1; i <= 20; i++) {
  const mdate = DateApp.NewDate(
    randomDate(new Date(2020, 0, 1), new Date()).toString()
  ).formatTimeAgo()
  tableDownload.push({
    chapter: i,
    bahasa: 'English',
    vol: i + 1,
    group: 'Achul',
    release: mdate ? mdate : '',
    status: false,
    title: makeid(15),
  })
}

const activateMultiDownload = ref(false)
const testDownload = () => {
  activateMultiDownload.value = !activateMultiDownload.value
}
const clickRowTable = (c: TableDownload) => {
  activateMultiDownload.value = !activateMultiDownload.value
}
///TESTING

//get server
const servers = ref<manga.Server[] | null>(null)
GetServer().then(res => {
  servers.value = res.filter(item => item.status_active == true)
})

const urldata = ref('')
const selectedServer = ref(1)
const { GetPasteData } = useClipboardData()
const goGetchDownload = () => {
  // auto fetch paste URL(only)
  GetPasteData().then(res => {
    const pasteURL = IsURL(res)
    if (pasteURL) {
      urldata.value = pasteURL.href
    }
  }) //
}

//clearDownload
const clearDownload = () => {
  urldata.value = ''
  selectedServer.value = 1
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
