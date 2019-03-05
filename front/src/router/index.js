import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '../components/HelloWorld'
import Login from '../components/Login'
import TODOPage from '../components/TODOPage'
import PersonalArea from '../components/PersonalArea'
import Register from '../components/Register'
import Create from '../components/Create'
import Generated from '../components/Generated'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/Login',
      name: 'Login',
      component: Login
    },
    {
      path: '/PersonalArea',
      name: 'PersonalArea',
      component: PersonalArea
    },
    {
      path: '/TODOPage',
      name: 'TODOPage',
      component: TODOPage
    },
    {
      path: '/PersonalArea/Register',
      name: 'Register',
      component: Register
    },
    {
      path: '/PersonalArea/Create',
      name: 'Create',
      component: Create
    },
    {
      path: 'PersonalArea/Create/Generated',
      name: 'Generated',
      component: Generated
    }
  ]
})
