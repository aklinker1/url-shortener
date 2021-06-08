import './framework.css';
import { createApp } from "vue";
import App from "./App.vue";
import "./firebase/init";
import router from './router';

createApp(App).use(router).mount("#app");
