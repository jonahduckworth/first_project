<template>
    <div class="home-container">
        <h1>Add Item</h1>
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
        submitForm() {
            // Add code to handle the submission of the form data, such as sending it to a server or storing it locally
            console.log("Title:", this.title);
            console.log("Price:", this.price);
            console.log("Image:", this.image);
        },
    },
};
</script>
