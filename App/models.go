package App

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	MyUserModel    *UserModel
	MyExpTypeModel *ExpTypeModel
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

//完整的用户属性
type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	Auth     []string      `bson:"auth"`
}

//未注册用户没有Id等其他属性
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

//未注册用户没有Id等其他属性
type ExpType_notSaved struct {
	TypeName string   `bson:"typename"`
	Para     []string `bson:"para"`
	Unit     []string `bson:"unit"`
}
