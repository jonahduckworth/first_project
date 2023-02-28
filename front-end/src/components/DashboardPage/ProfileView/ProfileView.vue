<template>
    <div class="back-button" @click="goBack">&lt; Back</div>
    <div class="profile-container">
        <h3>Profile</h3>
        <div class="profile-info">
            <p>Name: {{ name }}</p>
            <p>Email: {{ email }}</p>
        </div>
    </div>
</template>

<script>
import GetProfile from "./GetProfile.js";

export default {
    mixins: [GetProfile],
    methods: {
        goBack() {
            this.$emit("update:showProfileView", false);
        },
    },
};
</script>
