import axios from 'axios';

const instance = axios.create();

instance.interceptors.request.use((config) => {
    return config;
});

export default instance;