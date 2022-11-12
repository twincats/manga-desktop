import { createApp } from 'vue'
import App from './App.vue'
import Router from './router/router'

import 'uno.css'

//create app
const app = createApp(App)
app.use(Router)
app.mount('#app')
