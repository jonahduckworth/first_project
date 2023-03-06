<template>
    <div class="main-container">
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="image">Image:</label>
                <input type="file" id="image" @change="uploadImage" />
            </div>
            <div class="form-group">
                <label for="title">Title:</label>
                <input type="text" id="title" v-model="title" />
            </div>
            <div class="form-group">
                <label for="price">Price:</label>
                <input type="text" id="price" v-model="price" />
            </div>
            <div class="form-group">
                <button type="submit">Add Item</button>
            </div>
        </form>
    </div>
</template>

<script>
import { submit } from "./AddItem.js";
export default {
    data() {
        return {
            image: null,
            title: "",
            price: "",
        };
    },
    methods: {
        uploadImage(event) {
            this.image = event.target.files[0];
        },
        async submitForm() {
            try {
                await submit(this.title, this.price, this.image);
            } catch (error) {
                console.error(error);
            }
        },
    },
};
</script>
