import request from '@/utils/request'

export function login(username, password) {
  return request({
    url: '/login',
    method: 'post',
    data: {
      username,
      password
    }
  })
}

export function register(pojo){
  return request({
      url: '/register',
      method: 'post',
      data: pojo
  })
}

export function getInfo() {
  return request({
    url: '/admin/info',
    method: 'get'
  })
}
