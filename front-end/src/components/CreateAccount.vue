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
export default {
    data() {
        return {
            name: "",
            email: "",
            password: "",
        };
    },
    methods: {
        submit() {
            // Add your code to submit the form data to your back end here
        },
    },
};
</script>
