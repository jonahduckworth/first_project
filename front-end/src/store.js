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
        },
    },
    actions: {
        login({ commit }, user) {
            commit("setUser", user);
        },
    },
});

export default store;
