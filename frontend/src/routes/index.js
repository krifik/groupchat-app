import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/home",
    name: "home",
    component: () => import("../components/Chat.vue"),
  },
];

const router = createRouter({
  routes,
  history: createWebHistory(),
});

export default router;
