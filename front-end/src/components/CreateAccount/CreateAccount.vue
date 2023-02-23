<template>
    <div class="form-container">
        <h1>Create an Account</h1>
        <form>
            <div class="form-group">
                <label for="name">Name:</label>
                <input
                    type="text"
                    id="name"
                    v-model="name"
                    class="form-control"
                />
            </div>
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
                Submit
            </button>
        </form>
    </div>
</template>

<style scoped>
@import "./CreateAccount.css";
</style>

<script>
import createAccount from "./CreateAccount.js";

export default {
    mixins: [createAccount],
};
</script>
