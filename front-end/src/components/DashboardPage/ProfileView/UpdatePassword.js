import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

export default {
    methods: {
        async updatePassword() {
            const id = this.$store.state.id;
            try {
                console.log("Password reset request sent!");
                await axiosInstance.post(`/users/${id}/updatePassword`);
            } catch (error) {
                console.error(error);
                console.log("Error sending password reset request.");
            }
        },
    },
};
