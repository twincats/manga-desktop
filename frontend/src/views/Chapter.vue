<template>
  <div class="text-left px-2 pt-2">
    <div class="flex pb-2">
      <div v-if="result" class="w-45 mr-2">
        <a-image
          :preview="false"
          :width="180"
          :height="250"
          fit="cover"
          :src="`/file/${MangaTitleURL(result.title)}/cover.webp`"
        />
      </div>
      <a-typography class="w-full">
        <a-typography-title :heading="5" style="margin-top: 0.25rem"
          >{{ result?.title }}
        </a-typography-title>
        <a-typography-paragraph>
          Status Complete : {{ result?.status_end }} <br />
          <div v-if="result">
            Manga Alternative Title :
            <span v-if="result?.alter.length == 0">NO</span> <br />
            <ul v-if="result?.alter">
              <li v-for="(m, i) in result?.alter" :key="i">
                {{ m.title }}
              </li>
            </ul>
            {{ selectedKeys }}
          </div>
        </a-typography-paragraph>
      </a-typography>
      <div>
        <a-space>
          <a-dropdown-button @click="showModalSelect">
            Select
            <template #icon>
              <icon-down />
            </template>
            <template #content>
              <a-doption @click="clearSelect">Clear</a-doption>
              <a-doption @click="selectAll">Select All</a-doption>
            </template>
          </a-dropdown-button>
          <a-button-group>
            <a-button status="success">Add</a-button>
            <a-button status="danger">Delete</a-button>
          </a-button-group>
        </a-space>
      </div>
    </div>
    <div>
      <div>
        <a-table
          id="chapterTable"
          row-key="id"
          :data="result?.chapter"
          :row-selection="rowSelection"
          :pagination="pprop"
          v-model:selected-keys="selectedKeys"
          size="small"
        >
          <template #columns>
            <a-table-column
              v-if="false"
              :width="110"
              title="Chapter ID"
              data-index="id"
            ></a-table-column>
            <a-table-column
              :width="88"
              title="Chapter"
              data-index="chapter"
              align="center"
            >
            </a-table-column>
            <a-table-column
              title="Title"
              :ellipsis="true"
              :tooltip="true"
              data-index="title"
            >
            </a-table-column>
            <a-table-column
              title="Volume"
              data-index="volume"
              align="center"
            ></a-table-column>
            <a-table-column
              :ellipsis="true"
              :tooltip="true"
              title="Group"
              align="center"
              data-index="group"
            >
            </a-table-column>
            <a-table-column :width="185" title="Download Time">
              <template #cell="{ record }: { record: manga.Chapter }">
                {{
                  DateApp.NewDate(record.created_at.toString()).Format(
                    'DD-MM-YYYY HH:mm:ss'
                  )
                }}
              </template>
            </a-table-column>
            <a-table-column title="Bahasa" align="center">
              <template #cell="{ record }: { record: manga.Chapter }">
                <i-twemoji-flag-for-flag-indonesia
                  v-if="record.language.id == 2"
                />
                <i-twemoji-flag-for-flag-united-kingdom v-else />
              </template>
            </a-table-column>
            <a-table-column :width="100" title="Read" align="center">
              <template #cell="{ record }: { record: manga.Chapter }">
                <a-badge
                  :count="record.status_read ? 0 : 1"
                  dot
                  :offset="[0, -2]"
                >
                  <a-button
                    size="mini"
                    @click="$router.push(`/page/${record.id}`)"
                    >Read</a-button
                  >
                </a-badge>
              </template>
            </a-table-column>
          </template>
        </a-table>
      </div>
      <!-- <div class="col-span-2 pl-2">
        <a-space>
          <a-dropdown-button @click="showModalSelect">
            Select
            <template #icon>
              <icon-down />
            </template>
            <template #content>
              <a-doption @click="clearSelect">Clear</a-doption>
              <a-doption @click="selectAll">Select All</a-doption>
            </template>
          </a-dropdown-button>
          <a-button-group>
            <a-button status="success">Add</a-button>
            <a-button status="danger">Delete</a-button>
          </a-button-group>
        </a-space>
      </div> -->
    </div>
    <!-- modal select chapter -->
    <a-modal
      v-model:visible="modalState.modalSelect"
      title="Select Chapter Range"
      @ok="selectChapRange"
    >
      <div class="text-center">
        <a-input-group>
          <a-input
            @keydown="keySelectEvent"
            :input-attrs="{ id: 'fromInput' }"
            :style="{ width: '160px' }"
            v-model="modalState.select.from"
            placeholder="From"
          />
          <a-input
            @keydown="keySelectEvent"
            :input-attrs="{ id: 'toInput' }"
            :style="{ width: '160px' }"
            v-model="modalState.select.to"
            placeholder="To"
          />
        </a-input-group>
      </div>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { GetMangaWithChapter } from '@wails/go/manga/Manga'
import { manga } from '@wails/go/models'
import { IconDown } from '@arco-design/web-vue/es/icon'
import { MangaTitleURL, DateApp, GetBreakPoints } from '@/composable/helper'
import type { PaginationProps, TableRowSelection } from '@arco-design/web-vue'

const props = defineProps<{
  mid: string
}>()

const result = ref<manga.Manga>()

const bp = GetBreakPoints()

// const tableColumn = [{ title: 'Manga', dataIndex: 'title' }]
const tableData = reactive<manga.Chapter[]>([])

GetMangaWithChapter(Number(props.mid)).then(res => {
  result.value = res
  //sort desc
  res.chapter.sort((a, b) => Number(b.chapter) - Number(a.chapter))
  res.chapter.forEach(item => {
    tableData.push(item)
  })
  // console.log(res)
})

const pprop = reactive<PaginationProps>({
  hideOnSinglePage: true,
})

const setAutopaginationSize = () => {
  if (bp.lg.value) {
    pprop.pageSize = 9
  } else {
    pprop.pageSize = 17
  }
}
setAutopaginationSize()
watch(
  () => bp.lg.value,
  () => {
    setAutopaginationSize()
  }
)

const rowSelection = reactive<TableRowSelection>({
  type: 'checkbox',
  showCheckedAll: true,
  onlyCurrent: false,
})
const selectedKeys = ref<number[]>([])

const modalState = reactive({
  modalSelect: false,
  select: {
    from: '',
    to: '',
  },
})
const showModalSelect = async () => {
  modalState.modalSelect = true
  await nextTick()
  var d = document.getElementById('fromInput')
  d?.focus()
}

const keySelectEvent = (e: KeyboardEvent) => {
  const allowedKey = ['.', 'Backspace', 'Delete', 'ArrowLeft', 'ArrowRight']
  if (isNaN(Number(e.key)) && !allowedKey.includes(e.key)) {
    e.preventDefault()
    if (e.key == '-') {
      if ((e.target as HTMLElement).id == 'fromInput') {
        document.getElementById('toInput')?.focus()
      }
    } else if (e.key == 'Enter') {
      console.log('Enter ditekan')
      if ((e.target as HTMLElement).id == 'fromInput') {
        document.getElementById('toInput')?.focus()
      } else if ((e.target as HTMLElement).id == 'toInput') {
        // if (Number(modalState.select.from) > Number(modalState.select.to)) {
        //   console.log('tidak bisa angka lebih besar ke kecil')
        // } else {
        //   selectedKeys.value.push(1)
        // }
        selectChapRange()
      }
    }
  } else {
    if (
      isNaN(Number((e.target as HTMLInputElement).value + '.')) &&
      e.key == '.'
    ) {
      e.preventDefault()
    }
  }
}

const selectChapRange = () => {
  const from = Number(modalState.select.from)
  const to = Number(modalState.select.to)
  selectedKeys.value = []
  if (from <= to) {
    result.value?.chapter.forEach(it => {
      if (it.chapter >= from && it.chapter <= to) {
        selectedKeys.value.push(it.id)
      }
    })
  } else if (to <= from) {
    result.value?.chapter.forEach(it => {
      if (it.chapter >= to && it.chapter <= from) {
        selectedKeys.value.push(it.id)
      }
    })
  } else {
    console.log('error tidak sesuai')
  }
  modalState.modalSelect = false
}

const clearSelect = () => {
  selectedKeys.value = []
}

const selectAll = () => {
  const newm = result.value?.chapter.map(it => it.id)
  selectedKeys.value = newm ? newm : []
}
</script>

<style lang="less">
#chapterTable {
  tr:nth-child(even) {
    td:nth-child(2) {
      color: var(--app-main);
    }
  }
  tr:hover {
    td {
      background-color: rgba(var(--primary-6), 0.1);
    }
  }
}
</style>
