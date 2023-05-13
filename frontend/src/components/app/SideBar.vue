<template>
  <div class="menu-sidebar">
    <a-menu
      :style="{ width: '200px', height: '100%' }"
      :default-open-keys="['0']"
      :default-selected-keys="['0_2']"
      :collapsed="$props.collapse"
    >
      <a-sub-menu key="0">
        <template #icon><icon-apps></icon-apps></template>
        <template #title>Manga</template>
        <a-menu-item key="0_0" @click="$router.push('/')">Home</a-menu-item>
        <a-menu-item key="0_1">Chapter</a-menu-item>
        <a-menu-item key="0_2" @click="$router.push('/download')"
          >Download</a-menu-item
        >
      </a-sub-menu>
      <a-sub-menu key="1">
        <template #icon><icon-desktop /></template>
        <template #title>Admin</template>
        <a-menu-item key="1_0" @click="$router.push('/dlcode')"
          >Download JSCode</a-menu-item
        >
        <a-menu-item key="1_1">Menu 2</a-menu-item>
        <a-menu-item key="1_2">Menu 3</a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="2">
        <template #icon><icon-bulb></icon-bulb></template>
        <template #title>Navigation 3</template>
        <a-menu-item key="2_0">Menu 1</a-menu-item>
        <a-menu-item key="2_1">Menu 2</a-menu-item>
        <a-sub-menu key="2_2" title="Navigation 4">
          <a-menu-item key="2_2_0">Menu 1</a-menu-item>
          <a-menu-item key="2_2_1">Menu 2</a-menu-item>
        </a-sub-menu>
      </a-sub-menu>
      <a-sub-menu key="3">
        <template #icon><icon-settings /></template>
        <template #title>Setting</template>
        <a-menu-item key="3_1">Theme</a-menu-item>
        <a-menu-item key="3_1">Download</a-menu-item>
      </a-sub-menu>
    </a-menu>
  </div>
</template>

<script lang="ts" setup>
import {
  IconApps,
  IconDesktop,
  IconBulb,
  IconSettings,
} from '@arco-design/web-vue/es/icon'

interface Props {
  collapse?: boolean
}

//props with default
withDefaults(defineProps<Props>(), {
  collapse: true,
})
</script>

<style>
.menu-sidebar {
  /* box-sizing: border-box; */
  float: left;
  height: calc(100vh - 28px);
  padding: 0;
  margin: 0;
  background-color: var(--color-neutral-2);
}
.arco-menu-item {
  line-height: 30px !important;
}

.arco-trigger-menu-item {
  height: 26px;
  line-height: 26px;
  font-size: 12px;
}
</style>
