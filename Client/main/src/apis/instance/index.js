import axios from 'axios';
import dotenv from 'dotenv';

dotenv.config();
console.log(process.env.REACT_APP_BASEURL)
const instance = axios.create({
    baseURL: process.env.REACT_APP_BASEURL,
    timeout: 3000,
});

instance.interceptors.request.use((config) => {
    const accessToken = sessionStorage.getItem("accessToken") ||localStorage.getItem("accessToken");

    if (accessToken) {
        config.headers["Access-Token"] = accessToken;
    }

    return config;
});

export default instance;