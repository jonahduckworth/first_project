<template>
    <div>
        <div class="container-shadow"></div>
        <div class="form-container">
            <div style="height: 25px; background-color: white"></div>
            <form>
                <div class="form-group">
                    <label for="email">Email:</label>
                    <input
                        type="email"
                        id="email"
                        v-model="email"
                        class="form-control"
                        :class="{
                            'is-invalid': emailError || emailFormatError,
                        }"
                    />
                    <div class="invalid-feedback" v-if="emailError">
                        Please enter email.
                    </div>
                    <div class="invalid-feedback" v-if="emailFormatError">
                        Please enter a valid email address.
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
                        Please enter password.
                    </div>
                    <div class="invalid-feedback" v-if="loginError">
                        Invalid email or password.
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

<style scoped>
@import "./LoginPage.css";
</style>

<script>
import login from "./Login.js";

export default {
    mixins: [login],
    data() {
        return {
            email: "",
            password: "",
            emailError: false,
            passwordError: false,
            loginError: false,
            emailFormatError: false,
        };
    },
    methods: {
        async handleSubmit() {
            this.emailError = false;
            this.passwordError = false;
            this.loginError = false;
            this.emailFormatError = !this.validateEmail(this.email);
            if (!this.email) {
                this.emailError = true;
            }
            if (!this.password) {
                this.passwordError = true;
            }
            if (!this.email || !this.password || this.emailFormatError) {
                return;
            }
            const response = await this.submit();
            if (response.data.userExists) {
                this.$router.push({ path: "/dashboard" });
            } else {
                this.loginError = true;
            }
        },
        validateEmail(email) {
            // Source: https://stackoverflow.com/a/46181/1420960
            const emailRegex =
                /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
            return emailRegex.test(String(email).toLowerCase());
        },
    },
};
</script>
