<template>
  <v-app>
    <v-main>
      <v-card>
        <v-toolbar color="purple darken-4" dark flat>
          <v-toolbar-title @click="$router.push({ name: 'Home' })"
            >Eight-Sup</v-toolbar-title
          >
          <v-spacer></v-spacer>
          <v-btn
            v-if="type==='client'"
            icon
            @click="$router.push({ name: 'WantPropClient' })"
            ><!-- it's reachable when login-->
            <v-icon>mdi-charity</v-icon>
          </v-btn>
          <v-btn
            v-if="type==='streamer'"
            icon
            @click="$router.push({ name: 'WantPropStreamer' })"
            ><!-- it's reachable when login-->
            <v-icon>mdi-charity</v-icon>
          </v-btn>
          <v-btn
            v-if="type==='streamer'"
            icon
            @click="$router.push({ name: 'ListClient' })"
            ><!-- it's reachable when login-->
            <v-icon>mdi-format-list-bulleted</v-icon>
          </v-btn>
          <v-btn
            v-if="type==='client'"
            icon
            @click="$router.push({ name: 'ListStreamer' })"
            ><!-- it's reachable when login-->
            <v-icon>mdi-format-list-bulleted</v-icon>
          </v-btn>
          <v-btn
            v-if="!!token"
            icon
            @click="$router.push({ name: 'History' })"
            ><!-- it's reachable when login-->
            <v-icon>mdi-history</v-icon>
          </v-btn>
          <v-btn
            v-if="type==='client'"
            icon
            @click="$router.push({ name: 'Package' })"
            ><!-- it's reachable when client login-->
            <v-icon>mdi-magnify</v-icon>
          </v-btn>
          <v-btn
            v-if="type==='streamer'"
            icon
            @click="$router.push({ name: 'Plan' })"
            ><!-- it's reachable when streamer login-->
            <v-icon>mdi-plus</v-icon>
          </v-btn>
          <v-btn icon @click="$router.push({ name: 'Login' })">
            <v-icon v-if="!!token">mdi-account-remove</v-icon>
            <v-icon v-else>mdi-account</v-icon>
          </v-btn>
        </v-toolbar>
      </v-card>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "App",
    computed: {
    ...mapState(["token","type"]),
  },
  created() {
    const token = localStorage.getItem("token");
    const username = localStorage.getItem("username");
    const type = localStorage.getItem("type");
    if (token) this.$store.commit("setToken", token);
    if (username) this.$store.commit("setUsername", username);
    if (type) this.$store.commit("setType", type);
  },
};
</script>
