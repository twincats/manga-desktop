import { createApp } from 'vue'
import { monacoLoader } from './libs/monacoLoader'
import App from './App.vue'
import router from './router/router'

import 'virtual:uno.css'
import './assets/style/app.less'

//create app
const app = createApp(App)
app.use(router)
app.mount('#app')

// load monaco worker
monacoLoader()

//set default dark theme
document.body.setAttribute('arco-theme', 'dark')
