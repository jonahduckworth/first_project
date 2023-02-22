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
.form-container {
    width: 400px;
    margin: 0 auto;
    padding: 40px;
    background-color: #f7f7f7;
    border-radius: 10px;
}

.form-group {
    margin-bottom: 25px;
}

.form-control {
    width: 100%;
    padding: 10px;
    font-size: 16px;
    border-radius: 5px;
    border: 1px solid #ccc;
}

.btn {
    width: 100%;
    padding: 10px;
    font-size: 16px;
    border-radius: 5px;
    background-color: #007bff;
    color: #fff;
    border: 0;
    cursor: pointer;
}
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
