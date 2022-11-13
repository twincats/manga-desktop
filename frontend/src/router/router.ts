import { createRouter, createWebHashHistory } from 'vue-router'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/Home.vue'),
    },
    {
      path: '/convert',
      name: 'convert',
      component: () => import('../views/Convert.vue'),
    },
    {
      path: '/chapter/:mid',
      name: 'chapter',
      props: true,
      component: () => import('../views/Chapter.vue'),
    },
    {
      path: '/page/:cid',
      name: 'page',
      props: true,
      component: () => import('../views/Page.vue'),
    },
    {
      path: '/test',
      name: 'test',
      component: () => import('../views/Test.vue'),
    },
  ],
})
