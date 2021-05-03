import axios from "axios";

const API_URL = process.env.REACT_APP_SERVER_URL;

export function getTask(id) {
    const url = new URL(id,API_URL+'/task/').toString()
    return axios.get(url).then(resp => resp.data)
}