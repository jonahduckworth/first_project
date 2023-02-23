<template>
    <div>
        <div class="container-shadow"></div>
        <div class="form-container">
            <div style="height: 45px; background-color: white"></div>
            <button class="btn btn-primary" @click="goToLoginPage">
                Login
            </button>
            <div style="height: 40px; background-color: white"></div>
            <button class="btn btn-primary" @click="goToCreateAccountPage">
                Create Account
            </button>
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
