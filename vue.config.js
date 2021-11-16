'use strict';

/**
 * @type {import('@vue/cli-service').ProjectOptions}
 */

const path = require('path');

module.exports = {
  outputDir: path.join(__dirname, 'public'),
  indexPath: 'main.html',
  devServer: {
    proxy: 'http://localhost:3100'
  }
}
