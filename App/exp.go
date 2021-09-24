package App

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func GetExpList(c *gin.Context) {
	userId, _ := c.Get("G_userId")

	tmpExpCase := []ExpCase_full{}
	MyExpCaseModel.DB.Find(bson.M{"user": bson.ObjectIdHex(userId.(string))}).All(&tmpExpCase)

	for i := 0; i < len(tmpExpCase); i++ {
		tmpExpType := ExpType{}
		MyExpTypeModel.DB.FindId(tmpExpCase[i].Type).One(&tmpExpType)
		fmt.Printf("%v", "tmpExpType: ")
		fmt.Printf("%v", tmpExpType.TypeName)
		tmpExpCase[i].TypeName = tmpExpType.TypeName
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Get Exp List",
		Data:    tmpExpCase,
	})
}

func GetExpIdList(c *gin.Context) {
	userId, _ := c.Get("G_userId")

	tmpExpCase := []ExpCase{}
	MyExpCaseModel.DB.Find(bson.M{"user": bson.ObjectIdHex(userId.(string))}).All(&tmpExpCase)

	Ids := []string{}

	for i := 0; i < len(tmpExpCase); i++ {
		Ids = append(Ids, (tmpExpCase[i].Id.Hex()))
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Get Exp Id List",
		Data:    Ids,
	})
}

func GetData(c *gin.Context) {
	expcase := c.Param("expid")

	tmpExpCase := ExpCase{}
	MyExpCaseModel.DB.FindId(bson.ObjectIdHex(expcase)).One(&tmpExpCase)

	tmpExpType := ExpType{}
	MyExpTypeModel.DB.FindId(tmpExpCase.Type).One(&tmpExpType)

	// Data Result
	type DataResult struct {
		Column []string             `json:"column"`
		Data   []map[string]float32 `json:"data"`
	}

	tmpExpData := []ExpData{}
	MyExpDataModel.DB.Find(bson.M{"labid": bson.ObjectIdHex(expcase)}).All(&tmpExpData)

	var tmpData []map[string]float32

	for i := 0; i < len(tmpExpData); i++ {
		tmpData = append(tmpData, (tmpExpData[i].Data))
	}

	result := &DataResult{Column: tmpExpType.Para, Data: tmpData}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Get Exp Data",
		Data:    result,
	})
}

func GetExpTypeList(c *gin.Context) {
	// Data Result
	type TypeExtract struct {
		Id       bson.ObjectId `bson:"_id"`
		TypeName string        `bson:"typename"`
	}

	tmpExpType := []TypeExtract{}
	MyExpTypeModel.DB.Find(nil).All(&tmpExpType)

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Get Exp Type List",
		Data:    tmpExpType,
	})
}

func AddExpCaseFromFile(c *gin.Context) {
	//解析post的数据
	userId, _ := c.Get("G_userId")
	file, _ := c.FormFile("exp_file")
	expType := c.PostForm("exp_type")
	expMod := c.PostForm("exp_mod")
	expName := c.PostForm("exp_name")

	fmt.Printf("open file %v", file.Filename)
	fmt.Printf("exp type %v", expType)
	fmt.Printf("exp name %v", expName)

	// 插入新exp case
	i := bson.NewObjectId()
	newCase := ExpCase{Id: i,
		Type:    bson.ObjectIdHex(expType),
		ExpName: expName,
		ModId:   bson.ObjectIdHex(expMod),
		Time:    time.Now(),
		User:    bson.ObjectIdHex(userId.(string)),
	}
	err := MyExpCaseModel.DB.Insert(&newCase)
	if err != nil {
		panic(err)
	}

	// 读数据
	fileContent, _ := file.Open()
	csvLines, _ := csv.NewReader(fileContent).ReadAll()

	tmpType := ExpType{}
	MyExpTypeModel.DB.FindId(bson.ObjectIdHex(expType)).One(&tmpType)

	for _, line := range csvLines[1:] {
		tmpData := ExpData_toSave{}
		tmpData.Data = make(map[string]float32)
		tmpData.LabID = i

		for i := 0; i < len(tmpType.Para); i++ {
			value, _ := strconv.ParseFloat(line[i], 32)
			tmpData.Data[tmpType.Para[i]] = float32(value)
		}

		err := MyExpDataModel.DB.Insert(&tmpData)
		if err != nil {
			panic(err)
		}
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Create new exp case and save data",
	})
}
