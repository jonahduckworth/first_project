<template>
    <div class="main-container">
        <div v-for="item in items" :key="item.ID">
            <p>{{ item.Title }}</p>
            <p>{{ item.Price }}</p>
            <img :src="item.ImageUrl" />
        </div>
    </div>
</template>

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
