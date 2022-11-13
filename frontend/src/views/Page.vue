<template>
  <div>
    Page Vue
    <div>{{ cid }}</div>
    <div>
      {{ page }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetPage } from '@wails/go/manga/Manga'

// const Route = useRoute()
const props = defineProps<{
  cid: string
}>()

const page = ref()
GetPage(Number(props.cid)).then(res => {
  page.value = res
})
</script>

<style scoped></style>
