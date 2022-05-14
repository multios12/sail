import axios from "axios"

const secretMode = localStorage.getItem("")

export default axios.interceptors.response.use(config => {
    if (secretMode === "true") {
    } else {
    }
  });