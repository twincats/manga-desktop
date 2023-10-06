<template>
  <div class="p-2">
    <div class="mb-2">Add Server</div>
    <!-- <div class="grid grid-cols-3 h-[calc(100vh-7rem)]">
      <div class="bg-red">
        <div>div 1-1</div>
        <div>div 1-2</div>
      </div>
      <div class="bg-green">div 2</div>
      <div class="bg-blue">div 3</div>
    </div> -->

    <div class="grid grid-cols-6 h-[calc(100vh-6rem)]">
      <a-split class="col-span-4" :min="0.5" :max="0.7">
        <template #second>
          <div class="bg-green mx-1 h-full">Right</div>
        </template>
        <template #first>
          <div>
            <a-split
              :disabled="true"
              direction="vertical"
              :min="0.3"
              :max="0.7"
              class="h-[calc(100vh-6.25rem)]"
            >
              <template #first>
                <div class="bg-red mx-1 h-full">TOP</div>
              </template>
              <template #second
                ><div class="bg-blue mx-1 h-full">Bottom</div></template
              >
            </a-split>
          </div>
        </template>
      </a-split>
      <div class="bg-purple mx-1 h-full col-span-2">div 3</div>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped></style>
