import { createApp } from 'vue'
import App from './App.vue'
import Router from './router/router'

import 'uno.css'
import './assets/style/app.less'

//create app
const app = createApp(App)
app.use(Router)
app.mount('#app')

document.body.setAttribute('arco-theme', 'dark')
