import Vue from "vue";
import VueRouter from "vue-router";
import ListClient from "../views/ListClient.vue";
import ListStreamer from "../views/ListStreamer.vue";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import Register from "../views/Register.vue";
import Package from "../views/Package.vue";
import Plan from "../views/Plan.vue";
import History from "../views/History.vue";
import WantProp from "../views/WantProp.vue";

// import store from '../store'

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/wantProp",
    name: "WantProp",
    component: WantProp,
    auth: true,
  },
  {
    path: "/listClient",
    name: "ListClient",
    component: ListClient,
    auth: true,
  },
  {
    path: "/listStreamer",
    name: "ListStreamer",
    component: ListStreamer,
    auth: true,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
  },
  {
    path: "/package",
    name: "Package",
    component: Package,
    auth: true,
  },
  {
    path: "/plan",
    name: "Plan",
    component: Plan,
    auth: true,
  },
  {
    path: "/history",
    name: "History",
    component: History,
    auth: true,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  const route = routes.find((el) => el.path === to.path);
  if (route.auth && !localStorage.getItem("token")) next("/login");
  else next();
});

export default router;
