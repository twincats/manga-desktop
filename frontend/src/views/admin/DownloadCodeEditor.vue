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
            v-model="TestResult"
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
import { MangaParser } from '@/composable/downloads/parser'
import { UpdateChapJS, UpdatePagesJS } from '@wails/go/manga/Manga'
import { Message, Modal } from '@arco-design/web-vue'
import '@arco-design/web-vue/es/message/style/index'
import '@arco-design/web-vue/es/modal/style/index'

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
const { servers, getSelectedServer, refreshServer } = UseServer()
const TestResult = ref('')
const selectedServerID = ref<number | string>('')
const urlTest = ref('')

// LOAD CHAPTER JS
const loadChapFunc = () => {
  if (
    typeof selectedServerID.value === 'number' &&
    selectedServerID.value > 1
  ) {
    const sv = getSelectedServer(selectedServerID.value)
    JSCode.value = sv?.chap_jscode!
  } else {
    Message.warning('Please Select Server Before LOAD ChapterJS')
  }
}

// LOAD PAGE JS
const loadPageFunc = () => {
  if (
    typeof selectedServerID.value === 'number' &&
    selectedServerID.value > 1
  ) {
    const sv = getSelectedServer(selectedServerID.value)
    JSCode.value = sv?.page_jscode!
  } else {
    Message.warning('Please Select Server Before LOAD PagesJS')
  }
}

// SAVE CHAPTER JS
const saveChapFunc = () => {
  if (
    typeof selectedServerID.value === 'number' &&
    selectedServerID.value > 1
  ) {
    if (JSCode.value != '') {
      UpdateChapJS(selectedServerID.value, JSCode.value)
        .then(stat => {
          if (stat) {
            Message.success('ChapterJS SAVED')
          }
          refreshServer()
        })
        .catch(e => {
          console.log('Error UpdateChapJS : ', e)
        })
    } else {
      Message.warning('Editor must contain valid JSCODE')
    }
  } else {
    Message.warning('Please Select Server and Entry Editor JSCODE')
  }
}

// SAVE PAGE JS
const savePageFunc = () => {
  if (
    typeof selectedServerID.value === 'number' &&
    selectedServerID.value > 1
  ) {
    if (JSCode.value != '') {
      UpdatePagesJS(selectedServerID.value, JSCode.value)
        .then(stat => {
          console.log('UpdatePagesJS : ', stat)
          if (stat) {
            Message.success('PagesJS SAVED')
          }
          refreshServer()
        })
        .catch(e => {
          console.log('Error UpdatePagesJS : ', e)
        })
    } else {
      Message.warning('Editor must contain valid JSCODE')
    }
  } else {
    Message.warning('Please Select Server and Entry Editor JSCODE')
  }
}

const mp = new MangaParser()
// TEST CHAPTER JS
const testChapFunc = () => {
  if (urlTest.value != '' && JSCode.value != '') {
    mp.fetch(urlTest.value)
      .then(out => {
        const func = mp.getFuncSync(JSCode.value, 'p')
        func(mp, mp.getParser(out))
          .then((returnVal: unknown) => {
            if (returnVal) {
              TestResult.value = JSON.stringify(returnVal, null, 4)
            }
          })
          .catch((e: unknown) => {
            TestResult.value = JSON.stringify(e, null, 4)
          })
          .finally(() => {
            if (!mp.isEmptyChapter()) {
              TestResult.value = JSON.stringify(mp.chapter, null, 4)
            }
          })
      })
      .catch(e => {
        Modal.error({
          title: 'Error Test ChapterJS',
          okText: 'OK',
          content: () => {
            return ['Error ChapterJS : ', h('br'), JSON.stringify(e)]
          },
        })
      })
  } else {
    Message.warning('Please fill Valid URL TEST and JSCode in Editor')
  }
}

// TEST PAGE JS
const testPageFunc = () => {
  if (urlTest.value != '' && JSCode.value != '') {
    mp.fetch(urlTest.value)
      .then(out => {
        const func = mp.getFuncSync(JSCode.value, 'p')
        func(mp, mp.getParser(out))
          .then((returnVal: unknown) => {
            if (returnVal) {
              TestResult.value = JSON.stringify(returnVal, null, 4)
            }
          })
          .catch((e: unknown) => {
            TestResult.value = JSON.stringify(e, null, 4)
          })
          .finally(() => {
            if (!mp.isEmptyPage()) {
              TestResult.value = JSON.stringify(mp.pages, null, 4)
            }
          })
      })
      .catch(e => {
        Modal.error({
          title: 'Error Test PagesJS',
          okText: 'OK',
          content: () => {
            return ['Error PagesJS : ', h('br'), JSON.stringify(e)]
          },
        })
      })
  } else {
    Message.warning('Please fill Valid URL TEST and JSCode in Editor')
  }
}

// CLEAR function
const clearFunc = () => {
  JSCode.value = ''
  selectedServerID.value = ''
  urlTest.value = ''
  TestResult.value = ''

  Message.info('Data is Cleared')
}
</script>

<style scoped></style>
