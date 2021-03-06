import Cookies from 'js-cookie'

const TokenKey = 'sessionid'

export function getToken() {
  return Cookies.get(TokenKey)
}

// export function setToken(token) {
//   return Cookies.set(TokenKey, token)
// }

export function setToken() {
  return Cookies.set()
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}
