<template>
  <div class="flex justify-center items-center min-h-screen">
    <div class="bg-dark-400 rounded p-5">
      <a-space direction="vertical" size="large" :style="{ width: '600px' }">
        <div class="text-center text-lg text-orange-500">Manga Config</div>
        <a-input
          v-model="configForm.mangaPath"
          placeholder="Masukan manga folder path"
        />
        <div class="float-right">
          <a-button @click="saveConfig">Submit</a-button>
        </div>
      </a-space>
    </div>
    <div id="overlay" v-if="statusLoading">
      <div id="text" class="text-center">
        <div style="--un-bg-opacity: 0.9" class="bg-dark-50 rounded p-3">
          <div class="w-50vw lg:w-30vw">
            <span>Saving Chapters</span><br /><br />
            <a-spin
              style="--primary-6: var(--orange-4)"
              tip="This may take a while..."
            />
            <br />
            <div class="title">
              {{ title }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { SetConfig, AutoScanDirs } from '@wails/go/manga/Config'
import { EventsOn } from '@wails/runtime/runtime'
import { manga } from '@wails/go/models'
import { useMangaState } from '@/store'

const mangaState = useMangaState()

const emit = defineEmits<{
  (e: 'update:statusConfig', status: boolean): void
}>()

const configForm = reactive({
  mangaPath: '',
})

const statusLoading = ref(false)
const title = ref('')

const saveConfig = () => {
  let config = new manga.Config()
  config.manga_folder = configForm.mangaPath
  statusLoading.value = true
  SetConfig(config).then(() => {
    AutoScanDirs()
      .then(res => {
        statusLoading.value = false
        emit('update:statusConfig', true)
        mangaState.loadManga()
      })
      .catch(() => {
        emit('update:statusConfig', false)
      })

    EventsOn('status_scans', (mTitle: string) => {
      title.value = mTitle
    })
  })
}
</script>

<style scoped>
#overlay {
  position: fixed; /* Sit on top of the page content */
  display: block; /* Hidden by default */
  width: 100%; /* Full width (cover the whole page) */
  height: 100%; /* Full height (cover the whole page) */
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5); /* Black background with opacity */
  z-index: 2; /* Specify a stack order in case you're using a different order for other elements */
}
#text {
  position: absolute;
  top: 50%;
  left: 50%;
  color: white;
  transform: translate(-50%, -50%);
  -ms-transform: translate(-50%, -50%);
}

.title {
  line-height: 1.5;
  overflow: hidden;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
}
</style>
