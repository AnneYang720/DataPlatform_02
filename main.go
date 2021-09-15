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

	exp := r.Group("/exp")
	{
		exp.POST("/addexptype", App.AddExpType)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.Authorize())
	{
		admin.GET("/info", App.GetUserInfo)
	}

	// article := r.Group("/article")
	// {
	// 	article.GET("/", App.GetPage)

	// 	article.GET("/all", App.GetAllArticles)

	// 	article.GET("/aid/:aid", App.GetArticleByAid)

	// 	article.GET("/title/:title", App.GetArticlesByTitle)

	// 	article.GET("/tag/:tag", App.GetArticlesByTag)

	// 	article.GET("/publisher/:publisher", App.GetArticlesByPublisher)

	// 	//article.Use(Authorize())

	// 	article.POST("/publish", App.PublishArticle)

	// 	article.DELETE("/:aid", App.DeleteArticleByAid)

	// 	article.PUT("/:aid", App.ModifyArticleByAid)

	// }

	// comment := r.Group("/comment")
	// {
	// 	comment.GET("/id/:id", App.GetCommentsById)

	// 	//comment.Use(Authorize())

	// 	comment.POST("/publish", App.AddComment)

	// 	comment.PUT("/:id", App.ModifyCommentByCid)

	// 	comment.DELETE("/:id", App.DeleteCommentByCid)

	// }

	// like := r.Group("/like")
	// {
	// 	like.GET("/id/:id", App.GetLikesById)
	// 	//like.Use(Authorize())

	// 	like.POST("/likeit", App.LikeIt)

	// 	like.DELETE("/:lid", App.UnlikeIt)

	// }

	r.Run(":9999")
}
