<template>
  <div id="PersonalArea">
    <h1>Личный кабинет</h1>
    <div id="Buttons">
      <button id="btn1" type="button" v-on:click="createMathTask">Создать задание</button>
      <button type="button" v-on:click="showMathTasks">Созданные задания</button>
      <button type="button" v-on:click="showForum">Форум</button>
      <button type="button" v-on:click="registerNewProff">Зарегистрировать нового преподавателя</button>
      <sidebar-menu :menu="menu" />
    </div>
  </div>
</template>

<script>
import router from '../router'

export default {
  name: 'PersonalArea',
  data () {
    return {
      menu: [
        {
          header: true,
          title: 'Main Navigation'
        },
        {
          href: '/',
          title: 'Dashboard',
          icon: 'fa fa-user'
        },
        {
          title: 'Charts',
          icon: 'fa fa-chart-area',
          child: [
            {
              href: '/charts/sublink',
              title: 'Sub Link'
            }
          ]
        }
      ]
    }
  },
  methods: {
    showForum () {
      router.push({name: 'TODOPage'})
    },
    createMathTask () {
      router.push({name: 'Create'})
    },
    showMathTasks () {
      router.push({name: 'TODOPage'})
    },
    registerNewProff () {
      router.push({name: 'Register'})
    }
  }
}
</script>

<style scoped>
  #PersonalArea {
    background-color: #FFFFFF;
    border: 1px solid #CCCCCC;
    padding: 20px;
    margin-top: 10px;
  }

  #Buttons {
    display: flex;
    align-items: center;
    flex-direction: column;
  }
</style>
