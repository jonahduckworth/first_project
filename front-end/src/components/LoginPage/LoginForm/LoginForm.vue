<template>
    <div>
        <div class="main-container">
            <div style="height: 15px; background-color: white"></div>
            <form>
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
                        type="password"
                        id="password"
                        v-model="password"
                        class="form-control"
                        :class="{ 'is-invalid': passwordError }"
                    />
                    <div class="invalid-feedback" v-if="passwordError">
                        Please enter password
                    </div>
                    <div
                        class="invalid-feedback"
                        v-if="!loginError && !passwordError"
                        style="height: 10px; background-color: white"
                    ></div>
                    <div class="invalid-feedback" v-if="loginError">
                        Invalid email or password
                    </div>
                </div>
                <div style="height: 20px; background-color: white"></div>
                <button @click.prevent="handleSubmit" class="btn btn-primary">
                    Login
                </button>
            </form>
        </div>
    </div>
</template>
<script>
import login from "./Login.js";
import { validateEmail } from "@/utils.js";

export default {
    mixins: [login],
    data() {
        return {
            email: "",
            password: "",
            passwordError: false,
            loginError: false,
            emailError: false,
        };
    },
    methods: {
        async handleSubmit() {
            this.emailError = false;
            this.passwordError = false;
            this.loginError = false;
            this.emailError = !validateEmail(this.email);
            if (!this.email) {
                this.emailError = true;
            }
            if (!this.password) {
                this.passwordError = true;
            }
            if (!this.email || !this.password || this.emailError) {
                return;
            }
            const response = await this.submit();
            if (response.data.userExists) {
                this.$store.dispatch("login", response.data.id);
                this.$router.push({ path: "/dashboard" });
            } else {
                this.loginError = true;
            }
        },
    },
};
</script>
