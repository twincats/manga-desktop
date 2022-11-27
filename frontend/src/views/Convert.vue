<template>
  <div class="flex flex-col h-full">
    <div>
      <div id="tool" class="text-left">
        <div class="w-150">
          <div>
            Quality: &nbsp;&nbsp;
            <strong style="color: rgb(var(--primary-6))">{{
              data.quality
            }}</strong>
          </div>
          <div class="px-2">
            <a-slider
              :min="60"
              :show-ticks="true"
              :step="5"
              :max="90"
              v-model="data.quality"
            />
          </div>
        </div>
        <div class="w-full min-h-40px">
          <div>
            Resize : &nbsp;&nbsp;
            <strong
              v-if="data.resizeStatus"
              style="color: rgb(var(--primary-6))"
              >{{ data.resize }}</strong
            >
          </div>
          <div>
            <div v-if="data.resizeStatus" class="flex">
              <a-checkbox class="mr-4" v-model="data.resizeStatus"></a-checkbox>
              <a-slider
                :min="900"
                :max="1400"
                :show-ticks="true"
                :step="50"
                v-model="data.resize"
                class="mr-1"
              />
            </div>
            <a-button @click="data.resizeStatus = true" v-else size="mini" long
              >Resize</a-button
            >
          </div>
        </div>
        <div style="flex: 0 0 132px; padding-right: 0">
          <div>Delete :</div>
          <a-space direction="vertical" size="mini">
            <a-radio-group v-model="data.delete">
              <a-radio :value="true">Yes</a-radio>
              <a-radio :value="false">No</a-radio>
            </a-radio-group>
          </a-space>
        </div>
        <div style="flex: 0 0 160px; padding-right: 0">
          <div>Convert Engine :</div>
          <a-space direction="vertical" size="mini">
            <a-radio-group v-model="data.covert">
              <a-radio :value="1">libwebp</a-radio>
              <a-radio :value="2">Ext</a-radio>
            </a-radio-group>
          </a-space>
        </div>
        <div style="flex: 0 0 188px; padding-right: 0">
          <div>Format Image :</div>
          <a-space direction="vertical" size="mini">
            <a-checkbox-group v-model="data.format">
              <a-checkbox :value="1">*.jpg, *png</a-checkbox>
              <a-checkbox :value="2">*webp</a-checkbox>
            </a-checkbox-group>
          </a-space>
        </div>
      </div>
    </div>
    <div class="h-20 mb-1" style="background-color: var(--color-bg-4)">
      <a-input
        v-model="data.title"
        size="small"
        :input-attrs="{ class: 'text-center' }"
        placeholder="Please Select Manga"
        class="mt-1"
        readonly
      />
      <a-input
        ref="hiddenSearch"
        v-model="output.hiddenSearch"
        size="small"
        :input-attrs="{ class: 'text-center' }"
        placeholder="Search Here"
        class="my-1"
        allow-clear
      />
    </div>
    <div
      id="mtable"
      class="h-full overflow-auto p-3 mb-1 text-left"
      style="background-color: var(--color-bg-4)"
    >
      <a-table
        v-if="tableManga.length > 0"
        size="small"
        :show-header="false"
        :pagination="false"
        :columns="columnManga"
        :data="tableMangaFilter"
        @row-click="(c:TableManga)=>selectManga(c)"
      >
        <template v-slot:tr="{ rowIndex }">
          <tr class="my-tr" :id="'manga-' + rowIndex" />
        </template>
        <template #empty>
          <div class="arco-empty">
            <div class="arco-empty-image">
              <i-mdi-emoticon-sad-outline />
            </div>
            <div class="arco-empty-description">
              <div class="my-2">
                No Manga Data Filtered : <br />
                "{{ output.hiddenSearch }}"
              </div>
              <a-button
                @click="output.hiddenSearch = ''"
                cllass="mt-2"
                size="mini"
                >Clear Filter</a-button
              >
            </div>
          </div>
        </template>
      </a-table>
      <div
        v-else
        class="flex h-full justify-center"
        style="background-color: var(--color-bg-2)"
      >
        <a-space>
          <a-spin :size="32" />
          <a-spin :size="32" />
          <a-spin :size="32" />
        </a-space>
      </div>
    </div>
    <div>
      <div id="status_convert" v-html="statusConvert"></div>
      <div id="mprogress" class="mt-1">
        <a-progress
          status="success"
          :show-text="false"
          :stroke-width="20"
          :percent="output.progress"
        />
      </div>
    </div>
    <div class="text-xs flex justify-between px-3 py-2 h-80px">
      <div class="status">
        <div style="min-width: 128px">
          Size Before :
          <strong class="text-red-500" v-if="output.sizeBefore > 0"
            >{{ output.sizeBefore }} MB</strong
          >
        </div>
        <div style="min-width: 117px">
          Size After :
          <strong class="text-blue-500" v-if="output.sizeAfter > 0"
            >{{ output.sizeAfter }} MB</strong
          >
        </div>
        <div style="min-width: 85px">
          Percent :
          <strong class="text-green-500" v-if="output.percent > 0"
            >{{ output.percent }}%</strong
          >
        </div>
      </div>
      <div class="w-full"></div>
      <div class="w-full"></div>
      <div class="bmenu">
        <a-space>
          <a-button @click="reset">Reset</a-button>
          <a-button>Cover Fix</a-button>
          <a-button>Save Log</a-button>
          <a-button type="primary" @click="clickTest">Convert</a-button>
        </a-space>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetMangas } from '@wails/go/manga/Manga'
import { MangaTitleURL } from '@/composable/helper'

/// COMPONENT INTERFACE
interface TableManga {
  id: number
  title: string
}

interface ConvertProps {
  mid?: number | null
}

///INITIAL VARIABLE
const initialData = {
  title: '',
  quality: 60,
  resize: 1000,
  resizeStatus: true,
  delete: true,
  covert: 1,
  format: [1],
}

const initialOutput = {
  sizeBefore: 0,
  sizeAfter: 0,
  percent: 0,
  progress: 0,
  hiddenSearch: '',
}

const { mid = null } = defineProps<ConvertProps>()
const data = reactive({ ...initialData })
const output = reactive({ ...initialOutput })
const statusConvert = ref('')
const hiddenSearch = ref<HTMLInputElement | null>(null)

const tableManga = reactive<TableManga[]>([])
const columnManga = [{ title: 'Manga', dataIndex: 'title' }]

///INITIAL METHODS
GetMangas().then(resp => {
  //map to change Title from DB to URL manga
  const manga = resp.map(item => {
    item.title = MangaTitleURL(item.title)
    return item
  })

  //sort mapped data
  manga.sort((a, b) => {
    if (a.title < b.title) {
      return -1
    }
    if (a.title < b.title) {
      return 1
    }
    return 0
  })

  //fill table
  manga.forEach(item => {
    tableManga.push({ title: item.title, id: item.id })
  })

  // auto select manga from params
  if (mid) {
    const sel = manga.find(item => item.id == mid)
    if (sel) {
      data.title = sel.title
    }
  }
})

onStartTyping(() => {
  hiddenSearch.value?.focus()
})

onMounted(() => {
  hiddenSearch.value?.focus()
})

///APPLICATION METHODS

//select manga
const selectManga = (c: TableManga) => {
  data.title = c.title
}

//reset button
const reset = () => {
  Object.assign(data, initialData)
  Object.assign(output, initialOutput)
}

const tableMangaFilter = computed(() => {
  const s = output.hiddenSearch
  return tableManga.filter(item => {
    return item.title.toLocaleLowerCase().includes(s.toLocaleLowerCase())
  })
})

//// CODE BELOW IS TESTING CODE

const sleep = (ms: number) => new Promise(r => setTimeout(r, ms))

const clickTest = async () => {
  for (let i = 1; i <= 10; i++) {
    statusConvert.value += '<span>Hello world</span><br/>'
    await sleep(1000)
  }
}
</script>

<style lang="less" scoped>
#tool {
  border-bottom: 1px solid white;
  display: flex;
  justify-content: space-between;
  > div {
    padding: 0.5rem;
    font-size: 0.75rem;
    display: flex;
    flex-direction: column;
    &:not(:last-child) {
      border-right: 1px solid white;
    }
    > div {
      height: 100%;
    }
  }
}
#status_convert {
  // class="h-25 p-3 overflow-y-scroll text-left"
  height: 5.5rem;
  overflow-y: scroll;
  text-align: left;
  padding: 0.75rem;
  background-color: var(--color-bg-4);
}
.status {
  width: 100%;
  display: flex;
  justify-content: space-between;
  div {
    text-align: left;
    align-self: center;
    padding-right: 0.75rem;
  }
}

.bmenu {
  display: flex;
}

.arco-radio,
button,
.arco-checkbox {
  font-size: 0.75rem !important;
}
</style>

<style>
#mprogress .arco-progress-line {
  border-radius: 0 !important;
}
#mprogress .arco-progress-line-bar {
  border-radius: 0 !important;
}
#mtable .arco-table-td {
  font-size: 0.75rem !important;
}
</style>
