<template>
  <h2>Позиции</h2>
  <table>
    <thead>
    <tr>
      <th>Тикер</th>
      <th>Тип</th>
      <th>Название</th>
      <th>Лоты</th>
      <th title="Средняя цена покупки позиции" >Стоимость</th>
      <th title="Текущая цена позиции на рынке" >Рыночная цена</th>
      <th title="Сколько денег можно получить, продав позицию сейчас">Цена продажи</th>
      <th>Доход</th>
      <th>%</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="position in positions" v-bind:key="position.ticker">
      <td>{{ position.ticker }}</td>
      <td v-if="position.instrumentType==='Stock'">Акция</td>
      <td v-if="position.instrumentType==='Bond'">Облигация</td>
      <td v-if="position.instrumentType==='Etf'">Фонд</td>
      <td v-if="position.instrumentType==='Currency'">Валюта</td>
      <td>
        <a target="_blank" v-bind:href="'https://tinkoff.ru/invest/'+position.instrumentType.toString().toLowerCase()+'s/'+position.ticker+'/'"
        >{{ position.name }}</a>
      </td>
      <td>{{ position.balance }}</td>
      <td>
        {{
          new Intl.NumberFormat('ru-RU', {style: 'currency', currency: position.averagePositionPrice.currency})
              .format(position.averagePositionPrice.value)
        }}
      </td>
      <td>
        {{
          new Intl.NumberFormat('ru-RU', {style: 'currency', currency: position.averagePositionPrice.currency})
              .format(position.averagePositionPrice.value+position.expectedYield.value/position.balance)
        }}
      </td>
      <td>
        {{
          new Intl.NumberFormat('ru-RU', {style: 'currency', currency: position.averagePositionPrice.currency})
              .format(position.balance * position.averagePositionPrice.value+position.expectedYield.value)
        }}
      </td>
      <td v-bind:class="{
        profit: parseFloat(position.expectedYield.value, 10) >= 0,
        loss: parseFloat(position.expectedYield.value, 10) < 0
      }">
        {{
          new Intl.NumberFormat('ru-RU', {style: 'currency', currency: position.averagePositionPrice.currency})
              .format(position.expectedYield.value)
        }}
      </td>
      <td v-bind:class="{
        profit: parseFloat(position.expectedYield.value, 10) >= 0,
        loss: parseFloat(position.expectedYield.value, 10) < 0
      }">
        {{
          Number(position.expectedYield.value / (position.balance * position.averagePositionPrice.value))
              .toLocaleString('ru-RU',{style: 'percent', minimumFractionDigits:2})
        }}
      </td>
    </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  name: 'positions',
  props: {'positions': Array},
}
</script>

<style scoped>
td.profit {
  color: green;
}

td.loss {
  color: red;
}
</style>

