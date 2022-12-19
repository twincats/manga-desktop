<template>
  <div>
    <div
      class="flex justify-center flex-wrap flex-row-reverse lg:overflow-hidden"
    >
      <div v-if="page" v-for="(p, i) in page.pages" :key="i">
        <img
          class="max-w-full lg:ml-1 mb-1 select-none"
          draggable="false"
          :src="`file/${page.path}/${p}`"
          @contextmenu.prevent="openContextMenu"
        />
      </div>
    </div>

    <div v-if="page">
      <page-nav @change="updatePageChange" :chaps="page.chaps"></page-nav>
    </div>
    <teleport to="#main">
      <context-menu ref="refMenu" :width="200">
        <li>Image URL</li>
        <li @click="$router.push('/chapter/' + page?.manga_id)">
          Home Chapter
        </li>
        <div class="divider"></div>
        <li>Delete Image</li>
      </context-menu>
    </teleport>
  </div>
</template>

<script setup lang="ts">
import { GetPage } from '@wails/go/manga/Manga'
import type { manga } from '@wails/go/models'
import { UseContextMenu } from '@/composable/helper'

// const Route = useRoute()
const props = defineProps<{
  cid: string
}>()

const { refMenu, openContextMenu } = UseContextMenu()
const router = useRouter()
const updatePageChange = (chapId: number) => {
  console.log('event change :', chapId)
  router.push({ name: 'page', params: { cid: chapId } })
}

//main data page
const page = ref<manga.Page | null>(null)

//fetching function
const fetchData = async (cid: number) => {
  page.value = await GetPage(cid)
}
//initial fetch Data
fetchData(Number(props.cid))

const el = document.getElementById('main')
const scroll = useScroll(el, { behavior: 'smooth' })
//watch page params
watch(
  () => Number(props.cid),
  async (newcid, _) => {
    await fetchData(newcid)
    scroll.y.value = 0
    // if (newcid > oldcid) {
    // scroll.y.value = 0
    // } else {
    //   scroll.y.value = 1000 * 1000
    // }
  }
)

//mouse event
// window.addEventListener('mousedown', function handleMouseButtonDown(event) {
//   if (event.button === 0) {
//     // left mouse button
//     console.log('button 0')
//   } else if (event.button === 1) {
//     // middle mouse button
//     console.log('button 1')
//   } else if (event.button === 2) {
//     // right mouse button
//     console.log('button 2')
//   } else if (event.button === 3) {
//     // back mouse button
//     console.log('button 3')
//     // event.preventDefault()
//   } else if (event.button === 4) {
//     // forward mouse button
//     console.log('button 4')
//     // event.preventDefault()
//   } else {
//     // other mouse button
//     console.log('button tak tahu')
//   }
// })
</script>

<style scoped>
#ppnav {
  position: absolute;
  bottom: 135px;
  width: 165px;
  background-color: var(--color-bg-1);
  left: 50%;
  transform: translateX(-50%);
  opacity: 0.9;
  border-radius: 0.5rem;
}
</style>
