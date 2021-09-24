<template>
  <div>
    <br>

    <el-table
      :data="expList"
      :row-style="{height:40+'px'}"
      style="width:90%;margin-left:5%;margin-top:2%">

      <el-table-column
        prop="Id"
        label="实验ID"
        min-width="20%">
        <template slot-scope="scope">
          <span style="color: DodgerBlue">{{ scope.row.Id }}</span>
        </template>
      </el-table-column>

      <el-table-column
        prop="ExpName"
        label="实验名称"
        min-width="20%">
      </el-table-column>

      <el-table-column
        prop="TypeName"
        label="实验类型"
        min-width="20%">
      </el-table-column>

      <el-table-column
        prop="Time"
        label="实验时间"
        min-width="20%">
      </el-table-column>

      <el-table-column
        label="结果"
        min-width="6%">
        <template slot-scope="scope">
          <el-button @click.native.stop="toResult(scope.row)" type="text" size="small">查看</el-button>
        </template>
      </el-table-column>
      
    </el-table>
    

  </div>



</template>
<script>
import expApi from '@/api/exp'

export default {
    data(){
        return {
          expList: [], //首页用包名得到的列表
        }
    },
    created () {
        this.fetchExpList()
    },
    methods: {
        fetchExpList(){
            expApi.getExpList().then(response =>{
                this.expList = response.Data;
            }).catch((err) => {
                this.expList = []
          });
        },

        toResult(row){
          this.$router.push({
            path:'/exp/result',
            query: {expId:row.Id}
          })
        },
    },
    watch: {
      '$route': 'fetchTasksList'
    },
}
      
</script>

<style rel="stylesheet/scss" lang="scss">
.el-dialog {
  // // transform: translateY(-50%);
  // //border-radius: 10px;
  // // width: 500px;
  // // height: 500px!important;
  .el-dialog__header{  
    background: #f7f7f7;
    text-align: left;   
    font-weight: 600;
  }
}
.el-table--enable-row-hover .el-table__body tr:hover>td{
	background-color: rgba(185,211,249,0.75);
}
</style>