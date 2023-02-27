<template>
    <div>
        <div
            class="back-button"
            v-if="showLoginForm || showCreateAccountForm"
            @click="goBack"
        >
            &lt; Back
        </div>
        <div class="home-container-shadow"></div>
        <div class="home-container">
            <div style="height: 15px; background-color: white"></div>
            <div class="login-create-account">
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
@import "./HomePage.css";
</style>
<script>
import LoginForm from "@/components/LoginPage/LoginPage.vue";
import CreateAccountForm from "@/components/CreateAccount/CreateAccount.vue";

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
