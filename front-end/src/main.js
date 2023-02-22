import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHashHistory } from "vue-router";
import LoginPage from "./components/LoginPage/LoginPage.vue";
import CreateAccount from "./components/CreateAccount/CreateAccount.vue";

const routes = [
    { path: "/", component: LoginPage },
    { path: "/create-account", component: CreateAccount },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

const app = createApp(App);

app.use(router);

app.mount("#app");
