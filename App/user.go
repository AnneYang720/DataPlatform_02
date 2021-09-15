package App

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

//用户注册
func Register(c *gin.Context) {
	//解析post的数据存到postUser内
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	postUser := User_notRegistered{Auth: []string{"user"}}

	json.Unmarshal(con, &postUser)
	postUser.Password = HashPassword(postUser.Password)

	//检查用户名是否已经被注册
	tmpUser := User{}
	MyUserModel.DB.Find(bson.M{"username": postUser.Username}).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
	if hexid == "" {
		err := MyUserModel.DB.Insert(&postUser)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, &ApiResponse{
			Code:    200,
			Flag:    true,
			Message: "register success",
		})
	} else {
		c.JSON(http.StatusOK, &ApiResponse{
			Code:    400,
			Flag:    false,
			Message: "username has existed.",
		})
	}
}

//用户登陆
func Login(c *gin.Context) {
	session := sessions.Default(c)                                                                           //change
	option := sessions.Options{MaxAge: 3600, Path: "/", Domain: "localhost", HttpOnly: false, Secure: false} //change
	session.Options(option)                                                                                  //change

	//解析post的数据存到postUser内
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	postUser := User{}
	json.Unmarshal(con, &postUser)

	//检查用户名和密码是否匹配
	tmpUser := User{}
	MyUserModel.DB.Find(bson.M{"username": postUser.Username}).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))

	check := VerifyPassword(tmpUser.Password, postUser.Password)
	fmt.Printf("%v", check)

	if check {
		session.Set("userid", hexid) //change
		session.Save()               //change
		c.JSON(http.StatusOK, &ApiResponse{
			Code:    200,
			Flag:    true,
			Message: &ObjectID{Id: hexid},
		})
	} else {
		c.JSON(http.StatusOK, &ApiResponse{
			Code:    400,
			Flag:    false,
			Message: "username and password do not match",
		})
	}
}

func GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("G_userId")
	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: userId,
	})
}

// // GetUserByID 根据ID查询用户
// func GetUserByUid (c *gin.Context) {
// 	tmpUser := User{}
// 	MyuserModel.DB.FindId(bson.ObjectIdHex(c.Param("uid"))).One(&tmpUser)
// 	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
// 	if (hexid == "") {
// 		c.JSON(http.StatusOK, &ApiResponse {
// 			Code: 400,
// 			Type: "fail",
// 			Message:  "user id does not exist",
// 		})
// 	} else {
// 		tmpUser.Id = fmt.Sprintf("%x", string(tmpUser.Id))
// 		c.JSON(http.StatusOK, &ApiResponse {
// 			Code: 200,
// 			Type: "success",
// 			Message:  &tmpUser,
// 		})
// 	}
// }

// // 根据用户名查询用户
// func GetUserByUsername (c *gin.Context) {
// 	tmpUser := User{}
// 	MyuserModel.DB.Find(bson.M{"username": c.Param("username")}).One(&tmpUser)
// 	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
// 	if (hexid == "") {
// 		c.JSON(http.StatusOK, &ApiResponse {
// 			Code: 400,
// 			Type: "fail",
// 			Message:  "user does not exist",
// 		})
// 	} else {
// 		tmpUser.Id = fmt.Sprintf("%x", string(tmpUser.Id))
// 		c.JSON(http.StatusOK, &ApiResponse {
// 			Code: 200,
// 			Type: "success",
// 			Message:  &tmpUser,
// 		})
// 	}
// }

// //修改用户信息
// func ModifyUserByUid (c *gin.Context) {
// 	//解析post的数据存到postUser内
// 	con,_ := ioutil.ReadAll(c.Request.Body) //获取post的数据
// 	postUser := User{}
// 	json.Unmarshal(con, &postUser)

// 	tmpUser := User{}
// 	MyuserModel.DB.FindId(bson.ObjectIdHex(c.Param("uid"))).One(&tmpUser)
// 	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
// 	if (hexid == "") {
// 		c.JSON(http.StatusOK, &ApiResponse {
// 			Code: 400,
// 			Type: "fail",
// 			Message:  "user does not exist",
// 		})
// 	} else {
// 		//更新
// 		if (postUser.Username == "") {
// 			postUser.Username = tmpUser.Username
// 		}
// 		if (postUser.Email == "") {
// 			postUser.Email = tmpUser.Email
// 		}
// 		if (postUser.Password == "") {
// 			postUser.Password = tmpUser.Password
// 		}
// 		if (postUser.Phone == "") {
// 			postUser.Phone = tmpUser.Phone
// 		}
// 		MyuserModel.DB.Update(bson.M{"_id": bson.ObjectIdHex(c.Param("uid"))}, bson.M{"$set": bson.M{
// 			"username": postUser.Username,
// 			"email": postUser.Email,
// 			"password": postUser.Password,
// 			"phone": postUser.Phone,
// 		}})
// 		c.JSON(http.StatusOK, &ApiResponse {
// 			Code: 200,
// 			Type: "success",
// 			Message:  "modify user success",
// 		})
// 	}

// }

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	check := true
	if err != nil {
		check = false
	}
	return check
}
