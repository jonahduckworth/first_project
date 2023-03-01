import { createStore } from "vuex";

const store = createStore({
    state() {
        return {
            id: 0,
        };
    },
    mutations: {
        setUser(state, id) {
            state.id = id;
            console.log(state.id);
        },
    },
    actions: {
        login({ commit }, user) {
            commit("setUser", user);
            console.log("Logged in");
        },
    },
});

export default store;
