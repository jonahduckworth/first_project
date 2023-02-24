import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

export default {
    data() {
        return {
            email: "",
            password: "",
        };
    },
    methods: {
        submit() {
            return axiosInstance
                .post("/login", {
                    email: this.email,
                    password: this.password,
                })
                .then((response) => {
                    console.log(response.data);
                    return response;
                })
                .catch((error) => {
                    console.log(error);
                });
        },
    },
};
