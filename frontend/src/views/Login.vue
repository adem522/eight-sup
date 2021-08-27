<template>
  <div
    style="
      width: 50%;
      display: flex;
      align-items: center;
      justify-content: flex-start;
      flex-direction: column;
      margin: 0 auto;
    "
  >
    <v-text-field
      v-model="username"
      :rules="[rules.required]"
      label="Username"
      style="width: 100%"
    />
    <v-text-field
      v-model="password"
      :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
      :rules="[rules.required, rules.min]"
      :type="show ? 'text' : 'password'"
      label="Password"
      hint="At least 8 characters"
      @click:append="show = !show"
      style="width: 100%"
    />
    <v-btn @click="login" style="width: 100%">LOGIN</v-btn>
    <v-btn @click="$router.push({ name: 'Register' })" style="width: 100%"
      >REGISTER</v-btn
    >
  </div>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  name: "Login",
  computed: {
    ...mapState(["host"]),
  },
  data() {
    return {
      show: false,
      username: "",
      password: "",
      rules: {
        required: (value) => !!value || "Required.",
        min: (v) => v.length >= 1 || "Min 1 characters", //issues#4
      },
    };
  },
  methods: {
    login() {
      axios
        .post(this.host + "/user/login", {
          username: this.username,
          password: this.password,
          output: "json",
        })
        .then((res) => {
          localStorage.setItem("token", res.data.token);
          this.$store.commit("setToken", res.data.token);
          localStorage.setItem("type", res.data.type);
          this.$store.commit("setType", res.data.type);
          localStorage.setItem("username", this.username);
          this.$store.commit("setUsername", this.username);
          this.$router.push("/");
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
          return error;
        });
    },
  },
  created() {
    this.$store.commit("logout");
  },
};
</script>
