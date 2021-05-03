import axios from "axios";

const API_URL = process.env.REACT_APP_SERVER_URL;

const register = (username, password) => {
  return axios.post(API_URL + "/registration", {
    username,
    password,
  });
};

const login = (username, password) => {
  return axios
    .post(API_URL + "/login", {
      username,
      password,
    })
    .then((response) => {
      if (response.data.Authorization) {
        console.log(JSON.stringify(response.data))
        localStorage.setItem("token", JSON.stringify(response.data));
        localStorage.setItem("user", JSON.stringify(username));
      }

      return response.data;
    });
};

const logout = () => {
  localStorage.removeItem("user");
  localStorage.removeItem("token");
};

const getCurrentUser = () => {
  return JSON.parse(localStorage.getItem("user"));
};

export default {
  register,
  login,
  logout,
  getCurrentUser,
};
