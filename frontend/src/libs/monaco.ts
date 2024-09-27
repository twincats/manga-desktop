import 'monaco-editor/esm/vs/editor/editor.all.js'

// language feature
import 'monaco-editor/esm/vs/language/typescript/monaco.contribution'
import 'monaco-editor/esm/vs/language/css/monaco.contribution'
import 'monaco-editor/esm/vs/language/json/monaco.contribution'
import 'monaco-editor/esm/vs/language/html/monaco.contribution'

// basic language syntax
import 'monaco-editor/esm/vs/basic-languages/typescript/typescript.contribution'
import 'monaco-editor/esm/vs/basic-languages/javascript/javascript.contribution'
import 'monaco-editor/esm/vs/basic-languages/css/css.contribution'
import 'monaco-editor/esm/vs/basic-languages/html/html.contribution'
import 'monaco-editor/esm/vs/basic-languages/sql/sql.contribution'

import {
  editor as monacoEditor,
  languages as monacoLanguage,
  Uri as MonacoURI,
} from 'monaco-editor/esm/vs/editor/editor.api'

export { monacoEditor, monacoLanguage, MonacoURI }
