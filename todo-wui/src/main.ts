import {createApp} from "vue";
import VueAxios from "vue-axios";

import AppView from "@/core/views/AppView.vue";
import router from "@/core/routers/router";
import axios from "@/core/composables/axios";
import authUseCasePlugin from "@/plugins/auth-use-case-plugin";
import todoUseCasePlugin from "@/plugins/todo-use-case-plugin";
import "@/core/assets/main.css";
import "@fortawesome/fontawesome-free/css/all.min.css"

const app = createApp(AppView);

app.use(router);
app.use(VueAxios, axios);
app.use(authUseCasePlugin, router);
app.use(todoUseCasePlugin, router);

app.mount("#app");
