import {createApp} from "vue";
import axios from "axios";
import VueAxios from "vue-axios";

import AppView from "@/core/views/AppView.vue";
import router from "@/core/routers/router";
import httpClient from "@/core/composables/http-client";
import authUseCasePlugin from "@/plugins/auth-use-case-plugin";
import todoUseCasePlugin from "@/plugins/todo-use-case-plugin";
import "@/core/assets/main.css";
import "@fortawesome/fontawesome-free/css/all.min.css"

const app = createApp(AppView);

app.use(router);
app.use(VueAxios, axios);
app.use(authUseCasePlugin, router, httpClient);
app.use(todoUseCasePlugin, router, httpClient);

app.mount("#app");
