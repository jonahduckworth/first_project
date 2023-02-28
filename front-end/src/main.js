import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHashHistory } from "vue-router";
import HomePage from "./components/HomePage/HomePage.vue";
import DashboardPage from "./components/DashboardPage/DashboardPage.vue";

const routes = [
    { path: "/", component: HomePage },
    { path: "/dashboard", component: DashboardPage },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

const app = createApp(App);

app.use(router);

app.mount("#app");
