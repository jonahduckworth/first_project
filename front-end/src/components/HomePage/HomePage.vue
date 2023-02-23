<template>
    <div>
        <div
            class="back-button"
            v-if="showLoginForm || showCreateAccountForm"
            @click="goBack"
        >
            &lt; Back
        </div>
        <div class="container-shadow"></div>
        <div class="form-container">
            <div class="login-create-account">
                <div style="height: 45px; background-color: white"></div>
                <button class="btn btn-primary" @click="showLoginForm = true">
                    Login
                </button>
                <div style="height: 40px; background-color: white"></div>
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
