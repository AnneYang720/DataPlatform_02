// store/permission.js
import { asyncRouterMap, constantRouterMap } from '@/router';

function hasPermission(roles, route) {
    if (route.meta && route.meta.roles) {
        return roles.some(role => route.meta.roles.includes(role))
    } else {
        return true
    }
}

const permission = {
    state: {
        routers: constantRouterMap, //最终的RouteMap
        addRouters: []
    },
    mutations: {
        SET_ROUTERS: (state, routers) => {
            state.addRouters = routers;
            state.routers = constantRouterMap.concat(routers);
        }
    },
    actions: {
        GenerateRoutes({ commit }, data) {
            return new Promise(resolve => {
                const {roles} = data;
                const accessedRouters = asyncRouterMap.filter(v => {
                    if (roles.indexOf('admin') >= 0) return true;
                    if (hasPermission(roles, v)) {
                        if (v.children && v.children.length > 0) {
                            v.children = v.children.filter(child => {
                                if (hasPermission(roles, child)) {
                                    return child;
                                }
                                return false;
                            })
                            return v;
                        } else{
                            return v;
                        }
                    }
                    return false;
                })
                commit('SET_ROUTERS', accessedRouters);
                resolve();
            })
        }
    }
};

export default permission;


