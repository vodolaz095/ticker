'use strict';

import {reactive} from 'vue';

const store = {
  state: reactive({
    title: String,
    timestamp: new Date(),
    positions: [],
    currencies: [],
  }),
  setTitle(newValue) {
    this.state.title = newValue;
  },
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
