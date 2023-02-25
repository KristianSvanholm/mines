import axios from 'axios'

let api;

function startApi(){
    api = axios.create({
        baseURL: 'http://localhost:8080/api'
    });

    api.interceptors.response.use((res) => {
        return res.data;
    });
}

startApi();

export const fetchField = () => {return api.get('/getField')}

