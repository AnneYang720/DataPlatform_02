<template>
  <div>
    <br>
      <div style="line-height:60px;margin-left:5%;font-size:20px"> 新建仿真 </div>
      
      <el-form label-width="100px" style="margin-left:5%">  

        <el-form-item label="型号">
          <el-select v-model="chosenMod" >
            <el-option v-for="item in modList" :label="item.ModName" :key="item.Id" :value="item.Id"/>
          </el-select>
        </el-form-item>
        
        <el-form-item label="仿真工况">
          <el-select v-model="chosenCon" >
            <el-option v-for="item in conList" :label="item.ConName" :key="item.Id" :value="item.Id"/>
          </el-select>
        </el-form-item>

		<el-form-item label="仿真名称">
          <el-col :span="8">
            <el-input v-model="simName" placeholder="选填写仿真名称"></el-input>
          </el-col>
        </el-form-item>

        <el-form-item label="输入文件">
          <el-col :span="4">
            <el-upload
              ref="uploadInputFile"
              action="a"
              multiple
              :auto-upload="false"
              :show-file-list="true"
              :file-list="uploadInputList"
              :on-change="handleInputUploadChange">
              <el-button type="primary" slot="trigger">选取文件</el-button>
            </el-upload>
          </el-col>
        </el-form-item>

		<el-form-item label="输出文件">
          <el-col :span="4">
            <el-upload
              ref="uploadOutputFile"
              action="a"
              multiple
              :auto-upload="false"
              :show-file-list="true"
              :file-list="uploadOutputList"
              :on-change="handleOutputUploadChange">
              <el-button type="primary" slot="trigger">选取文件</el-button>
            </el-upload>
          </el-col>
        </el-form-item>
        

      </el-form>

      <el-button style="margin-left:5%" type="primary"  :loading="loading" @click="handleCreate()">创 建</el-button>


  </div>
</template>
<script>

import expApi from '@/api/exp'
import simApi from '@/api/sim'

export default {
    data(){
        return {
			modList: [],
			conList: [],
			simName: '',
			chosenMod: '',
			chosenCon: '',
			uploadInputList: [],
			uploadOutputList: [],

			loading: false,
        }
    },
    created () {
        this.fetchList()
    },
    methods: {
        // 获取型号和工况
        fetchList(){
			simApi.getConList().then(response =>{
				this.conList = response.Data
			}).catch(() => {
				this.conList = []
			});
			expApi.getModList().then(response =>{
				this.modList = response.Data
			}).catch(() => {
				this.modList = []
			});
        },
        
        beforeCreate(){
			if(this.chosenMod==''){
				this.$message.error('请选择型号')
				return false
			}
			if(this.chosenCon==''){
				this.$message.error('请选择工况')
				return false
			}
			if(this.simName==''){
				this.$message.error('请填写仿真名称')
				return false
			}
			if(this.uploadInputList.length==0){
				this.$message.error('仿真输入文件必须上传')
				return false
			}
			if(this.uploadOutputList.length==0){
				this.$message.error('仿真输出文件必须上传')
				return false
			}
			return true
        },

        // Final
        handleCreate(){
			this.$confirm('确认创建该仿真', '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			}).then(async() => {
				if(this.beforeCreate()){
					this.loading = true
					let formData = new FormData();

					this.uploadInputList.forEach((val,index) => {
						// inputfile.push(val.raw)
						formData.append('sim_inputfile', val.raw)
					})

					this.uploadOutputList.forEach((val,index) => {
						// outputfile.push(val.raw)
						formData.append('sim_outputfile', val.raw)
					})

					formData.append('sim_con', this.chosenCon)
					formData.append('sim_mod', this.chosenMod)
					formData.append('sim_name', this.simName)

					console.log(formData)

					simApi.createSim(formData).then(async(response) => {
						if(response.Flag){//如果成功
							this.$message.success('新任务创建成功')
							this.closeCreate() // 清空前端用户上传数据
						}else{
							this.$message.error('新任务创建失败')
						}
					})
							
					this.loading = false
				}
			})
        },

        closeCreate(){
			this.uploadInputList = []
			this.uploadOutputList = []
			this.chosenMod = ''
			this.chosenCon = ''
			this.simName = ''
        },

		// 文件上传，文件不能重复
        handleInputUploadChange(file, uploadInputList) {
            let existFile = uploadInputList.slice(0, uploadInputList.length - 1).find(f => f.name === file.name);
                if (existFile) {
                    console.log('当前文件已经存在!');
                    this.$message({
                      message: '当前文件已经存在! 新选中文件未上传',
                      type: 'error'
                    })
                    uploadInputList.pop();
                }
            this.uploadInputList = uploadInputList
        },
        handleOutputUploadChange(file, uploadOutputList) {
            let existFile = uploadOutputList.slice(0, uploadOutputList.length - 1).find(f => f.name === file.name);
                if (existFile) {
                    console.log('当前文件已经存在!');
                    this.$message({
                      message: '当前文件已经存在! 新选中文件未上传',
                      type: 'error'
                    })
                    uploadOutputList.pop();
                }
            this.uploadOutputList = uploadOutputList
        },

    }
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

.el-tag + .el-tag {
  margin-left: 10px;
}
.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}

</style>