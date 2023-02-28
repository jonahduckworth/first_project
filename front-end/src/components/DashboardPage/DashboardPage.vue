<template>
    <div v-if="!showProfileView">
        <div class="container">
            <div class="container-row">
                <div class="tile" @click="showProfileView = true">
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
    <ProfileView
        v-if="showProfileView"
        @update:showProfileView="showProfileView = false"
        :name="name"
        :email="email"
    />
</template>

<script>
import ProfileView from "./ProfileView/ProfileView.vue";
import GetProfile from "./ProfileView/GetProfile.js";

export default {
    mixins: [GetProfile],
    components: {
        ProfileView,
    },
    data() {
        return {
            showProfileView: false,
        };
    },
    methods: {
        signOut() {
            this.$router.push({ path: "/" });
        },
    },
};
</script>

<style scoped>
@import "@/css/DashboardPage.css";
</style>
