import { createRouter, createWebHistory } from "vue-router";
import { checkToken } from "../helper/auth";
const routes = [
  {
    path: "/",
    name: "root",
    component: () => import("../components/Chat.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/register",
    name: "register",
    component: () => import("../components/pages/Register.vue"),
    meta: { requiresVisitor: true },
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem("token")) {
        next({ name: "root" });
      } else next();
    },
  },
  {
    path: "/login",
    name: "login",
    component: () => import("../components/pages/Login.vue"),
    meta: { requiresVisitor: true },
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem("token")) {
        next({ name: "root" });
      } else next();
    },
  },
];

const router = createRouter({
  routes,
  history: createWebHistory(),
});

router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page.
    if (localStorage.getItem("token")) {
      checkToken()
        .then((res) => {
          if (res.status === 200) next();
          else if (res.status === 401 || res.status === 500) next({ name: "login" });
          else next();
        })
        .catch((err) => {
          next({ name: "login" });
        });
      return;
    } else {
      next({
        name: "login",
      });
    }
  } else {
    next(); // make sure to always call next()!
  }
});

export default router;
