<template>
  <div>
    <div v-if="page" v-for="(p, i) in page.pages" class="flex justify-center">
      <img class="" :key="i" :src="`file/${page.path}/${p}`" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetPage } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'

// const Route = useRoute()
const props = defineProps<{
  cid: string
}>()

const page = ref<manga.Page>()
GetPage(Number(props.cid)).then(res => {
  page.value = res
})
</script>

<style scoped></style>
