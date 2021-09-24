import request from '@/utils/request'
export default{
    getExpList(){
        return request({
            url: `/exp/getlist`, 
            method: 'get'
        });
    },

    getExpId(){
        return request({
            url: `/exp/getidlist`,//ES6写法
            method: 'get'
        });
    },
    
    getData(expId){
        return request({
            url: `/exp/getdata/${expId}`,//ES6写法
            method: 'get'
        });
    },
    
    getExpTypeList(){
        return request({
            url: `/exp/gettypelist`,
            method: 'get'
        });
    },
    
    createExp(pojo){
        return request({
            url: `/exp/createcase`,
            method: 'post',
            data: pojo
        })
    },


    getModList(){
        return request({
            url: `/getmodlist`,
            method: 'get'
        });
    },
} 
