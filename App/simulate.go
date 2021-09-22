package App

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func UploadSingleInput(c *gin.Context) {
	//解析post的数据
	file, _ := c.FormFile("file")
	fileContent, _ := file.Open()

	defer fileContent.Close()

	gridFile, _ := MyGridFS.Create(file.Filename)

	_, err := io.Copy(gridFile, fileContent)

	err = gridFile.Close()

	// err := MyReactorModel.DB.Insert(&newReactorMod)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: gridFile.Id(),
	})
}

func DownloadFile(c *gin.Context) {
	id := c.Param("id")
	file, err := MyGridFS.OpenId(bson.ObjectIdHex(id))
	size := file.Size()

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name()))
	c.Header("Content-Type", "application/text/plain")
	c.Header("Accept-Length", fmt.Sprintf("%d", size))

	io.Copy(c.Writer, file)
	err = file.Close()
	c.Writer.WriteHeader(http.StatusOK)

	if err != nil {
		panic(err)
	}
}

func AddSimCase(c *gin.Context) {
	//解析post的数据
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	newCase := SimCase_notSaved{SimTime: time.Now()}
	json.Unmarshal(con, &newCase)

	err := MySimCaseModel.DB.Insert(&newCase)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "New simulation case added.",
	})
}

func AddSimOutputType(c *gin.Context) {
	//解析post的数据
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	newType := SimOutputType_notSaved{}
	json.Unmarshal(con, &newType)

	//检查该类型实验是否已存在
	tmpType := SimOutputType{}
	MySimOutTypeModel.DB.Find(bson.M{"typename": newType.TypeName}).One(&tmpType)
	hexid := fmt.Sprintf("%x", string(tmpType.Id))
	if hexid == "" {
		err := MySimOutTypeModel.DB.Insert(&newType)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, &ApiResponse{
			Code:    200,
			Flag:    true,
			Message: "New simulation output file type added.",
		})
	} else {
		c.JSON(http.StatusOK, &ApiResponse{
			Code:    400,
			Flag:    false,
			Message: "This simulation output file type is existed.",
		})
	}
}

func AddSimData(c *gin.Context) {
	typeid := c.Param("typeid")
	caseid := c.Param("caseid")
	fileid := c.Param("fileid")

	// open output csv file
	file, err := MyGridFS.OpenId(bson.ObjectIdHex(fileid))
	fmt.Printf("open file %v", file.Name())
	csvLines, err := csv.NewReader(file).ReadAll()

	tmpType := SimOutputType{}
	MySimOutTypeModel.DB.FindId(bson.ObjectIdHex(typeid)).One(&tmpType)
	fmt.Printf("type name %v", tmpType.TypeName)

	for _, line := range csvLines[2:] {
		// fmt.Println("time  %v ", line[0])
		tmpData := SimData_toSaved{}
		tmpData.Data = make(map[string]interface{})
		tmpData.SimCaseID = bson.ObjectIdHex(caseid)
		for i := 0; i < len(tmpType.Para); i++ {
			tmpData.Data[tmpType.Para[i]], _ = strconv.ParseFloat(line[i], 32)
		}
		err = MySimDataModel.DB.Insert(&tmpData)
		if err != nil {
			panic(err)
		}
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "New simulation output data successfully added.",
	})
}
