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
            :style="{ width: '100%' }"
            :input-attrs="{ class: 'text-center' }"
            placeholder="URL"
            allow-clear
          >
            <template #append>
              <a-space size="mini">
                <a-button type="secondary" class="px-2">
                  <span style="font-size: 1rem; line-height: 1rem">
                    <i-mdi:eraser />
                  </span>
                </a-button>
                <a-button type="secondary" class="px-2">
                  <span
                    style="
                      color: rgb(255, 117, 24);
                      font-size: 1rem;
                      line-height: 1rem;
                    "
                    ><i-mdi:cloud-search-outline
                  /></span>
                </a-button>
              </a-space>
            </template>
          </a-input>
        </div>
        <div class="h-[calc(100vh-11rem)] px-2 overflow-auto">
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
              <div class="w-full">
                Total Chapter : 42 <br />
                Download : <br />
                Selected : <br />
              </div>
              <div class="p-2 self-end">
                <a-button>Read</a-button>
              </div>
            </div>
          </div>
          <div>
            <a-table
              :data="tableDownload"
              :pagination="{ pageSize: 6 }"
              size="small"
            >
              <template #columns>
                <a-table-column title="chapter" data-index="chapter" />
                <a-table-column title="Vol" data-index="vol" />
                <a-table-column title="Judul Chapter" data-index="title" />
                <a-table-column title="Bahasa" data-index="bahasa" />
                <a-table-column title="Group" data-index="group" />
                <a-table-column title="Release" data-index="release" />
                <a-table-column title="Status" data-index="status" />
                <a-table-column :width="50" title="Check" align="center">
                  <template #cell>
                    <span v-if="true" class="text-green-600"
                      ><icon-check
                    /></span>
                    <span v-else class="text-red-600"><icon-minus /></span>
                  </template>
                </a-table-column>
                <a-table-column :width="100" title="Download" align="center">
                  <template #cell="{ record }: { record: TableDownload }">
                    <a-button size="mini">Download</a-button>
                  </template>
                </a-table-column>
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

for (let i = 1; i <= 10; i++) {
  tableDownload.push({
    chapter: i,
    bahasa: 'English',
    vol: i + 1,
    group: 'Achul',
    release: '2 Hari yang lalu',
    status: false,
    title: makeid(15),
  })
}
</script>

<style lang="less">
#download {
  .arco-input-append {
    padding: 0 0 0 0.25rem;
  }
}
</style>
