export default function ({ $axios }) {
  $axios.onRequest((config) => {
    // set authorization header
    $axios.defaults.headers.common.Authorization = localStorage.getItem('token')
    return config
  })
}
