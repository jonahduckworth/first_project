<template>
    <div class="form-container">
        <h1>Login</h1>
        <form>
            <div class="form-group">
                <label for="email">Email:</label>
                <input
                    type="email"
                    id="email"
                    v-model="email"
                    class="form-control"
                />
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input
                    type="password"
                    id="password"
                    v-model="password"
                    class="form-control"
                />
            </div>
            <button @click.prevent="submit" class="btn btn-primary">
                Submit
            </button>
        </form>
        <p>
            Don't have an account yet?
            <router-link to="/create-account">Create one</router-link>.
        </p>
    </div>
</template>
<style scoped>
@import "../css/LoginPage.css";
</style>
<script>
import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

export default {
    data() {
        return {
            email: "",
            password: "",
        };
    },
    methods: {
        submit() {
            axiosInstance
                .post("/login", {
                    email: this.email,
                    password: this.password,
                })
                .then((response) => {
                    console.log(response.data);
                })
                .catch((error) => {
                    console.log(error);
                });
        },
    },
};
</script>
