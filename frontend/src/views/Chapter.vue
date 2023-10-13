<template>
  <div class="text-left px-2 pt-2">
    <div class="flex pb-2">
      <div v-if="result" class="w-45 mr-2">
        <img
          draggable="false"
          class="select-none object-cover w-180px h-250px"
          :src="`/file/${MangaTitleURL(result.title)}/cover.webp`"
        />
      </div>
      <div class="w-full select-none">
        <a-typography>
          <a-typography-title
            :ellipsis="{
              rows: 1,
              showTooltip: true,
            }"
            :heading="5"
            style="margin-top: 0.25rem"
            >{{ result?.title }}
          </a-typography-title>
          <a-typography-paragraph>
            Status Complete :
            <strong>
              <span v-if="result?.status_end">Yes</span>
              <span class="text-red-600" v-else>NO</span>
            </strong>
            <br />
            <div v-if="result">
              Manga Alternative Title :
              <div class="h-105px">
                <a-list
                  :bordered="false"
                  size="small"
                  :max-height="120"
                  :scrollbar="true"
                >
                  <a-list-item
                    style="padding: 0.05rem; font-size: 12px"
                    v-for="(alt, i) in result.alter"
                    :key="i"
                  >
                    <a-list-item-meta :title="alt.title"> </a-list-item-meta>
                  </a-list-item>
                  <template #empty> No </template>
                </a-list>
              </div>
            </div>
          </a-typography-paragraph>
        </a-typography>
        <div class="flex">
          <div class="w-full pr-2">
            <a-typography-paragraph
              style="margin: 0"
              :ellipsis="{
                rows: 2,
                showTooltip: true,
              }"
            >
              {{ (selectedChapter.length ? 'Chapter ' : '') + selectedChapter }}
            </a-typography-paragraph>
          </div>
          <div class="h-44px grid items-end">
            <a-space>
              <a-button-group>
                <a-button @click="selectToggle">Select</a-button>
                <a-dropdown-button @click="showModalSelect">
                  Multi
                  <template #icon>
                    <icon-down />
                  </template>
                  <template #content>
                    <a-doption @click="clearSelect">Clear</a-doption>
                    <a-doption @click="selectAll">Select All</a-doption>
                  </template>
                </a-dropdown-button>
              </a-button-group>
              <a-button-group>
                <a-button @click="showModalAdd" status="success">Add</a-button>
                <a-button
                  @click="deleteClick"
                  status="danger"
                  :disabled="!(selectedKeys.length > 0)"
                  >Delete</a-button
                >
              </a-button-group>
            </a-space>
          </div>
        </div>
      </div>
    </div>

    <div>
      <div>
        <a-table
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
              cell-class="chap"
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
    <!-- modal add chapter -->
    <a-modal
      v-model:visible="modalState.modalAdd"
      :mask-closable="false"
      title="Add Chapter"
      @before-ok="handleAdd"
    >
      <div class="text-center">
        <a-space direction="vertical" class="w-full">
          <a-input-group class="w-full">
            <a-input-number
              ref="refFromChapter"
              placeholder="Please Enter"
              v-model="mAdd.chapter"
              name="madd_chapter"
              @keydown.-.prevent="addMulti"
              :min="0"
            />
            <a-input-number
              ref="refToChapter"
              v-if="mAdd.multi"
              placeholder="Please Enter"
              v-model="mAdd.chapterTo"
              :min="mAdd.chapter"
              @keydown.enter.prevent=""
            />
          </a-input-group>
          <div>
            <a-space>
              <a-radio-group type="button" v-model="mAdd.lang">
                <a-radio value="English"
                  ><i-twemoji-flag-for-flag-united-kingdom /> English</a-radio
                >
                <a-radio value="Indonesia"
                  ><i-twemoji-flag-for-flag-indonesia /> Indonesia</a-radio
                >
              </a-radio-group>
              <a-radio-group type="button" v-model="mAdd.multi">
                <a-radio :value="false">Single</a-radio>
                <a-radio :value="true">Multi</a-radio>
              </a-radio-group>
              <a-radio-group
                :disabled="!mAdd.multi"
                type="button"
                v-model="mAdd.lvl"
              >
                <a-radio value="1">1</a-radio>
                <a-radio value="0.5">0.5</a-radio>
                <a-radio value="0.1">0.1</a-radio>
              </a-radio-group>
            </a-space>
          </div>
          <div class="grid grid-cols-4 gap-2">
            <a-input-number
              :disabled="modalState.addState.disableLoop"
              placeholder="Loop Level"
              class="col-span-3"
              :input-attrs="{ class: 'text-center' }"
              v-model="mAdd.looplvl"
              :min="0"
              mode="button"
            />
            <a-input-number
              :disabled="true"
              :input-attrs="{ class: 'text-center' }"
              v-model="addChapterList.length"
              :min="0"
            />
          </div>
        </a-space>
      </div>
    </a-modal>
  </div>
</template>
<script setup lang="ts">
import { GetMangaWithChapter } from '@wails/go/manga/Manga'
import { manga } from '@wails/go/models'
import { IconDown } from '@arco-design/web-vue/es/icon'
import { MangaTitleURL, DateApp, GetBreakPoints } from '@/composable/helper'
import {
  type PaginationProps,
  type TableRowSelection,
  type InputNumberInstance,
  Modal,
} from '@arco-design/web-vue'

interface Madd {
  chapter: number | undefined
  chapterTo: number | undefined
  lang: string
  multi: boolean
  lvl: string
  looplvl: number
}

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

const rowSelectionDefault: TableRowSelection = {
  type: 'checkbox',
  showCheckedAll: true,
  onlyCurrent: false,
}

const rowSelection = ref<TableRowSelection | undefined>()
const selectedKeys = ref<number[]>([])

const modalState = reactive({
  modalSelect: false,
  modalAdd: false,
  select: {
    from: '',
    to: '',
  },
  addState: {
    disableLoop: true,
    totalChaps: 1,
  },
})

const mAdd = reactive<Madd>({
  chapter: undefined,
  chapterTo: undefined,
  lang: 'English',
  multi: false,
  lvl: '1',
  looplvl: 1,
})
const selectToggle = () => {
  if (rowSelection.value) {
    rowSelection.value = undefined
    selectedKeys.value = []
  } else {
    rowSelection.value = rowSelectionDefault
  }
}
const showModalSelect = async () => {
  if (rowSelection.value == undefined) {
    rowSelection.value = rowSelectionDefault
  }
  modalState.select.from = ''
  modalState.select.to = ''
  modalState.modalSelect = true
  await nextTick()
  document.getElementById('fromInput')?.focus()
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
      // console.log('Enter ditekan')
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
  if (newm) {
    selectedKeys.value.push(...newm)
  }
  rowSelection.value = rowSelectionDefault
}

const refFromChapter = ref<InputNumberInstance>()
const refToChapter = ref<InputNumberInstance>()
const showModalAdd = async () => {
  mAdd.chapter = undefined
  mAdd.chapterTo = undefined
  mAdd.lang = 'English'
  mAdd.multi = false
  mAdd.lvl = '1'
  mAdd.looplvl = 1
  const lang = result.value ? result.value.chapter[0].language.lang : 'English'
  mAdd.lang = lang
  modalState.modalAdd = true
  await nextTick()
  refFromChapter.value?.focus()
}

const addMulti = async () => {
  mAdd.multi = true
  await nextTick()
  refToChapter.value?.focus()
}

watch(
  () => mAdd.lvl,
  lvl => {
    if (lvl == '0.1') {
      modalState.addState.disableLoop = false
    } else {
      modalState.addState.disableLoop = true
    }
  }
)

const addChapterList = computed<manga.Chapter[]>(() => {
  const latestChap = result.value ? result.value.chapter[0] : null
  const chapList: manga.Chapter[] = []
  const f = mAdd.chapter
  const t = mAdd.chapterTo
  const cl = ref(1)
  if (f) {
    if (t && t > f) {
      //no decimal number.integer
      //calculate cf
      switch (mAdd.lvl) {
        case '1':
          cl.value = t - f + 1
          break
        case '0.5':
          cl.value = (t - f) * 2 + 1
          break
        default:
          cl.value = t - f + (t - f) * mAdd.looplvl + 1
          break
      }
      // console.log('total cl', cl.value, chapList)
      //check added number already in chapter
      for (let i = 0; i < cl.value; i++) {
        const m = new manga.Chapter(latestChap)
        switch (mAdd.lvl) {
          case '1':
            m.chapter = f + i
            break
          case '0.5':
            m.chapter = i / 2 + f
            break
          default:
            m.chapter = chapSeq(i, mAdd.looplvl, f)
            break
        }
        //not add if already in chapter
        if (inChapter(m.chapter)) {
          chapList.push(m)
        }
      }
    } else {
      const m = new manga.Chapter(latestChap)
      m.chapter = f
      //not add if already in chapter
      if (inChapter(f)) {
        chapList.push(m)
      }
    }
  }

  return chapList
})

const handleAdd = async () => {
  console.log(JSON.parse(JSON.stringify(addChapterList.value)))
  await new Promise(resolve => setTimeout(resolve, 600))

  //reset to default
  return true
}

const chapSeq = (n: number, length: number, start: number): number => {
  const group = Math.floor(n / (length + 1))
  const index = n % (length + 1)

  return start + group + index * 0.1
}

const inChapter = (i: number): boolean => {
  if (result.value) {
    return result.value.chapter.findIndex(it => it.chapter == i) == -1
  } else {
    return false
  }
}

const deleteClick = () => {
  console.log('delete')
  const allChap = tableData.length == selectedKeys.value.length
  Modal.warning({
    title: 'Delete Confirmation',
    content: () => [
      h('div', 'Are you sure want to Delete Selected Chapter?'),
      allChap
        ? h('div', [
            h('span', 'Deleting All Chapter make Manga with '),
            h('span', { style: { color: 'red' } }, 'No Chapter'),
          ])
        : '',
    ],
    okText: 'Delete',
    cancelText: 'No',
    hideCancel: false,
    bodyClass: 'text-center',
    maskClosable: false,
  })
}

const selectedChapter = computed(() => {
  return tableData
    .filter(it => selectedKeys.value.includes(it.id))
    .map(it => {
      return it.chapter
    })
    .reverse()
})
</script>

<style lang="less" scoped>
:deep(tr:nth-child(even)) {
  .chap {
    color: var(--app-main);
  }
}

:deep(tr:hover) {
  td {
    background-color: rgba(var(--primary-6), 0.1) !important;
  }
}
</style>
