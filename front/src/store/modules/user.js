import { login, logout, getInfo } from '@/api/login'
import { getToken, setToken, removeToken } from '@/utils/auth'
import  { resetRouter } from '@/router'

const user = {
  state: {
    username: '',
    userId:'',
    roles: [],
    login: false
  },

  mutations: {
    SET_USERNAME: (state, username) => {
      state.username = username
    },
    SET_USERID: (state, userId) => {
      state.userId = userId
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_LOGIN: (state, logged) => {
      state.login = logged
    },
  },

  actions: {
    // 登录
    Login({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        login(userInfo.username, userInfo.password).then( response => {
          commit('SET_LOGIN', true)
          
          // console.log("set token")
          // console.log(response.Message)
          // setToken(response.Message)
          // console.log(getToken())
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 获取用户信息
    GetInfo({ commit, state }) {
      return new Promise((resolve, reject) => {
        getInfo().then(response => {
          // roles must be a non-empty array
          if (!response.Data.Auth || response.Data.Auth.length <= 0) {
            reject('getInfo: roles must be a non-null array!')
          }
          
          commit('SET_ROLES', response.Data.Auth)
          commit('SET_USERNAME', response.Data.Username)
          commit('SET_USERID', response.Data.Id)
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 登出
    LogOut({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        // logout(state.token).then(() => {
          // commit('SET_TOKEN', '')
          removeToken()
          resetRouter()
          resolve()
        // }).catch(error => {
        //   reject(error)
        // })
      })
    },

    // 前端 登出
    FedLogOut({ commit }) {
      return new Promise(resolve => {
        // commit('SET_TOKEN', '')
        removeToken()
        resetRouter()
        resolve()
      })
    }
  }
}

export default user
