import HttpClient from "@/core/http/HttpClient";
import type {HttpResponse, HttpRequestConfig} from "@/core/http/HttpClient";
import {loading} from "@/core/composables/loading";

const httpClient = new HttpClient({
    baseURL: import.meta.env.VITE_TODO_WBFF_BASE_URL,
});

httpClient.interceptors.request.use((config: HttpRequestConfig) => {
    loading.value = true;
    console.log("loading");

    return config;
});

httpClient.interceptors.response.use(
    (response: HttpResponse) => {
        loading.value = false;
        console.log("loaded");

        return response;
    },
    (error) => {
        loading.value = false;
        console.log("loaded");

        throw error;
    },
);

export default httpClient;
