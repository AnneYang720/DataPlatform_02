<template>
  <div>
    
    <br>
    <el-select v-model="chosenExpId" label="选中实验" style="margin-left:5%;margin-top:2%">
        <el-option v-for="item in expId" :label="item" :key="item" :value="item"/>
    </el-select>
    <el-button style="margin-left:1%" type="primary" @click="handleGetData()">查看实验数据</el-button>

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

import expApi from '@/api/exp'

export default {
  	data(){
		return {
			expId: [], //该用户所有实验的expId
			chosenExpId: '',
			tableColumnList: [],
			expData: [],
		}
  	},

	created () {
		this.fetchExpId()
		this.getParams()
	},

  	methods: {

		getParams(){
			if(this.$route.query.expId){
				this.chosenExpId = this.$route.query.expId
			}
		},
      
		fetchExpId(){
			expApi.getExpId().then(response =>{
				this.expId = response.Data
			}).catch(() => {
				this.expId = []
			});
		},


      	handleGetData(){
        	if(this.chosenExpId==''){
                this.$message.error('请选择实验ID')
                return false
            }            
            expApi.getData(this.chosenExpId).then(response =>{
                this.tableColumnList = response.Data.column
				this.expData = response.Data.data
            }).catch(() => {
                this.tableColumnList = []
				this.expData = []
            })
      	},
      
	},
  
	watch: {
		'$route': 'fetchExpId'
	},
}
      
</script>
