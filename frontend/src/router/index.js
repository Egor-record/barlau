import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/components/home';
import Generator from '@/components/generator';
import Checker from '@/components/checker';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/generator',
      name: 'generator',
      component: Generator,
    },
    {
      path: '/checker',
      name: 'checker',
      component: Checker,
    },
  ],
});
