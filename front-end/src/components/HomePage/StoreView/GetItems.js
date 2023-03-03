import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

export function getItems() {
    return axiosInstance.get("/items").then((response) => response.data);
}
