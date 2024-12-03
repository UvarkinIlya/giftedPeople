import { createApp } from 'vue'
import Persons from './views/PersonsView.vue'
import router from './router'

import './assets/person.css'

const app = createApp(Persons)

app.use(router)

app.mount('#app')
