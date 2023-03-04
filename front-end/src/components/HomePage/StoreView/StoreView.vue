<template>
    <div class="store-container">
        <div class="container-row">
            <div class="item-container" v-for="item in items" :key="item.ID">
                <img :src="item.ImageUrl" />
                <p>{{ item.Title }}</p>
                <p>{{ item.Price }}</p>
            </div>
        </div>
    </div>
</template>

<style>
.store-container {
    width: 85%;
    min-width: 365px;
    height: 67%;
    margin: 0 auto;
    padding: 30px;
    background-color: white;
    border-radius: 0;
    border: 1px solid black;
    position: absolute;
    top: 58%;
    left: 50%;
    transform: translate(-50%, -50%);
    box-shadow: 15px 15px 0px 0px black;
}

.container-row {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-evenly;
}

.item-container {
    width: 100px;
    height: 200px;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: 0 auto;
    padding: 30px;
    background-color: white;
    border-radius: 0;
    border: 1px solid black;
    box-shadow: 5px 5px 0px 0px black;
}

img {
    width: 100px;
    height: 100px;
    object-fit: cover;
    margin-bottom: 20px;
}
</style>

<script>
import { getItems } from "./GetItems";

export default {
    data() {
        return {
            items: [],
        };
    },
    mounted() {
        this.getItems();
    },
    methods: {
        async getItems() {
            try {
                const items = await getItems();
                this.items = items;
                console.log(items);
            } catch (error) {
                console.error(error);
            }
        },
    },
};
</script>
