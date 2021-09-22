package App

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func AddExpType(c *gin.Context) {
	//解析post的数据
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

func AddExpCase(c *gin.Context) {
	//解析post的数据
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	newCase := ExpCase_notSaved{Time: time.Now()}
	json.Unmarshal(con, &newCase)

	err := MyExpCaseModel.DB.Insert(&newCase)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "New experiment case added.",
	})
}

func AddExpData(c *gin.Context) {
	//解析post的数据
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	newData := ExpData_notSaved{}
	json.Unmarshal(con, &newData)

	tmpCase := ExpCase{}
	MyExpCaseModel.DB.FindId(newData.LabID).One(&tmpCase)

	tmpType := ExpType{}
	MyExpTypeModel.DB.FindId(tmpCase.Type).One(&tmpType)

	for i := 0; i < len(newData.Data); i++ {
		tmpData := ExpData_toSave{}
		tmpData.Data = make(map[string]float32)
		tmpData.LabID = newData.LabID
		for j := 0; j < len(tmpType.Para); j++ {
			fmt.Printf("Element[%v] = %v\n", tmpType.Para[i], newData.Data[i][j])
			tmpData.Data[tmpType.Para[j]] = newData.Data[i][j]
		}
		err := MyExpDataModel.DB.Insert(&tmpData)
		if err != nil {
			panic(err)
		}
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "New data added.",
	})
}
