<template>
  <div id="App">
    <div id="nav">
      <router-link v-if="authenticated" to="/login" v-on:click.native="logout()" replace>Выйти из учётной записи
      </router-link>
    </div>
    <router-view @authenticated="setAuthenticated"/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      authenticated: false,
      mockAccount: {
        username: 'hse',
        password: 'hsepass'
      }
    }
  },
  mounted () {
    if (!this.authenticated) {
      this.$router.replace({name: 'HelloWorld'})
    }
  },
  methods: {
    setAuthenticated (status) {
      this.authenticated = status
    },
    logout () {
      this.authenticated = false
    }
  }
}
</script>

<style>
  #App {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
