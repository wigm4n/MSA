<template>
    <div id="Login">
        <h1>Личный кабинет преподавателя</h1>
        <input type="text" name="username" v-model="input.username" placeholder="Логин" />
        <input type="password" name="password" v-model="input.password" placeholder="Пароль" />
        <button type="button" v-on:click="login()">Войти</button>
        <router-link to="/TODOPage" replace>Забыли пароль?</router-link>
    </div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      input: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    login () {
      if (this.input.username != '' && this.input.password != '') {
        if (this.input.username == this.$parent.mockAccount.username && this.input.password == this.$parent.mockAccount.password) {
          this.$emit('authenticated', true)
          this.$router.replace({name: 'PersonalArea'})
        } else {
          alert('The username and / or password is incorrect')
        }
      } else {
        alert('A username and password must be present')
      }
    }
  }
}
</script>

<style scoped>
    #Login {
        width: 500px;
        border: 1px solid #CCCCCC;
        background-color: #FFFFFF;
        margin: auto;
        margin-top: 200px;
        padding: 20px;
        display: flex;
        flex-direction: column;
    }
</style>
