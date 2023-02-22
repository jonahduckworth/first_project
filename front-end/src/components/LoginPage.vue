<template>
    <div>
        <form @submit.prevent="submitForm">
            <div>
                <label>Email:</label>
                <input type="email" v-model="email" />
            </div>
            <div>
                <label>Password:</label>
                <input type="password" v-model="password" />
            </div>
            <button type="submit">Login</button>
        </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            email: "",
            password: "",
        };
    },
    methods: {
        submitForm() {
            const formData = new FormData();
            formData.append("email", this.email);
            formData.append("password", this.password);

            const xhr = new XMLHttpRequest();
            xhr.open("POST", "http://localhost:8080/login", true);
            xhr.onload = () => {
                if (xhr.status === 200) {
                    const response = JSON.parse(xhr.responseText);
                    if (response.success) {
                        this.$router.push({ path: "/dashboard" });
                    } else {
                        alert("Login failed");
                    }
                }
            };
            xhr.send(formData);
        },
    },
};
</script>
