import { createApp } from "vue";
import App from "./App.vue";
import BootstrapVueNext from "bootstrap-vue-next";
import "./assets/styles/styles.css";

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue-next/dist/bootstrap-vue-next.css";

const app = createApp(App);

app.use(BootstrapVueNext);

app.mount("#app");
