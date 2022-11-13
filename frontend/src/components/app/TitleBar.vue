<template>
  <div id="titleBar" style="--wails-draggable: drag">
    <div class="left">
      <img class="w-5" src="../../assets/images/favicon.png" />
    </div>
    <div class="mid" @dblclick="toggleMaximize">AppName</div>
    <div class="right" style="--wails-draggable: no-drag">
      <div class="" onclick="window.runtime.WindowMinimise()">
        <i-mdi:window-minimize />
      </div>
      <div class="" @click="toggleMaximize">
        <i-mdi:window-maximize v-if="maxStatus" />
        <i-mdi:window-restore v-else />
      </div>
      <div class="close" onclick="window.runtime.Quit()">
        <i-mdi:close />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { WindowToggleMaximise } from '@wails/runtime'
const maxStatus = ref(true)

const toggleMaximize = () => {
  maxStatus.value = !maxStatus.value
  WindowToggleMaximise()
}
</script>

<style lang="less">
@title: rgb(15, 15, 15);

#titleBar {
  background-color: @title;
  display: flex;
  user-select: none;
  justify-content: center;
  align-items: center;
  height: 30px;

  .left {
    margin-left: 0.5rem;
  }
  .mid {
    width: 100%;
  }
  .right {
    display: flex;
    div {
      padding: 0.3rem 0.75rem;
      &:hover {
        background-color: lighten(@title, 10%);
      }
    }
    .close:hover {
      background-color: red !important;
    }
  }
}
</style>
