import axios from "axios";
import {loading} from "@/core/composables/loading";

const httpClient = axios.create({
    baseURL: import.meta.env.VITE_TODO_WBFF_BASE_URL,
});

httpClient.interceptors.request.use(
    config => {
        loading.value = true;
        console.log("loading");

        return config;
    }, error => {
        loading.value = false;
        console.log("loaded");

        return Promise.reject(error);
    },
);

httpClient.interceptors.response.use(
    response => {
        loading.value = false;
        console.log("loaded");

        return response;
    }, error => {
        loading.value = false;
        console.log("loaded");

        return Promise.reject(error);
    },
);

export default httpClient;
