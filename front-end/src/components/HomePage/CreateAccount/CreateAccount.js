import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

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
            return new Promise((resolve, reject) => {
                axiosInstance
                    .post("/users", {
                        name: this.name,
                        email: this.email,
                        password: this.password,
                    })
                    .then((response) => {
                        resolve(response);
                    })
                    .catch((error) => {
                        reject(error.response);
                    });
            });
        },
    },
};
