import { createRouter, createWebHashHistory } from "vue-router";
import AddTask from "../components/AddTask.vue";

const routers = [
   {
      // add task
      path: '/task',
      name: 'task',
      component: AddTask,
   },
];

const router = createRouter({
   history: createWebHashHistory(),
   routers,
});

export default router;