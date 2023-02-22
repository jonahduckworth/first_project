<template>
    <div>
        <h1>Create an Account</h1>
        <form>
            <div>
                <label for="name">Name:</label>
                <input type="text" id="name" v-model="name" />
            </div>
            <div>
                <label for="email">Email:</label>
                <input type="email" id="email" v-model="email" />
            </div>
            <div>
                <label for="password">Password:</label>
                <input type="password" id="password" v-model="password" />
            </div>
            <button @click.prevent="submit">Submit</button>
        </form>
    </div>
</template>
<script>
import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

export default {
    data() {
        return {
            name: "",
            email: "",
            password: "",
        };
    },
    methods: {
        submit() {
            axiosInstance
                .post("/users", {
                    name: this.name,
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
