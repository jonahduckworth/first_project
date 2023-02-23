<template>
    <div>
        <div class="create-account-text">Create Account</div>
        <div class="container-shadow"></div>
        <div class="form-container">
            <form>
                <div class="form-group">
                    <label for="name"> Name:</label>
                    <input
                        type="text"
                        id="name"
                        v-model="name"
                        class="form-control"
                    />
                </div>
                <div class="form-group">
                    <label for="email"> Email:</label>
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
                <button
                    @click.prevent="submit"
                    class="btn btn-primary"
                    @click="goToLoginPage"
                >
                    Create
                </button>
            </form>
        </div>
    </div>
</template>

<style scoped>
@import "./CreateAccount.css";
</style>

<script>
import createAccount from "./CreateAccount.js";

export default {
    mixins: [createAccount],
    methods: {
        goToLoginPage() {
            this.$router.push({ path: "/login" });
        },
    },
};
</script>
