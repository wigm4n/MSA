<template>
  <div id="Create">
    <div id="Context">
      <div id="TaskName">
        <p>Укажите название задания</p>
        <input type="text" v-model="createInput.TaskName">
      </div>
      <div id="TaskType">
        <p>Выберите тип задачи</p>
        <select v-model="createInput.TaskType">
          <option>Тип №1</option>
          <option>Тип №2</option>
          <option>Тип №3</option>
          <option>Тип №4</option>
          <option>Тип №5</option>
          <option>Тип №6</option>
        </select>
      </div>
      <div id="Size">
        <p>Укажите длину выборки</p>
        <input type="number" v-model="createInput.Size" min="1" oninput="validity.valid||(value='');">
      </div>
      <div id="HowMuch">
        <p>Укажите количество выборок</p>
        <input type="number" v-model="createInput.HowMuch" min="1" oninput="validity.valid||(value='');">
      </div>
      <div id="Min">
        <p>Минимальное значение выборки</p>
        <input type="number" v-model="createInput.Min" min="1" oninput="validity.valid||(value='');">
      </div>
      <div id="Max">
        <p>Максимальное значение выборки</p>
        <input type="number" min="1" v-model="createInput.Max" oninput="validity.valid||(value='');">
      </div>
      <button type="button" v-on:click="generate">Сгенерировать и отправить студентам</button>
    </div>
  </div>
</template>

<script>
import JQuery from 'jquery'
import router from '../router'

let $ = JQuery

export default {
  name: 'Create',
  data () {
    return {
      createInput: {
        TaskName: '',
        TaskType: '',
        Size: 1,
        HowMuch: 1,
        Min: 1,
        Max: 1
      }
    }
  },
  methods: {
    generate () {
      var self = this

      // TODO sync
      $.ajax({
        url: '/test',
        type: 'POST',
        data: {
          'TaskName': this.createInput.TaskName,
          'TaskType': this.createInput.TaskType,
          'Size': this.createInput.Size,
          'HowMuch': this.createInput.HowMuch,
          'Min': this.createInput.Min,
          'Max': this.createInput.Max
        },
        success: function (data) {
          console.log('Прибыли данные: ' + data)
          self.$emit('close', data)
        },
        error: function (xhr, status, error) {
          console.log('Ошибка сохранения: ' + error)

          console.log(error)
          self.showDismissibleAlert = true

          console.log('Data send')
        }
      })
      router.push({name: 'Generated'})
      console.log('Router pushed')
    }
  }
}
</script>

<style scoped>
  #Create {
    background-color: #FFFFFF;
    border: 1px solid #CCCCCC;
    padding: 20px;
    margin-top: 10px;
  }

  #Context {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  #TaskName {
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  #TaskType {
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  #Size {
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  #HowMuch {
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  #Min {
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  #Max {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
</style>
