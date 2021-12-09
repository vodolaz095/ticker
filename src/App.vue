<template>
  <h1>{{ title }}</h1>
  <clock :key="timestamp" :timestamp="timestamp"></clock>
  <positions :key="timestamp" :positions="positions"></positions>
  <currencies :key="timestamp" :currencies="currencies"></currencies>
  <p>&copy; 2021 <a href="https://github.com/vodolaz095/ticker/">https://github.com/vodolaz095/ticker</a>. Все права защищены.</p>
</template>

<script>
import currencies from './components/currencies.vue';
import positions from './components/positions.vue';
import clock from './components/clock.vue';
import doGet from './lib/http.js';
import store from './lib/store.js';

let t;

function upd() {
  return doGet('/api/portfolio.json')
      .then(function (resp) {
        if (resp.statusCode === 200) {
          store.setTitle(resp.body.title);
          store.setPositions(resp.body.positions);
          store.setCurrencies(resp.body.currencies);
          store.setTimestamp(new Date(1000*resp.body.timestamp));
          return Promise.resolve();
        }
        throw new Error(`wrong status code ${resp.statusCode}`);
      });
}

export default {
  name: 'App',
  created() {
    window.document.title = 'Ticker';
  },
  components: {
    clock,
    currencies,
    positions,
  },
  data() {
    return store.state;
  },
  unmounted: function () {
    clearInterval(t);
  },
  mounted: function () {
    upd();
    t = setInterval(function () {
      upd().catch(function (err) {
        console.error(err);
        clearInterval(t);
      });
    }, 1000);
  }
}
</script>

<style>
body {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background-color: lavender;
  color: #2c3e50;
}

table {
  border: black solid 3px;
  width: 100%;
  border-collapse: collapse;
}
td,th {
  border: black solid 1px;
}
tr:hover {
  color: white;
  background-color: #2c3e50;
}
a:hover {
  color: white;
}
</style>
