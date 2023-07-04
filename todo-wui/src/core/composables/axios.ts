import axios from "axios";

axios.interceptors.request.use(
    config => {
        const token = localStorage.getItem("token")
        if (token != null) {
            config.headers.Authorization = `Bearer ${token}`
        }

        return config
    },
);

export default axios;