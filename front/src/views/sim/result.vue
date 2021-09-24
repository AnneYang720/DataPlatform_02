<template>
  <div>
    
    <br>
    <el-select v-model="chosenSimId" label="选中仿真" style="margin-left:5%;margin-top:2%">
        <el-option v-for="item in simId" :label="item" :key="item" :value="item"/>
    </el-select>
    <el-button style="margin-left:1%" type="primary" @click="handleGetData()">查看仿真数据</el-button>

  <el-table
    :data="expData"
    :row-style="{height:40+'px'}"
    style="width:90%;margin-left:5%;margin-top:2%">
    
      <el-table-column
      v-for="item in tableColumnList" 
      :label="item"
      :property="item"
      :key="item"
      align="center">
      </el-table-column>

   </el-table>
    
  </div>
</template>
<script>

import simApi from '@/api/sim'

export default {
  	data(){
		return {
			simId: [], //该用户所有实验的simId
			chosenSimId: '',
			tableColumnList: [],
			expData: [],
		}
  	},

	created () {
		this.fetchSimId()
		this.getParams()
	},

  	methods: {

		getParams(){
			if(this.$route.query.simId){
				this.chosenSimId = this.$route.query.simId
			}
		},
      
		fetchSimId(){
			simApi.getSimId().then(response =>{
				this.simId = response.Data
			}).catch(() => {
				this.simId = []
			});
		},


      	handleGetData(){
        	if(this.chosenSimId==''){
                this.$message.error('请选择仿真ID')
                return false
            }            
            expApi.getData(this.chosenSimId).then(response =>{
                this.tableColumnList = response.Data.column
				this.expData = response.Data.data
            }).catch(() => {
                this.tableColumnList = []
				this.expData = []
            })
      	},
      
	},
  
	watch: {
		'$route': 'fetchSimId'
	},
}
      
</script>
