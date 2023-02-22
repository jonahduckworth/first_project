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
@import "./LoginPage.css";
</style>

<script>
import loginPage from "./LoginPage.js";

export default {
    mixins: [loginPage],
};
</script>
