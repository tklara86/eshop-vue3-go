import axios from 'axios';


function getData(endpoint) {
  return axios.get(`${import.meta.env.VITE_VUE_APP_API_BASE_URL}/${endpoint}`)
}

export default { getData };


