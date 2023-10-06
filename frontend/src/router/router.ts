import { createRouter, createWebHashHistory } from 'vue-router'
import { typedProps } from './param'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/Home.vue'),
    },
    {
      path: '/convert/:mid?',
      name: 'convert',
      props: typedProps({ mid: Number }),
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
      path: '/download',
      name: 'download',
      component: () => import('../views/Download.vue'),
    },
    {
      path: '/test',
      name: 'test',
      component: () => import('../views/Test.vue'),
    },
    {
      path: '/dlcode',
      name: 'dlcode',
      component: () => import('../views/admin/DownloadCodeEditor.vue'),
    },
    {
      path: '/addalter/:mid',
      name: 'addalter',
      props: typedProps({ mid: Number }),
      component: () => import('../views/admin/AddAlter.vue'),
    },
    {
      path: '/serverstat',
      name: 'serverstat',
      component: () => import('../views/admin/ServerStat.vue'),
    },
    {
      path: '/server/add',
      name: 'serveradd',
      component: () => import('../views/admin/AddServer.vue'),
    },
  ],
})
