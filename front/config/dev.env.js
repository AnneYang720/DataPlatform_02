'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  // BASE_API: '"http://easy-mock.anneyang.me/mock/614a9ce1e235063d550645d4/database_02/"',
  BASE_API: '"http://127.0.0.1:9000/api"',
})
