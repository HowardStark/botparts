import Vue from 'vue'
import Router from 'vue-router'
import Menu from '@/components/Menu'
import Splash from '@/components/Splash'
import Systems from '@/components/Systems'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/menu',
      name: 'Menu',
      component: Menu
    },
    {
      path: '/',
      name: 'Splash',
      component: Splash
    },
    {
      path: '/systems',
      name: 'Systems',
      component: Systems
    }
  ]
})
