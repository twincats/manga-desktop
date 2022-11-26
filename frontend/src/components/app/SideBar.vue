<template>
  <div class="menu-demo">
    <a-menu
      :style="{ width: '200px', height: '100%' }"
      :default-open-keys="['0']"
      :default-selected-keys="['0_2']"
      :collapsed="$props.collapse"
    >
      <a-sub-menu key="0">
        <template #icon><icon-apps></icon-apps></template>
        <template #title>Navigation 1</template>
        <a-menu-item key="0_0">Home</a-menu-item>
        <a-menu-item key="0_1">Chapter</a-menu-item>
        <a-menu-item key="0_2">Download</a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="1">
        <template #icon><icon-bug></icon-bug></template>
        <template #title>Navigation 2</template>
        <a-menu-item key="1_0">Menu 1</a-menu-item>
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
      <a-menu-item key="3">
        <template #icon><icon-settings /></template>
        Menu 1</a-menu-item
      >
    </a-menu>
  </div>
</template>

<script lang="ts" setup>
import {
  IconApps,
  IconBug,
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

<style scoped>
.menu-demo {
  /* box-sizing: border-box; */
  float: left;
  height: calc(100vh - 28px);
  padding: 0;
  margin: 0;
  background-color: var(--color-neutral-2);
}
</style>
