'use strict';

/**
 * @type {import('@vue/cli-service').ProjectOptions}
 */

const path = require('path');

module.exports = {
  outputDir: path.join(__dirname, 'public'),
  publicPath: '/public/',
  indexPath: 'main.html',
  integrity: true,
  devServer: {
    proxy: 'http://localhost:3100'
  }
}
