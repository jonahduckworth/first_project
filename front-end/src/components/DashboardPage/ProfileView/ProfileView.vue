<template>
    <div>
        <div class="back-button" @click="goBack">&lt; Back</div>
        <div class="profile-container">
            <h3>Profile</h3>
            <div class="profile-info">
                <p>Name: {{ name }}</p>
                <p>Email: {{ email }}</p>
                <div class="password-container" v-if="!editingPassword">
                    <p>
                        Password: ********
                        <button @click="editingPassword = true">Edit</button>
                    </p>
                </div>
                <form v-if="editingPassword" @submit.prevent="updatePassword">
                    <p>
                        Password:
                        <input
                            type="text"
                            v-model="password"
                            placeholder="New Password"
                        />
                        <button type="submit">Save</button>
                    </p>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import GetProfile from "./GetProfile.js";
import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

export default {
    mixins: [GetProfile],
    data() {
        return {
            password: "",
            editingPassword: false,
        };
    },
    methods: {
        goBack() {
            this.$emit("update:showProfileView", false);
        },
        async updatePassword() {
            try {
                const response = await axiosInstance.put(`/updatePassword`, {
                    email: this.email,
                    password: this.password,
                });
                if (response.data.passwordUpdated) {
                    this.editingPassword = false;
                    this.password = "";
                } else {
                    console.error("Failed to update password");
                }
            } catch (error) {
                console.error(error);
            }
        },
    },
};
</script>
