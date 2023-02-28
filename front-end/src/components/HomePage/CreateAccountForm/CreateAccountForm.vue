<template>
    <div>
        <div class="login-container">
            <div style="height: 15px; background-color: white"></div>
            <form>
                <div class="form-group">
                    <label for="name">Name:</label>
                    <input
                        type="text"
                        id="name"
                        v-model="name"
                        class="form-control"
                        :class="{ 'is-invalid': nameError }"
                    />
                    <div class="invalid-feedback" v-if="nameError">
                        Please enter your name
                    </div>
                </div>
                <div class="form-group">
                    <label for="email">Email:</label>
                    <input
                        type="email"
                        id="email"
                        v-model="email"
                        class="form-control"
                        :class="{
                            'is-invalid': emailError,
                        }"
                    />
                    <div class="invalid-feedback" v-if="emailError">
                        Please enter a valid email address
                    </div>
                </div>
                <div class="form-group">
                    <label for="password">Password:</label>
                    <input
                        type="text"
                        id="password"
                        v-model="password"
                        class="form-control"
                        :class="{ 'is-invalid': passwordError }"
                    />
                    <div
                        class="invalid-feedback"
                        v-if="!passwordError"
                        style="height: 10px; background-color: white"
                    ></div>
                    <div class="invalid-feedback" v-if="passwordError">
                        Please enter a password
                    </div>
                </div>
                <div style="height: 20px; background-color: white"></div>
                <button
                    @click.prevent="handleSubmit"
                    class="btn btn-primary"
                    @click="goToRoot"
                >
                    Create
                </button>
            </form>
        </div>
    </div>
</template>

<script>
import createAccount from "./CreateAccount.js";
import { validateEmail } from "@/utils.js";

export default {
    mixins: [createAccount],
    data() {
        return {
            name: "",
            email: "",
            password: "",
            nameError: false,
            emailError: false,
            passwordError: false,
        };
    },
    methods: {
        async handleSubmit() {
            this.nameError = false;
            this.passwordError = false;
            this.emailError = false;
            if (!this.name) {
                this.nameError = true;
            }
            if (!this.email) {
                this.emailError = true;
            }
            if (!this.password) {
                this.passwordError = true;
            }
            this.emailError = !validateEmail(this.email);
            if (
                !this.name ||
                !this.email ||
                !this.password ||
                this.emailError
            ) {
                return;
            }
            const response = await this.submit();
            if (response.data.userCreated) {
                this.$emit("update:showCreateAccountForm", false);
                this.$emit("update:showLoginForm", true);
            }
        },
    },
};
</script>
