import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost:8081",
});

export function submit(title, price, image) {
    const formData = new FormData();
    formData.append("title", title);
    formData.append("price", price);
    formData.append("image", image);
    formData.append("quantity", 1);

    return new Promise((resolve, reject) => {
        axiosInstance
            .post("/addItem", formData)
            .then((response) => {
                resolve(response);
            })
            .catch((error) => {
                reject(error.response);
            });
    });
}
