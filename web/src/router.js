import Vue from 'vue'
import VueRouter from 'vue-router'
import Join from './views/Join.vue'
import Intro from './views/Intro.vue'
import Experiment from './views/Experiment.vue'
import Farewell from './views/Farewell.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Join',
    component: Join,
  },
  {
    path: '/intro',
    name: 'Intro',
    component: Intro,
    meta: {
      auth: true,
    },
  },
  {
    path: '/experiment/:number',
    name: 'Experiment',
    component: Experiment,
    meta: {
      auth: true,
    },
  },
  {
    path: '/farewell',
    name: 'Farewell',
    component: Farewell,
  },
]

const router = new VueRouter({
  routes,
})

Vue.router = router

export default router
