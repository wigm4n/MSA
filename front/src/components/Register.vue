<template>
  <div id="Register">
    <h1>Введите данные о новом преподавателе</h1>
    <h2>Логин и пароль будут отправлены преподавателю на почту</h2>
    <input type="text" name="email" v-model="input.email" placeholder="Почтовый ящик"/>
    <input type="text" name="name" v-model="input.name" placeholder="Имя"/>
    <input type="text" name="surname" v-model="input.surname" placeholder="Фамилия"/>
    <input type="text" name="lastname" v-model="input.lastname" placeholder="Отчество"/>
    <button type="button" v-on:click="sendData">Создать преподавателя</button>
  </div>
</template>

<script>
export default {
  name: 'Register',
  data () {
    return {
      input: {
        email: '',
        name: '',
        surname: '',
        lastname: ''
      }
    }
  },
  methods: {
    sendData () {
      return {
        registerInput: {
          email: this.input.email,
          name: this.input.name,
          surname: this.input.surname,
          lastname: this.input.lastname
        }
      }
    }
  }
}
</script>

<style scoped>
  #Register {
    width: 500px;
    border: 1px solid #CCCCCC;
    background-color: #FFFFFF;
    margin: auto;
    padding: 20px;
    display: flex;
    flex-direction: column;
  }
</style>
