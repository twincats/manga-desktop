<template>
  <div class="p-2 text-center">
    <div class="">Add Alter Manga Title</div>
    <div>{{ mangaData?.title }}</div>
    <div class="w-[70%] mx-auto">
      <a-input-search
        class="my-2"
        placeholder="Please enter something"
        :input-attrs="{ class: 'text-center', spellcheck: 'false' }"
        search-button
        @search="searchfunc"
      >
        <template #button-icon>
          <icon-save />
        </template>
      </a-input-search>
      <a-list v-if="mangaData" :data="alters" size="small" hoverable>
        <!-- <a-list-item
          v-if="mangaData.alter.length == 0"
          style="padding: 0.25rem 0.5rem"
          >No Alternative</a-list-item
        > -->
        <template #empty>
          <div class="select-none" style="padding: 0.25rem 0.5rem">
            No Alternative
          </div>
        </template>
        <template #item="{ item }: { item: manga.Alter }">
          <a-list-item class="list-demo-item" action-layout="vertical">
            {{ item.title }}
          </a-list-item>
        </template>
      </a-list>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetMangaWithAlter } from '@wails/go/manga/Manga'
import { manga } from '@wails/go/models'
import { IconSave } from '@arco-design/web-vue/es/icon'
import { Message } from '@arco-design/web-vue'
import '@arco-design/web-vue/es/message/style/index'
import { SetAlter } from '@wails/go/manga/Alter'

interface ConvertProps {
  mid?: number | null
}
const { mid } = withDefaults(defineProps<ConvertProps>(), { mid: null })
const mangaData = ref<manga.Manga | null>(null)
const alters = ref<manga.Alter[]>([])

if (mid) {
  GetMangaWithAlter(mid).then(res => {
    mangaData.value = res
    alters.value = res.alter
  })
}

const searchfunc = (val: string, _: MouseEvent) => {
  if (val.toLowerCase() != mangaData.value?.title.toLowerCase() && val != '') {
    let statPush = false
    if (alters.value.length > 0) {
      alters.value.forEach(item => {
        if (val.toLowerCase() != item.title.toLowerCase()) {
          statPush = true
        }
      })
    } else {
      statPush = true
    }
    if (statPush) {
      const a = new manga.Alter()
      a.title = val
      a.manga_id = mid!
      SetAlter(a)
        .then(() => {
          alters.value.push(a)
        })
        .catch(er => {
          Message.error(er)
        })
    } else {
      //error title already in alternatives
      Message.error('Manga title Alternative already in DB')
    }
  } else {
    //error val sama title / val empty
    Message.error('Manga title Alternative is empty or Same as Manga Title')
  }
}
</script>

<style scoped></style>
