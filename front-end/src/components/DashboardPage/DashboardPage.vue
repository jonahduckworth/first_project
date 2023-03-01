<template>
    <div v-if="!showProfileView && !showStoreView">
        <div class="container">
            <div class="container-row">
                <div class="tile" @click="showProfileView = true">
                    <p>Profile</p>
                </div>
                <div class="tile" @click="showStoreView = true">
                    <p>Store</p>
                </div>
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
    <StoreView
        v-if="showStoreView"
        @update:showStoreView="showStoreView = false"
    />
</template>
<script>
import ProfileView from "./ProfileView/ProfileView.vue";
import StoreView from "./StoreView/StoreView.vue";
import GetProfile from "./ProfileView/GetProfile.js";

export default {
    mixins: [GetProfile],
    components: {
        ProfileView,
        StoreView,
    },
    data() {
        return {
            showProfileView: false,
            showStoreView: false,
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
