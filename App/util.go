package App

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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
