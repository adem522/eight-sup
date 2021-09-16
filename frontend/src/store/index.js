import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    username: "",
    type: "",
    token: null,
    storedInfo:[],
    host: "http://localhost:8080",
  },
  mutations: { 
    setUsername(state, val) {
      state.username = val;
    },
    setToken(state, val) {
      state.token = val;
    },
    setType(state, val) {
      state.type = val;
    },
    setStoredInfo(state,val){
      state.storedInfo=val;
    },
    logout(state) {
      localStorage.removeItem("token");
      localStorage.removeItem("username");
      localStorage.removeItem("type");
      state.storedInfo = [];
      state.username = "";
      state.type = "";
      state.token = null;
    },
  },
});

export default store;
