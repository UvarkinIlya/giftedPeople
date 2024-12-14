import { createRouter, createWebHashHistory } from "vue-router";
import Home from "../views/Home.vue";
import Questionnaire from "../views/Questionnaire.vue";

const routes = [{
    path: '/admin',
    component: Home
  }, {
    path: '/',
    component: Questionnaire
  }]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
