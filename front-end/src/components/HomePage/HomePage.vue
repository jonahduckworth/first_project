<template>
    <div>
        <div class="container-shadow"></div>
        <div class="form-container">
            <div class="login-create-account">
                <button class="btn btn-primary" @click="goToLoginPage">
                    Login
                </button>
                <span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span>
                <button class="btn btn-primary" @click="goToCreateAccountPage">
                    Create Account
                </button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    methods: {
        goToLoginPage() {
            this.$router.push({ path: "/login" });
        },
        goToCreateAccountPage() {
            this.$router.push({ path: "/create-account" });
        },
    },
};
</script>
