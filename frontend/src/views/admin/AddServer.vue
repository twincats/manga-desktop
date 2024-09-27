<template>
  <div class="p-2">
    <!-- <div class="grid grid-cols-3 h-[calc(100vh-7rem)]">
      <div class="bg-red">
        <div>div 1-1</div>
        <div>div 1-2</div>
      </div>
      <div class="bg-green">div 2</div>
      <div class="bg-blue">div 3</div>
    </div> -->

    <div class="grid grid-cols-2 h-[calc(100vh-72px)]">
      <!-- <div class="bg-purple mx-1 h-full col-span-2">div 3</div> -->
      <div class="grid grid-rows-2 grid-cols-1">
        <div class="h-[calc(50vh-50px-0.25rem)]">
          <div class="mb-1">ChapterJS</div>
          <Monaco
            :config="{
              language: 'javascript',
              theme: 'vs-dark',
              automaticLayout: true,
              minimap: { enabled: false },
              overviewRulerLanes: 0,
            }"
            :set-lib="lib"
          />
        </div>
        <div class="h-[calc(50vh-52px-0.45rem)]">
          <div class="my-1">PageJS</div>
          <Monaco
            :config="{
              language: 'javascript',
              theme: 'vs-dark',
              automaticLayout: true,
              minimap: { enabled: false },
              overviewRulerLanes: 0,
            }"
            :set-lib="lib"
          />
        </div>
      </div>
      <div class="flex flex-col lg:flex-row-reverse">
        <div class="h-[calc(50vh-35px)] pl-2 w-[calc(50vw-40px)] lg:h-full">
          <div class="mb-1">Entry Add SERVER</div>
          <div>
            <a-input class="mb-1" placeholder="Servername" allow-clear />
            <a-input class="mb-1" placeholder="Server URL" allow-clear />
            <div class="mb-1 grid grid-cols-2">
              <div>Active <a-switch /></div>
              <div>Search <a-switch /></div>
            </div>

            <a-input class="mb-1" placeholder="TEST URL" allow-clear />
            <a-button-group class="mb-1 w-full">
              <a-button long>TEST ChapterJS</a-button>
              <a-button long> TEST PageJS </a-button>
            </a-button-group>
            <a-button class="mb-1" long> CLEAR </a-button>
            <a-button class="mb-2" long> SAVE SERVER </a-button>
            <div
              class="border-1 border-solid border-orange-400 rounded-sm h-13 flex justify-center items-center"
            >
              <strong>
                <span class="text-green-500">"SUKSES TEST"</span>
                ChapterJS</strong
              >
            </div>
          </div>
        </div>
        <div
          class="h-[calc(50vh-55px-0.25rem)] lg:(h-[calc(100vh-91px)] w-[50vw])"
        >
          <div class="my-1 lg:mt-0">Test Result</div>
          <Monaco
            :config="{
              language: 'json',
              theme: 'vs-dark',
              automaticLayout: true,
              minimap: { enabled: false },
              overviewRulerLanes: 0,
              lineNumbers: 'off',
              matchBrackets: 'never',
              renderLineHighlight: 'none',
              readOnly: true,
            }"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import models from '@wails/go/models.ts?raw'
import parser from '@/composable/downloads/parser.ts?raw'
const lib = [
  models,
  parser,
  `const C = new MangaParser();const p = new Parser(C);`,
]
  .join('\n')
  .replace(/export type/g, 'type')
  .replace(/export class/g, 'class')
  .replace(/export namespace/g, 'namespace')
  .replace(`import { types } from '@wails/go/models'`, '')
  .replace(`import { Fetch, FetchPost } from '@wails/go/tool/Web'`, '')
</script>

<style scoped></style>
