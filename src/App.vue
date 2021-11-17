<template>
  <clock :key="timestamp" :timestamp="timestamp"></clock>
  <positions :key="timestamp" :positions="positions"></positions>
  <currencies :key="timestamp" :currencies="currencies"></currencies>
</template>

<script>
import {reactive} from 'vue';
import currencies from './components/currencies.vue';
import positions from './components/positions.vue';
import clock from './components/clock.vue';
let t;

function doGet(url) {
  return new Promise(function (resolve, reject) {
    const r = new XMLHttpRequest();
    r.onreadystatechange = function () {
      if (r.readyState === 4) {
        if (r.status > 204) {
          return reject(new Error(`wrong status code ${r.status}`));
        }
        let body;
        try {
          body = JSON.parse(r.responseText);
        } catch (err) {
          body = r.responseText;
        }
        resolve({
          statusCode: r.status,
          body
        });
      }
    };
    r.open('GET', url, true);
    r.send(null);
  });
}

const store = {
  state: reactive({
    timestamp: new Date(),
    positions: [],
    currencies: [],
  }),
  setPositions(newValue) {
    this.state.positions = newValue;
  },
  setCurrencies(newValue) {
    this.state.currencies = newValue;
  },
  setTimestamp(newValue) {
    this.state.timestamp = newValue;
  }
};

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
    function upd() {
      return doGet('/api/portfolio.json')
          .then(function (resp) {
            if (resp.statusCode === 200) {
              store.setPositions(resp.body.positions);
              store.setCurrencies(resp.body.currencies);
              store.setTimestamp(new Date(1000*resp.body.timestamp));
            }
          });
    }
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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
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
