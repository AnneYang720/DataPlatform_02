import Vue from 'vue'
import Router from 'vue-router'

// in development-env not use lazy-loading, because lazy-loading too many pages will cause webpack hot update too slow. so only in production use lazy-loading;
// detail: https://panjiachen.github.io/vue-element-admin-site/#/lazy-loading

Vue.use(Router)

/* Layout */
import Layout from '../views/layout/Layout'

/**
* hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
* alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
*                                if not set alwaysShow, only more than one route under the children
*                                it will becomes nested mode, otherwise not show the root menu
* redirect: noredirect           if `redirect:noredirect` will no redirct in the breadcrumb
* name:'router-name'             the name is used by <keep-alive> (must set!!!)
* meta : {
    title: 'title'               the name show in submenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar,
  }
**/
export const constantRouterMap = [
  { path: '/login', component: () => import('@/views/login/index'), hidden: true },
  { path: '/register', component: () => import('@/views/register/index'), hidden: true },
  { path: '/404', component: () => import('@/views/404'), hidden: true },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    name: 'Dashboard',
    hidden: true,
    children: [{
      path: 'dashboard',
      component: () => import('@/views/dashboard/index')
    }]
  },

  {
    path: '/exp',
    component: Layout,
    name: 'Experiment',
    meta: { title: '实验', icon: 'example' },
    children: [
      {
        path: 'create',
        name: 'CreateExp',
        component: () => import('@/views/exp/create'),
        meta: { title: '新建实验', icon: 'table', reuse: false }
      },
      {
        path: 'list',
        name: 'ExpList',
        component: () => import('@/views/exp/list'),
        meta: { title: '实验列表', icon: 'table', reuse: false  }
      },
      {
        path: 'result',
        name: 'ExpResult',
        component: () => import('@/views/exp/result'),
        meta: { title: '实验结果', icon: 'table', reuse: false  }
      }
    ]
  },

  {
    path: '/sim',
    component: Layout,
    name: 'Simulation',
    meta: { title: '仿真', icon: 'example' },
    children: [
      {
        path: 'create',
        name: 'Create',
        component: () => import('@/views/sim/create'),
        meta: { title: '新建仿真', icon: 'table', reuse: false }
      },
      {
        path: 'list',
        name: 'List',
        component: () => import('@/views/sim/list'),
        meta: { title: '仿真列表', icon: 'table', reuse: false  }
      },
      {
        path: 'result',
        name: 'Result',
        component: () => import('@/views/sim/result'),
        meta: { title: '仿真结果', icon: 'table', reuse: false  }
      }
    ]
  },


  { path: '*', redirect: '/404', hidden: true }
]

/**
 * asyncRouterMap
 * the routes that need to be dynamically loaded based on user roles
 */
 export const asyncRouterMap = [
  {
    path: '/permission',
    component: Layout,
    name: 'Permission',
    meta: {roles: ['admin'], title: '管理', icon: 'example'}, // you can set roles in root nav
    children: [
      {
        path: 'hardware',
        component: () => import('@/views/admin/worker'),
        name: 'Hardware',
        meta: { title: '硬件信息', roles: ['admin'] }
      },
      {
        path: 'user',
        component: () => import('@/views/admin/user'),
        name: 'User',
        meta: { title: '用户权限', roles: ['admin'] }
      }
    ]
  },

  { path: '*', redirect: '/404', hidden: true }
]


const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export function selfaddRoutes(params) {
  router.matcher = new Router().matcher;
  router.addRoutes(params)
}

export default router

