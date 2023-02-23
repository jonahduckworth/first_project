<template>
    <div>
        <div class="container-shadow"></div>
        <div class="form-container">
            <div style="height: 28px; background-color: white"></div>
            <form>
                <div class="form-group">
                    <label for="email">Email:</label>
                    <input
                        type="email"
                        id="email"
                        v-model="email"
                        class="form-control"
                    />
                </div>
                <div class="form-group">
                    <label for="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        v-model="password"
                        class="form-control"
                    />
                </div>
                <button @click.prevent="submit" class="btn btn-primary">
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
};
</script>
