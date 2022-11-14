<template>
  <div class="menu-bar">
    <div class="w-full flex justify-start">
      <a-button-group size="mini">
        <a-button @click="emit('clickMenu')"> <icon-list /> </a-button>
        <a-button @click="$router.push('/')"> M </a-button>
        <a-button @click="$router.push('/convert')"> CV </a-button>
        <a-button> <icon-download /> </a-button>
      </a-button-group>
    </div>
    <div class="w-full flex justify-end">
      <a-button-group size="mini">
        <a-button @click="$router.push('/home')"> H </a-button>
        <a-button @click="emit('clickToggleTheme')">
          <icon-moon-fill v-if="mode == 'dark'" />
          <icon-sun v-else />
        </a-button>
        <a-button> <icon-settings /> </a-button>
      </a-button-group>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useTheme } from '@/composable/helper'
import {
  IconList,
  IconMoonFill,
  IconSun,
  IconDownload,
  IconSettings,
} from '@arco-design/web-vue/es/icon'

const { mode } = useTheme()

const emit = defineEmits(['clickMenu', 'clickToggleTheme'])
</script>

<style lang="less" scoped>
.menu-demos {
  box-sizing: border-box;
  width: 100%;
  padding: 40px;
  background-color: var(--color-neutral-2);
}

.menu-bar {
  height: 28px;
  background-color: var(--color-menu-light-bg);
  display: flex;
  align-items: center;
}
</style>
