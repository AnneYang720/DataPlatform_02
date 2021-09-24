<template>
  <div>
    <br>
      <div style="line-height:60px;margin-left:5%;font-size:20px"> 新建实验</div>
      
      <el-form label-width="100px" style="margin-left:5%">  

        <el-form-item label="实验类型">
          <el-select v-model="chosenType" >
            <el-option v-for="item in expTypeList" :label="item.TypeName" :key="item.Id" :value="item.Id"/>
          </el-select>
        </el-form-item>

		<el-form-item label="型号">
          <el-select v-model="chosenMod" >
            <el-option v-for="item in modList" :label="item.ModName" :key="item.Id" :value="item.Id"/>
          </el-select>
        </el-form-item>

        <el-form-item label="数据">
          <el-col :span="4">
            <el-upload
              ref="uploadExpFile"
              action="a"
              :multiple="false"
              :auto-upload="false"
              :show-file-list="true"
              :file-list="uploadFileList"
              :on-change="handleFileCreateChange">
              <el-button type="primary" slot="trigger">选取文件</el-button>
            </el-upload>
          </el-col>
        </el-form-item>

        <el-form-item label="实验名称">
          <el-col :span="8">
            <el-input v-model="expName" placeholder="选填写实验名称"></el-input>
          </el-col>
        </el-form-item>

      </el-form>

      <el-button style="margin-left:5%" type="primary"  :loading="loading" @click="handleCreate()">创 建</el-button>


  </div>
</template>
<script>

import expApi from '@/api/exp'

export default {
    data(){
        return {
			expTypeList: [],
			modList: [],
			chosenType: '',
			expName: '',
			uploadFileList: [],
			chosenMod: '',
			loading: false,
        }
    },
    created () {
        this.fetchExpList()
    },
    methods: {
        // 获取所有实验类型
        fetchExpList(){
			expApi.getExpTypeList().then(response =>{
				this.expTypeList = response.Data
			}).catch(() => {
				this.expTypeList = []
			});
			expApi.getModList().then(response =>{
				this.modList = response.Data
			}).catch(() => {
				this.modList = []
			});
        },
        
        // 限制实验数据文件上传的个数只有一个，获取上传列表的最后一个
        handleFileCreateChange(file, uploadFileList) {
            if (uploadFileList.length > 0) {
                this.uploadFileList = [uploadFileList[uploadFileList.length - 1]] // 展示最后一次选择的文件
            }            
        },


        beforeCreate(){
			if(this.chosenType==''){
				this.$message.error('请选择实验类型')
				return false
			}
			if(this.uploadFileList.length==0){
				this.$message.error('实验数据文件必须上传')
				return false
			}
			if(this.expName==''){
				this.$message.error('请填写实验名称')
				return false
			}
			return true
        },


        // Final
        handleCreate(){
			this.$confirm('确认创建该实验', '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			}).then(async() => {
				if(this.beforeCreate()){
					this.loading = true
					let formData = new FormData();
					formData.append('exp_type', this.chosenType)
					formData.append('exp_mod', this.chosenMod)
					formData.append('exp_name', this.expName)
					formData.append('exp_file', this.uploadFileList[0].raw)

					expApi.createExp(formData).then(async(response) => {
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
          this.uploadFileList = []
		  this.chosenType = ''
		  this.expName = ''
		  this.chosenMod = ''
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