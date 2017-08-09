import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Match from '@/components/Match'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/match/:id',
      name: 'Match',
      component: Match
    }
  ]
})
