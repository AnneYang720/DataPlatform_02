package App

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func AddExpType(c *gin.Context) {
	//解析post的数据存到postUser内
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	newType := ExpType_notSaved{}
	json.Unmarshal(con, &newType)

	//检查该类型实验是否已存在
	tmpType := ExpType{}
	MyExpTypeModel.DB.Find(bson.M{"typename": newType.TypeName}).One(&tmpType)
	hexid := fmt.Sprintf("%x", string(tmpType.Id))
	if hexid == "" {
		err := MyExpTypeModel.DB.Insert(&newType)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, &ApiResponse{
			Code:    200,
			Flag:    true,
			Message: "New experiment type added.",
		})
	} else {
		c.JSON(http.StatusOK, &ApiResponse{
			Code:    400,
			Flag:    false,
			Message: "This experiment type is existed.",
		})
	}
}
