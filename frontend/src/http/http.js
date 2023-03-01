import axios from 'axios'

let api;

function startApi(){
    const host = process.env.NODE_ENV === 'development' ? "localhost:8080" : window.location.host;
    api = axios.create({
        baseURL: `http://${host}/api`
    });

    api.interceptors.response.use((res) => {
        return res.data;
    });
}

startApi();

export const fetchField = () => {return api.get('/getField')}

