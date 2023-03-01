<template>
    <div class="back-button" @click="goBack">&lt; Back</div>
    <div class="container">
        <div class="container-row">
            <div class="tile">
                <p>Tile 1</p>
            </div>
            <div class="tile">
                <p>Tile 2</p>
            </div>
            <div class="tile">
                <p>Tile 3</p>
            </div>
        </div>
        <div style="height: 15px; background-color: white"></div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            showStoreView: true,
        };
    },
    methods: {
        goBack() {
            this.$emit("update:showStoreView", false);
        },
    },
};
</script>

<style scoped>
@import "@/css/DashboardPage.css";
</style>
