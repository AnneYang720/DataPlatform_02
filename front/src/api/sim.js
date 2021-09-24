import request from '@/utils/request'
export default{
    getSimList(){
        return request({
            url: `/sim/getlist`, 
            method: 'get'
        });
    },

    getSimId(){
        return request({
            url: `/sim/getidlist`,//ES6写法
            method: 'get'
        });
    },
    
    
    createSim(pojo){
        return request({
            url: `/sim/createcase`,
            method: 'post',
            data: pojo
        })
    },

    getConList(){
        return request({
            url: `/getconlist`,
            method: 'get'
        });
    },
} 
