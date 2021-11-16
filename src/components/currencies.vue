<template>
  <h1>Валюты</h1>
  <table>
    <tr>
      <th>Валюта</th>
      <th>Доступно</th>
      <th>Зарезервированно</th>
    </tr>
    <tr v-for="item in currencies" v-bind:key="item.currency">
      <td>{{item.currency}}</td>
      <td>{{item.balance}}</td>
      <td>{{item.blocked}}</td>
    </tr>
  </table>
</template>

<script>
export default {
  name: 'currencies',
  data: function () {
    return {
      currencies: [
        {
          currency: 'EUR',
          balance: 0,
          blocked: 0
        },
        {
          currency: 'RUB',
          balance: 4.5,
          blocked: 0
        },
        {
          currency: 'USD',
          balance: 0.83,
          blocked: 0
        }
      ]
    }
  }
}
</script>

<style scoped>

</style>
