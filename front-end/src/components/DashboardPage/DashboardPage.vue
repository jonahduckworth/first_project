<template>
    <div v-if="!showProfileView && !showStoreView && !showAddItemView">
        <div class="container">
            <div class="container-row">
                <div class="tile" @click="showProfileView = true">
                    <p>Profile</p>
                </div>
                <div class="tile" @click="showStoreView = true">
                    <p>Store</p>
                </div>
                <div class="tile" @click="showAddItemView = true">
                    <p>Add Item</p>
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
    <AddItem
        v-if="showAddItemView"
        @update:showAddItemView="showAddItemView = false"
    />
</template>

<script>
import ProfileView from "./ProfileView/ProfileView.vue";
import StoreView from "./StoreView/StoreView.vue";
import AddItem from "./AddItem/AddItem.vue";
import GetProfile from "./ProfileView/GetProfile.js";

export default {
    mixins: [GetProfile],
    components: {
        ProfileView,
        StoreView,
        AddItem,
    },
    data() {
        return {
            showProfileView: false,
            showStoreView: false,
            showAddItemView: false,
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
