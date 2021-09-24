package App

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func AddReactorModel(c *gin.Context) {
	//解析post的数据
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	newReactorMod := ReactorMod_notSaved{}
	json.Unmarshal(con, &newReactorMod)

	err := MyReactorModel.DB.Insert(&newReactorMod)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "New reactor model added.",
	})
}

func AddWorkingCon(c *gin.Context) {
	//解析post的数据
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	newWorkingCon := WorkingCon_notSaved{}
	json.Unmarshal(con, &newWorkingCon)

	err := MyWConditionModel.DB.Insert(&newWorkingCon)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "New working condition added.",
	})
}

func GetModList(c *gin.Context) {
	// Data Result
	type ModExtract struct {
		Id      bson.ObjectId `bson:"_id"`
		ModName string        `bson:"modname"`
	}

	tmpModType := []ModExtract{}
	MyReactorModel.DB.Find(nil).All(&tmpModType)

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Get Mod List",
		Data:    tmpModType,
	})
}

func GetConList(c *gin.Context) {
	tmpConType := []WorkingCon{}
	MyWConditionModel.DB.Find(nil).All(&tmpConType)

	fmt.Printf("name %v ", tmpConType[0])

	fmt.Printf("id %v ", tmpConType[0].Id)
	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Get Con List",
		Data:    tmpConType,
	})
}
