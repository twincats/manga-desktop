<template>
  <div class="text-left px-2">
    <a-typography>
      <a-typography-title :heading="5"
        >Chapter Manga <br /><small>
          {{ result?.title }}
        </small></a-typography-title
      >
      <a-typography-paragraph>
        Status Complete : {{ result?.status_end }} <br />
        Manga Alternative Title :
        <span v-if="result?.alter.length == 0">NO</span> <br />
        <ul v-if="result?.alter">
          <li v-for="(m, i) in result?.alter" :key="i">
            {{ m.title }}
          </li>
        </ul>
      </a-typography-paragraph>
    </a-typography>
    <div>
      <a-table :data="result?.chapter">
        <template #columns>
          <a-table-column title="Chapter ID" data-index="id"></a-table-column>
          <a-table-column title="Chapter" align="center">
            <template #cell="{ record }">
              <span class="text-orange-500">{{ record.chapter }}</span>
            </template>
          </a-table-column>
          <a-table-column
            title="Volume"
            data-index="volume"
            align="center"
          ></a-table-column>
          <a-table-column
            title="Group"
            data-index="group"
            align="center"
          ></a-table-column>
          <a-table-column
            title="Download Time"
            data-index="created_at"
            align="center"
          ></a-table-column>
          <a-table-column title="Bahasa" data-index="language">
            <template #cell="{ record }">
              {{ record.language.lang }}
            </template>
          </a-table-column>
          <a-table-column title="Status Read" align="center">
            <template #cell="{ record }">
              <span v-if="record.status_read" class="text-green-600"
                ><icon-check
              /></span>
              <span v-else class="text-red-600"><icon-minus /></span>
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

const props = defineProps<{
  mid: string
}>()

const result = ref<manga.Manga>()

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
</script>

<style scoped></style>
