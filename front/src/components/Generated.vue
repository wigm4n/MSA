<template>
  <div id="Generated">
    <h1>Задание сгенерированно</h1>
    <button type="button" v-on:click="downloadArchive">Скачать архив сгенерированных выборок</button>
    <button type="button" v-on:click="downloadEach">Скачать расчётные показатели по каждый выборке</button>
    <button type="button" v-on:click="remove">Удалить</button>
  </div>
</template>

<script>
export default {
  name: 'Generated',
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
    downloadArchive () {

    },
    downloadEach () {

    },
    remove () {

    }
  }
}
</script>

<style scoped>
  #Generated {
    width: 500px;
    border: 1px solid #CCCCCC;
    background-color: #FFFFFF;
    margin: auto;
    padding: 20px;
    display: flex;
    flex-direction: column;
  }
</style>
