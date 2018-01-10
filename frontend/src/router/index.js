import Vue from 'vue'
import Router from 'vue-router'
import BoilerplateSplash from '@/components/BoilerplateSplash'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'BoilerplateSplash',
      component: BoilerplateSplash
    }
  ]
})
