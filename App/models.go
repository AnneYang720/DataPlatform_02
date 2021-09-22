package App

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	MyUserModel       *UserModel
	MyExpTypeModel    *ExpTypeModel
	MyExpCaseModel    *ExpCaseModel
	MyExpDataModel    *ExpDataModel
	MyReactorModel    *ReactorModel
	MyWConditionModel *WConditionModel
	MyGridFS          *mgo.GridFS
	MySimCaseModel    *SimCaseModel
	MySimOutTypeModel *SimOutTypeModel
	MySimDataModel    *SimDataModel
)

type ApiResponse struct {
	Code    int         `bson:"code"`
	Flag    bool        `bson:"flag"`
	Message interface{} `bson:"message"`
}

type ObjectID struct {
	Id string
}

type UserModel struct {
	DB *mgo.Collection
}

type ExpTypeModel struct {
	DB *mgo.Collection
}

type ExpCaseModel struct {
	DB *mgo.Collection
}

type ExpDataModel struct {
	DB *mgo.Collection
}

type ReactorModel struct {
	DB *mgo.Collection
}

type WConditionModel struct {
	DB *mgo.Collection
}

type SimCaseModel struct {
	DB *mgo.Collection
}

type SimOutTypeModel struct {
	DB *mgo.Collection
}

type SimDataModel struct {
	DB *mgo.Collection
}

//完整的用户属性
type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	Auth     []string      `bson:"auth"`
}

//未注册用户，没有Id等其他属性
type User_notRegistered struct {
	Username string   `bson:"username"`
	Password string   `bson:"password"`
	Auth     []string `bson:"auth"`
}

//完整的实验类型属性
type ExpType struct {
	Id       bson.ObjectId `bson:"_id"`
	TypeName string        `bson:"typename"`
	Para     []string      `bson:"para"`
	Unit     []string      `bson:"unit"`
}

//未保存实验类型，没有Id等其他属性
type ExpType_notSaved struct {
	TypeName string   `bson:"typename"`
	Para     []string `bson:"para"`
	Unit     []string `bson:"unit"`
}

// TODO 实验报告
//完整的实验名称属性
type ExpCase struct {
	Id      bson.ObjectId `bson:"_id"`
	Type    bson.ObjectId `bson:"type"`
	ExpName string        `bson:"expname"`
	ModId   bson.ObjectId `bson:"modid"`
	Time    time.Time     `bson:"time"`
}

//未保存实验名称，没有Id等其他属性
type ExpCase_notSaved struct {
	Type    bson.ObjectId `bson:"type"`
	ExpName string        `bson:"expname"`
	ModId   bson.ObjectId `bson:"modid"`
	Time    time.Time     `bson:"time"`
}

// TODO map[string]interface
// TODO 想为空 可以omitempty
//实验数据，没有Id等其他属性
type ExpData_toSave struct {
	LabID bson.ObjectId      `bson:"labid"`
	Data  map[string]float32 `bson:"data"`
}

//未保存实验数据，没有Id等其他属性
type ExpData_notSaved struct {
	LabID bson.ObjectId `bson:"labid"`
	Data  [][]float32   `bson:"data"`
}

//完整的型号
type ReactorMod struct {
	Id          bson.ObjectId `bson:"_id"`
	ModName     string        `bson:"modname"`
	Description string        `bson:"description"`
}

//未保存型号，没有Id等其他属性
type ReactorMod_notSaved struct {
	ModName     string `bson:"modname"`
	Description string `bson:"description"`
}

//完整的工况
type WorkingCon struct {
	Id      bson.ObjectId `bson:"_id"`
	ConName string        `bson:"conname"`
}

//未保存工况，没有Id等其他属性
type WorkingCon_notSaved struct {
	ConName string `bson:"conname"`
}

//未保存输入文件系列，没有Id等其他属性
type SimuInputs_notSaved struct {
	Boundary  bson.Binary `bson:"boundary.dat"`
	CoreInput bson.Binary `bson:"CoreInput.txt"`
}

// TODO 确定寿期是谁的属性
//完整的仿真实例属性
type SimCase struct {
	Id      bson.ObjectId   `bson:"_id"`
	SimDate time.Time       `bson:"simtime"`
	WConID  bson.ObjectId   `bson:"wconid"`
	RModID  bson.ObjectId   `bson:"rmodid"`
	Inputs  []bson.ObjectId `bson:"inputs"`
	Outputs []bson.ObjectId `bson:"outputs"`
}

//未保存仿真实例，没有Id等其他属性
type SimCase_notSaved struct {
	SimTime time.Time       `bson:"simtime"`
	WConID  bson.ObjectId   `bson:"wconid"`
	RModID  bson.ObjectId   `bson:"rmodid"`
	Inputs  []bson.ObjectId `bson:"inputs"`
	Outputs []bson.ObjectId `bson:"outputs"`
}

//完整的仿真输出文件类型属性
type SimOutputType struct {
	Id       bson.ObjectId `bson:"_id"`
	TypeName string        `bson:"typename"`
	Para     []string      `bson:"para"`
	Desp     []string      `bson:"desp"`
}

//未保存仿真输出文件类型，没有Id等其他属性
type SimOutputType_notSaved struct {
	TypeName string   `bson:"typename"`
	Para     []string `bson:"para"`
	Desp     []string `bson:"desp"`
}

//未保存仿真数据，没有Id等其他属性
type SimData_toSaved struct {
	SimCaseID bson.ObjectId          `bson:"simcaseid"`
	Data      map[string]interface{} `bson:"data"`
}

// TODO map里面类型不一样
