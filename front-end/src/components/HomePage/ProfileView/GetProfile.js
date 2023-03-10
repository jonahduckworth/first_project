import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

export default {
    data() {
        return {
            name: "",
            email: "",
        };
    },
    methods: {
        async getProfile() {
            const id = this.$store.state.id;
            try {
                const response = await axiosInstance.get(`/users/${id}`);
                this.name = response.data.name;
                this.email = response.data.email;
            } catch (error) {
                console.error(error);
            }
        },
    },
    created() {
        this.getProfile();
    },
};
