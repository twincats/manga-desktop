<template>
  <div class="m-1" style="--primary-6: 255, 117, 24; --primary-5: 204, 92, 18">
    <a-tabs default-active-key="1">
      <a-tab-pane key="1" title="Download">
        <div class="h-[calc(100vh-7.5rem)] px-2">Content of Tab Panel 1 2</div>
      </a-tab-pane>
      <a-tab-pane key="2" title="Web">
        <div class="h-[calc(100vh-7.5rem)] px-2">
          <a-input-search
            allow-clear
            class="w-full mb-2"
            placeholder="URL"
            search-button
          />
          <div class="h-[calc(100vh-11.2rem)] overflow-auto p-2 bg-gray-800">
            <div id="webID"></div>
          </div>
        </div>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped></style>
