import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import './setup/axious'

import './setup/sentry'

Vue.config.productionTip = false
Vue.prototype.$demoMode = process.env.VUE_APP_DEMO_MODE === 'true'

if (!Vue.prototype.$demoMode) {
  require('./setup/auth')
}

new Vue({
  router,
  vuetify,
  render: h => h(App),
}).$mount('#app')
