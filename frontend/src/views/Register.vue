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
    <v-text-field v-model="email" label="Email" style="width: 100%" />
    <v-text-field
      v-model="password"
      :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
      :rules="[rules.required, rules.min]"
      :type="show ? 'text' : 'password'"
      label="Password"
      hint="At least 1 characters"
      @click:append="show = !show"
      style="width: 100%"
    />
    <!-- issues #4 -->
    <v-text-field v-model="name" label="Name" style="width: 100%" />
    <v-text-field
      v-model="twitchUsername"
      :rules="[rules.required]"
      label="Twitch Username"
      style="width: 100%"
    />
    <v-radio-group v-model="type" style="width: 30%">
      <div
        style="
          width: 100%;
          display: flex;
          align-items: flex-start;
          justify-content: space-between;
        "
      >
        <v-radio
          v-for="type in types"
          :key="type.id"
          :label="type"
          :value="type"
        />
      </div>
    </v-radio-group>
    <v-btn @click="register" style="width: 100%">REGISTER</v-btn>
    <v-btn @click="$router.push({ name: 'Login' })" style="width: 100%"
      >LOGIN</v-btn
    >
    <v-btn @click="createExampleUsers" style="width: 100%">Create Example Users</v-btn>
  </div>
</template>

<script>
import axios from "axios";
import {mapState} from 'vuex'
export default {
  name: "Register",
    computed: {
    ...mapState(["host"]),
  },
  data() {
    return {
      username: "",
      email: "",
      password: "",
      name: "",
      twitchUsername: "",
      type: "",
      types: ["streamer", "client"],
      show: false,
      rules: {
        required: (value) => !!value || "Required.",
        min: (v) => v.length >= 1 || "Min 1 characters", //issues#4
      },
    };
  },
  methods: {
    register() {
      axios
        .post(this.host + "/user/register", {
          username: this.username,
          email: this.email,
          password: this.password,
          name: this.name,
          twitchUsername: this.twitchUsername,
          type: this.type,
          output: "json",
        })
        .then(() => {
          this.$router.push("/login");
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },
    createExampleUsers() {
      axios
        .get(this.host + "/user/createExampleUsers", {
          output: "json",
        })
        .then(() => {
          this.$router.push("/login");
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },
  },
};
</script>
