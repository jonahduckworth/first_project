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
                        :class="{ 'is-invalid': emailError }"
                    />
                    <div class="invalid-feedback" v-if="emailError">
                        Please enter your email.
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
                        Please enter your password.
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

.is-invalid {
    border-color: red;
}

.invalid-feedback {
    color: red;
    font-size: 10px;
    margin-top: 5px;
}
</style>

<script>
import login from "./Login.js";

export default {
    mixins: [login],
    data() {
        return {
            emailError: false,
            passwordError: false,
        };
    },
    methods: {
        async handleSubmit() {
            this.emailError = false;
            this.passwordError = false;
            if (!this.email) {
                this.emailError = true;
            }
            if (!this.password) {
                this.passwordError = true;
            }
            if (!this.email || !this.password) {
                return;
            }
            const response = await this.submit();
            if (response.data.userExists) {
                this.$router.push({ path: "/dashboard" });
            }
        },
    },
};
</script>
