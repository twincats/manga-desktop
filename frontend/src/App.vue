<template>
  <div>
    <a-config-provider :locale="enUS">
      <div
        v-if="statusConfig"
        style="
          --primary-6: 255, 117, 24;
          --color-primary-light-1: rgba(var(--primary-6), 0.2);
          --color-primary-light-3: rgba(var(--primary-6), 0.5);
          --primary-5: 255, 139, 61;
          --primary-7: 204, 92, 18;
        "
      >
        <menu-bar
          @clickMenu="collapse = !collapse"
          @clickToggleTheme="toggleTheme()"
        />
        <side-bar :collapse="collapse" />
        <div id="main">
          <router-view></router-view>
        </div>
        <footer-bar />
      </div>
      <manga-path @update:status-config="SetStatusConfig" v-else />
    </a-config-provider>
  </div>
</template>

<script lang="ts" setup>
import enUS from '@arco-design/web-vue/es/locale/lang/en-us'
import { EmitListenOnce, useTheme } from '@/composable/helper'
import { GetConfig } from '@wails/go/manga/Config'
import { useMangaState } from '@/store'

const mangaState = useMangaState()

const statusConfig = ref(false)
GetConfig().then(res => {
  if (res.manga_folder != '') {
    SetStatusConfig(true)
  }
})

const SetStatusConfig = (s: boolean) => {
  statusConfig.value = s
  mangaState.loadManga()
}

const router = useRouter()
const collapse = ref(true)
const { theme, toggleTheme } = useTheme()
EmitListenOnce('args', (res: string[]) => {
  if (res.length) {
    if (res.includes('convert')) {
      router.push('/convert')
    }
  }
})
</script>

<style>
#main {
  height: calc(100vh - 56px);
  overflow: auto;
  background-color: rgb(v-bind('theme.bg'));
  color-scheme: v-bind('theme.scheme');
}
</style>
