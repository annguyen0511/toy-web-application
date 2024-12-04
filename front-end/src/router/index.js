import { createRouter, createMemoryHistory } from "vue-router";
import Home from '../views/Home.vue';
import SignIn from '../views/SignIn.vue';
import SignUp from '../views/SignUp.vue';
const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/sign-in",
    name: "signIn",
    component: SignIn
  },
  {
    path: "/sign-up",
    name: "signUp",
    component: SignUp
  }
];

const router = createRouter({
  history: createMemoryHistory(),
  routes
});

export default router;