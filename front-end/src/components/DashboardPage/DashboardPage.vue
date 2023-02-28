<template>
    <div>
        <div class="container">
            <div class="container-row">
                <div class="tile">
                    <p>Profile</p>
                </div>
                <div class="tile"></div>
                <div class="tile"></div>
                <div class="tile"></div>
                <div class="tile"></div>
                <div class="tile"></div>
                <div class="tile"></div>
                <div class="tile" @click="signOut">
                    <p>Sign Out</p>
                </div>
            </div>
            <div style="height: 15px; background-color: white"></div>
        </div>
    </div>
</template>

<style scoped>
@import "@/css/DashboardPage.css";
</style>

<script>
export default {
    methods: {
        signOut() {
            this.$router.push({ path: "/" });
        },
    },
};
</script>
