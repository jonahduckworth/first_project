<template>
    <div class="navigation-bar">
        <div class="container-row">
            <div
                class="tile"
                @click="
                    showProfileView = true;
                    showStoreView = false;
                    showAddItemView = false;
                "
            >
                <p>Profile</p>
            </div>
            <div
                class="tile"
                @click="
                    showStoreView = true;
                    showProfileView = false;
                    showAddItemView = false;
                "
            >
                <p>Store</p>
            </div>
            <div
                class="tile"
                @click="
                    showAddItemView = true;
                    showProfileView = false;
                    showStoreView = false;
                "
            >
                <p>Add Item</p>
            </div>
            <div class="tile" @click="signOut">
                <p>Sign Out</p>
            </div>
        </div>
        <div style="height: 10px; background-color: white"></div>
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
            showProfileView: true,
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
@import "@/css/NavigationBar.css";
</style>
