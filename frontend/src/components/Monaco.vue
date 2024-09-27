<script setup lang="ts">
import { monacoEditor, monacoLanguage } from '../libs/monaco'

const props = defineProps<{
  onChange?: (value: string) => any
  initialValue?: string
  config?: monacoEditor.IStandaloneEditorConstructionOptions
  modelValue?: string
  setLib?: string
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const container = ref<null | HTMLElement>(null)

// set lib global
if (props.setLib != undefined) {
  monacoLanguage.typescript.javascriptDefaults.setDiagnosticsOptions({
    noSemanticValidation: true,
    noSyntaxValidation: false,
  })

  monacoLanguage.typescript.javascriptDefaults.setCompilerOptions({
    target: monacoLanguage.typescript.ScriptTarget.ESNext,
    allowNonTsExtensions: true,
    allowJs: true, //!!!Require for js
  })

  monacoLanguage.typescript.javascriptDefaults.addExtraLib(
    props.setLib,
    'global.d.ts'
  )
}

onMounted(() => {
  nextTick(() => {
    const editor = monacoEditor.create(container.value!, {
      value: props.modelValue ? props.modelValue : props.initialValue,
      ...props.config,
    })
    editor.onDidChangeModelContent(() => {
      ;(props.onChange || (() => {}))(editor.getValue())
      editor.onDidBlurEditorText(() => {
        emit('update:modelValue', editor.getValue())
      })
    })
    watch(
      () => props.modelValue,
      newVAl => {
        if (newVAl != undefined) {
          editor.setValue(newVAl)
        }
      }
    )
  })
})
</script>

<template>
  <div id="container" ref="container" style="height: 100%"></div>
</template>

<!-- <style>
.monaco-editor .cursors-layer > .cursor {
	display: none !important;
}
</style> -->
