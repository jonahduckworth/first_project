<template>
    <div>
        <div
            class="back-button"
            v-if="showLoginForm || showCreateAccountForm"
            @click="goBack"
        >
            &lt; Back
        </div>
        <div class="main-container">
            <div>
                <button class="btn btn-primary" @click="showLoginForm = true">
                    Login
                </button>
                <div style="height: 30px; background-color: white"></div>
                <button
                    class="btn btn-primary"
                    @click="showCreateAccountForm = true"
                >
                    Create Account
                </button>
            </div>
            <div v-if="showLoginForm">
                <LoginForm />
            </div>
            <div v-if="showCreateAccountForm">
                <CreateAccountForm
                    @update:showCreateAccountForm="
                        showCreateAccountForm = false
                    "
                    @update:showLoginForm="showLoginForm = true"
                />
            </div>
        </div>
    </div>
</template>

<style scoped>
@import "@/css/HomePage.css";
</style>

<script>
import LoginForm from "./LoginForm/LoginForm.vue";
import CreateAccountForm from "./CreateAccountForm/CreateAccountForm.vue";

export default {
    components: {
        LoginForm,
        CreateAccountForm,
    },
    methods: {
        goBack() {
            this.showLoginForm = false;
            this.showCreateAccountForm = false;
        },
    },
    data() {
        return {
            showLoginForm: false,
            showCreateAccountForm: false,
        };
    },
};
</script>
