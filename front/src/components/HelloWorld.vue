<template>
  <div>
    <img @click="clickFunc" src="../assets/1369096349_shkola-sber.jpg">
    <button v-on:click="navigateProfessor">Я преподаватель</button>
    <button v-on:click="navigateStudent">Я студент</button>
  </div>
</template>

<script>
import router from '../router'

export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  methods: {
    navigateStudent () {
      router.push({name: 'TODOPage'})
    },
    navigateProfessor () {
      router.push({name: 'Login'})
    },
    clickFunc () {
      window.open('https://www.vladtime.ru/econom/15547', '_blank')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
