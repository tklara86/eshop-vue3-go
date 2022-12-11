import axios from 'axios';

function getData(endpoint) {
  return axios.get(`http://localhost:9090/${endpoint}`)
}

export default {
  getData: getData
}
