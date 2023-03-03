<template>
    <div>
        <div class="profile-container">
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
                    <div class="invalid-feedback" v-if="errorMessage">
                        {{ errorMessage }}
                    </div>
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
            errorMessage: "",
        };
    },
    methods: {
        async updatePassword() {
            this.errorMessage = "";
            try {
                const response = await axiosInstance.put(`/updatePassword`, {
                    email: this.email,
                    password: this.password,
                });
                if (response.data.passwordUpdated) {
                    this.editingPassword = false;
                    this.password = "";
                } else {
                    this.errorMessage = "Failed to update password";
                }
            } catch (error) {
                console.error(error);
                this.errorMessage = error.response.data.error;
            }
        },
    },
};
</script>
