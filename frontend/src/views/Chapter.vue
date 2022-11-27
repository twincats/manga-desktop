<template>
  <div class="text-left px-2">
    <div class="flex">
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
          </div>
        </a-typography-paragraph>
      </a-typography>
      <div v-if="result" class="w-45">
        <img
          class="w-full mt-1"
          :src="`/file/${MangaTitleURL(result.title)}/cover.webp`"
          alt=""
        />
      </div>
    </div>
    <div>
      <a-table :data="result?.chapter" :pagination="pprop" size="small">
        <template #columns>
          <a-table-column
            :width="110"
            title="Chapter ID"
            data-index="id"
          ></a-table-column>
          <a-table-column :width="88" title="Chapter" align="center">
            <template #cell="{ record }">
              <span class="text-orange-500">{{ record.chapter }}</span>
            </template>
          </a-table-column>
          <a-table-column title="Title">
            <template #cell="{ record }: { record: manga.Chapter }">
              <div class="overflow-hidden h-5">{{ record.title }}</div>
            </template>
          </a-table-column>
          <a-table-column
            title="Volume"
            data-index="volume"
            align="center"
          ></a-table-column>
          <a-table-column title="Group" align="center">
            <template #cell="{ record }: { record: manga.Chapter }">
              <div class="overflow-hidden h-5">{{ record.group }}</div>
            </template>
          </a-table-column>
          <a-table-column :width="185" title="Download Time">
            <template #cell="{ record }: { record: manga.Chapter }">
              {{
                DateApp.NewDate(record.created_at).Format('DD-MM-YYYY HH:mm:ss')
              }}
            </template>
          </a-table-column>
          <a-table-column
            :width="10"
            title="Bahasa"
            data-index="language"
            align="center"
          >
            <template #cell="{ record }: { record: manga.Chapter }">
              <i-twemoji-flag-for-flag-indonesia
                v-if="record.language.lang_code == 'id'"
              />
              <i-twemoji-flag-for-flag-united-kingdom v-else />
            </template>
          </a-table-column>
          <a-table-column :width="50" title="Status" align="center">
            <template #cell="{ record }">
              <span v-if="record.status_read" class="text-green-600"
                ><icon-check
              /></span>
              <span v-else class="text-red-600"><icon-minus /></span>
            </template>
          </a-table-column>
          <a-table-column :width="100" title="Read" align="center">
            <template #cell="{ record }: { record: manga.Chapter }">
              <a-button size="mini" @click="$router.push(`/page/${record.id}`)"
                >Read</a-button
              >
            </template>
          </a-table-column>
        </template>
      </a-table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetMangaWithChapter } from '@wails/go/manga/Manga'
import { manga } from '@wails/go/models'
import { IconCheck, IconMinus } from '@arco-design/web-vue/es/icon'
import { MangaTitleURL, DateApp, GetBreakPoints } from '@/composable/helper'
import type { PaginationProps } from '@arco-design/web-vue'

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
})

const pprop = reactive<PaginationProps>({
  hideOnSinglePage: true,
})

const setAutopaginationSize = () => {
  if (bp.lg.value) {
    pprop.pageSize = 10
  } else {
    pprop.pageSize = 18
  }
}
setAutopaginationSize()
watch(
  () => bp.lg.value,
  () => {
    setAutopaginationSize()
  }
)
</script>

<style scoped></style>
