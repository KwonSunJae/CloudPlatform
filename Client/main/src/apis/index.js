import axios from 'axios';
import dotenv from 'dotenv';

dotenv.config();
console.log(process.env.REACT_APP_BASEURL)
const instance = axios.create({
    baseURL: process.env.REACT_APP_BASEURL,
});

instance.interceptors.request.use((config) => {
    return config;
});

export default instance;