package main

import (
	"time"

	"github.com/AnneYang720/DataBase2/App"
	middleware "github.com/AnneYang720/DataBase2/Middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

const (
	MongoDBHosts = "127.0.0.1:27017"
	AuthDatabase = "admin"
	AuthUserName = "xyang"
	AuthPassword = "GufKTQQYN3o2"
	MaxCon       = 300
)

func main() {
	//连接数据库
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	App.MyUserModel = &App.UserModel{
		DB: session.DB("database02").C("users"),
	}

	App.MyExpTypeModel = &App.ExpTypeModel{
		DB: session.DB("database02").C("exp_type"),
	}

	App.MyExpCaseModel = &App.ExpCaseModel{
		DB: session.DB("database02").C("exp_case"),
	}

	App.MyExpDataModel = &App.ExpDataModel{
		DB: session.DB("database02").C("exp_data"),
	}

	App.MyReactorModel = &App.ReactorModel{
		DB: session.DB("database02").C("reactor_mod"),
	}

	App.MyWConditionModel = &App.WConditionModel{
		DB: session.DB("database02").C("working_con"),
	}

	App.MyGridFS = session.DB("database02").GridFS("fs")

	App.MySimCaseModel = &App.SimCaseModel{
		DB: session.DB("database02").C("sim_case"),
	}

	App.MySimOutTypeModel = &App.SimOutTypeModel{
		DB: session.DB("database02").C("sim_outtype"),
	}

	App.MySimDataModel = &App.SimDataModel{
		DB: session.DB("database02").C("sim_data"),
	}

	//开启服务器
	r := gin.Default()

	//设置session信息存储
	c := session.DB("database02").C("sessions")
	store := mongo.NewStore(c, 3600, true, []byte("secret"))
	r.Use(sessions.Sessions("sessionid", store))

	//跨域
	r.Use(middleware.Passjs())

	//注册、登录
	r.POST("/register", App.Register)
	r.POST("/login", App.Login)

	//增加工况、型号等
	r.POST("/addreactormod", App.AddReactorModel)
	r.POST("/addworkingcon", App.AddWorkingCon)

	exp := r.Group("/exp")
	{
		exp.POST("/addtype", App.AddExpType)
		exp.POST("/addcase", App.AddExpCase)
		exp.POST("/adddata", App.AddExpData)
	}

	sim := r.Group("/sim")
	{
		sim.POST("/uploadinput", App.UploadSingleInput)
		sim.GET("/downloadfile/:id", App.DownloadFile)
		sim.POST("/addcase", App.AddSimCase)
		sim.POST("/addoutputtype", App.AddSimOutputType)
		sim.GET("/adddata/:typeid/:caseid/:fileid", App.AddSimData)
		// exp.POST("/addexpdata", App.AddExpData)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.Authorize())
	{
		admin.GET("/info", App.GetUserInfo)
	}

	r.Run(":9999")
}
