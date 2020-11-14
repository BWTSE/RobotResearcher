import Vue from 'vue'
import Auth from '@websanova/vue-auth'
import AuthBasic from '@websanova/vue-auth/drivers/auth/basic.js'
import AuthAxios from '@websanova/vue-auth/drivers/http/axios.1.x.js'
import AuthRouter from '@websanova/vue-auth/drivers/router/vue-router.2.x.js'

Vue.use(Auth, {
  auth: AuthBasic,
  http: AuthAxios,
  router: AuthRouter,
  refreshData: {
    enabled: false,
  },
  fetchData: {
    enabled: true,
    url: 'auth/status',
  },
  loginData: {
    url: 'auth/register',
  },
  authRedirect: '/',
})
