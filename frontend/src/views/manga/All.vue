<template>
  <div>manga all</div>
  <div>
    <!-- <div :style="{ minHeight: minH }">
      <div class="grid grid-cols-5 gap-2 lg:grid-cols-8">
        <div
          v-for="(manga, i) in mangaHome?.manga"
          :key="i"
          class="box relative"
          :class="today(manga.download_time)"
          @contextmenu.prevent="openContextMenu($event, manga)"
        >
          <div class="cover" @click="$router.push(`chapter/${manga.id}`)">
            <div class="h-203px lg:h-200px">
              <img
                :alt="manga.title"
                :src="`file/${MangaTitleURL(manga.title)}/cover.webp`"
                :onerror="errorLoadImage"
              />
            </div>
          </div>
          <div
            class="title select-none"
            @click="$router.push(`chapter/${manga.id}`)"
          >
            <span>{{ manga.title }}</span>
          </div>
          <div class="chapter">
            <button @click.stop="$router.push(`/page/${manga.chapter_id}`)">
              Chapter {{ manga.chapter }}
            </button>
          </div>
          <div
            v-if="
              DateApp.NewDate().Format('DD-MM-YYYY') ==
              DateApp.NewDate(manga.download_time.toString()).Format(
                'DD-MM-YYYY'
              )
            "
            class="absolute text-xs top-2 right-0 px-1 rounded-l border-b border-dark-100 drop-shadow bg-blue-500"
          >
            New
          </div>
        </div>
      </div>
    </div>
    <div class="mt-2">
      <a-pagination
        class="justify-center"
        v-model:current="nav.page"
        v-model:page-size="nav.limit"
        :hide-on-single-page="true"
        :auto-adjust="false"
        :total="mangaHome?.pagination?.total ? mangaHome?.pagination?.total : 0"
      />
    </div> -->
    {{ mangaHome?.manga }}
    <br />
    total : {{ mangaHome?.manga.length }}
  </div>
</template>

<script setup lang="ts">
import { GetMangaHome } from '@wails/go/manga/Manga'
import { manga } from '@wails/go/models'

const mangaHome = ref<manga.MangaHome | null>(null)
const loadManga = async () => {
  const start = Date.now()
  mangaHome.value = await GetMangaHome(null, null, 1, 0)
  const end = Date.now()
  console.log(`Execution time: ${end - start} ms`)
}
loadManga()
</script>

<style scoped></style>
