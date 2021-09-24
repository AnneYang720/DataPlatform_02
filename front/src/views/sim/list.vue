<template>
  <div>
    <br>

    <el-table
      :data="simList"
      :row-style="{height:40+'px'}"
      style="width:90%;margin-left:5%;margin-top:2%">

      <el-table-column
        prop="Id"
        label="仿真ID"
        min-width="20%">
        <template slot-scope="scope">
          <span style="color: DodgerBlue">{{ scope.row.Id }}</span>
        </template>
      </el-table-column>

      <el-table-column
        prop="SimName"
        label="仿真名称"
        min-width="20%">
      </el-table-column>

	<el-table-column
        prop="SimTime"
        label="仿真时间"
        min-width="20%">
      </el-table-column>

	  <el-table-column
        prop="RModName"
        label="型号"
        min-width="15%">
      </el-table-column>

	  <el-table-column
        prop="WConName"
        label="工况"
        min-width="10%">
      </el-table-column>

      <el-table-column
        prop="InputsName"
        label="输入文件"
		:formatter="formatList"
        min-width="20%">
      </el-table-column>

	  <el-table-column
        prop="OutputsName"
        label="输出文件"
		:formatter="formatList"
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
import simApi from '@/api/sim'

export default {
    data(){
        return {
			simList: [], //首页用包名得到的列表
        }
    },
    created () {
        this.fetchExpList()
    },
    methods: {
        fetchExpList(){
			simApi.getSimList().then(response =>{
				this.simList = response.Data;
			}).catch((err) => {
				this.simList = []
			});
        },

		toResult(row){
			this.$router.push({
				path:'/sim/result',
				query: {expId:row.Id}
			})
		},

		formatList(row, column) {
          return row[column.property].join(", ")
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