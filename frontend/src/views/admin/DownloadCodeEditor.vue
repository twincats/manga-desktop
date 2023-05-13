<template>
  <div class="h-full">
    <div class="py-2 text-center bg-[#1E1E1E] select-none">
      Download JSCode EDITOR
    </div>
    <div class="grid grid-cols-2 mt-3">
      <div style="height: calc(100vh - 100px)">
        <Monaco
          :config="{
            language: 'javascript',
            theme: 'vs-dark',
            automaticLayout: true,
            minimap: { enabled: false },
            overviewRulerLanes: 0,
          }"
          v-model="JSCode"
          :set-lib="lib"
        />
      </div>
      <div class="px-3">
        <div>
          <a-space direction="vertical" fill>
            <a-select placeholder="Server List" v-model="selectedServerID">
              <a-option v-for="(i, x) in servers" :key="x" :value="i.id">{{
                i.name
              }}</a-option>
            </a-select>
            <div class="grid grid-cols-2 gap-2">
              <a-button @click="loadChapFunc">Load ChapterJS</a-button>
              <a-button @click="loadPageFunc">Load PagesJS</a-button>
            </div>

            <div class="grid grid-cols-2 gap-2">
              <a-button @click="saveChapFunc" long status="success"
                >SAVE ChapterJS</a-button
              >
              <a-button @click="savePageFunc" long status="success"
                >SAVE PagesJS</a-button
              >
            </div>
            <div class="grid grid-cols-2 gap-2">
              <a-button @click="testChapFunc" long status="warning"
                >TEST ChapterJS</a-button
              >
              <a-button @click="testPageFunc" long status="warning"
                >TEST PagesJS</a-button
              >
            </div>
            <a-button @click="clearFunc" long>Clear</a-button>
            <a-input v-model="urlTest" placeholder="URL TEST" allow-clear />
          </a-space>
        </div>
        <div
          class="p-3 mt-2 bg-[#1E1E1E] h-[calc(100vh-364px)] overflow-y-auto"
        >
          <!-- <div class="whitespace-pre">
            {{ selectedServerID }}
            {{ JSON.stringify(TestResult, null, 4) }}
          </div> -->
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
            v-model="jsonCode"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import models from '@wails/go/models.ts?raw'
import parser from '@/composable/downloads/parser.ts?raw'
import { UseServer } from '@/composable/downloads/download'

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

const JSCode = ref('')
const { servers } = UseServer()
const TestResult = ref()
const selectedServerID = ref<number | string>('')
const urlTest = ref('')

const loadChapFunc = () => {}
const loadPageFunc = () => {}
const saveChapFunc = () => {}
const savePageFunc = () => {}
const testChapFunc = () => {}
const testPageFunc = () => {
  console.log(JSCode.value)
}
const clearFunc = () => {
  // console.log(jsonCode.value)
  JSCode.value = ''
  selectedServerID.value = ''
  urlTest.value = ''
  TestResult.value = ''
  jsonCode.value = ''
}

const jsonCode = ref('')
</script>

<style scoped></style>
