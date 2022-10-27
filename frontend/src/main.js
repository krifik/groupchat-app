import { createApp } from "vue";
// import "./style.css";
import App from "./App.vue";
import "./index.css";
import router from "../src/routes/index.js";
import { createPinia } from "pinia";

createApp(App).use(router).use(createPinia()).mount("#app");
