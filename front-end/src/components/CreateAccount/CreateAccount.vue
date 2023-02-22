<template>
    <div>
        <h1>Create an Account</h1>
        <form>
            <div>
                <label for="name">Name:</label>
                <input type="text" id="name" v-model="name" />
            </div>
            <div>
                <label for="email">Email:</label>
                <input type="email" id="email" v-model="email" />
            </div>
            <div>
                <label for="password">Password:</label>
                <input type="password" id="password" v-model="password" />
            </div>
            <button @click.prevent="submit">Submit</button>
        </form>
    </div>
</template>

<script>
import createAccount from "./CreateAccount.js";

export default {
    mixins: [createAccount],
};
</script>
