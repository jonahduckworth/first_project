import { createApp } from "vue";
import App from "./App.vue";
import store from "./store.js";
import { createRouter, createWebHashHistory } from "vue-router";
import LoginPage from "./components/LoginPage/LoginPage.vue";
import DashboardPage from "./components/DashboardPage/DashboardPage.vue";

const routes = [
    { path: "/", component: LoginPage },
    { path: "/dashboard", component: DashboardPage },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

const app = createApp(App);

app.use(store);
app.use(router);

app.mount("#app");
