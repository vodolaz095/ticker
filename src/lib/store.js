'use strict';

import {reactive} from 'vue';

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

export default store;
