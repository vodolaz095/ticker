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
  props: {'currencies':Array}
}
</script>

<style scoped>

</style>
