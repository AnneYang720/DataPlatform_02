package App

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

func AddSimCaseFromFile(c *gin.Context) {
	//解析post的数据
	form, _ := c.MultipartForm()
	inputfiles := form.File["sim_inputfile"]
	outputfile := form.File["sim_outputfile"]
	userId, _ := c.Get("G_userId")
	simCon := c.PostForm("sim_con")
	simMod := c.PostForm("sim_mod")
	simName := c.PostForm("sim_name")

	//新建simcase
	simcase_id := bson.NewObjectId()

	//存文件
	inputfilesId := []bson.ObjectId{}
	inputfilesName := []string{}
	for _, file := range inputfiles {
		//创建新的gridFile，并写入文件
		input_id := bson.NewObjectId()
		gridFile, _ := MyGridFS.Create(file.Filename)
		gridFile.SetId(input_id)
		fileContent, _ := file.Open()
		defer fileContent.Close()
		_, err := io.Copy(gridFile, fileContent)
		err = gridFile.Close()
		if err != nil {
			panic(err)
		}
		//存下Id和文件名
		inputfilesName = append(inputfilesName, file.Filename)
		inputfilesId = append(inputfilesId, input_id)
	}
	outputfilesId := []bson.ObjectId{}
	outputfilesName := []string{}
	for _, file := range outputfile {
		//创建新的gridFile，并写入文件
		output_id := bson.NewObjectId()

		//得知输出的类型
		typeName := strings.Split(file.Filename, ".")[0]
		tmpType := SimOutputType{}
		MySimOutTypeModel.DB.Find(bson.M{"typename": typeName}).One(&tmpType)
		typeId := fmt.Sprintf("%x", string(tmpType.Id))
		if typeId == "" {
			panic("输出类型不存在")
		}

		fmt.Printf("output name %v\n", file.Filename)
		fmt.Printf("file type %v\n", tmpType.TypeName)

		//存数据
		fileContent, _ := file.Open()
		csvLines, _ := csv.NewReader(fileContent).ReadAll()
		for _, line := range csvLines[2:] {
			// fmt.Printf("%v\n", line)
			tmpData := SimData_toSaved{}
			tmpData.Data = make(map[string]interface{})
			tmpData.SimCaseID = simcase_id
			tmpData.SimOutputID = output_id
			tmpData.SimOutputTypeID = tmpType.Id
			for i := 0; i < len(tmpType.Para); i++ {
				tmpData.Data[tmpType.Para[i]], _ = strconv.ParseFloat(line[i], 32)
			}
			err := MySimDataModel.DB.Insert(&tmpData)
			if err != nil {
				panic(err)
			}
		}

		//存下文件
		fileContent, _ = file.Open()
		gridFile, _ := MyGridFS.Create(file.Filename)
		gridFile.SetId(output_id)
		_, err := io.Copy(gridFile, fileContent)
		err = gridFile.Close()
		if err != nil {
			panic(err)
		}
		fileContent.Close()

		//存下Id和文件名
		outputfilesName = append(outputfilesName, file.Filename)
		outputfilesId = append(outputfilesId, output_id)
	}

	fmt.Printf("input files %v\n", inputfilesId)
	fmt.Printf("output files %v\n", outputfilesId)

	newCase := SimCase{Id: simcase_id,
		User:        bson.ObjectIdHex(userId.(string)),
		SimName:     simName,
		SimTime:     time.Now(),
		WConID:      bson.ObjectIdHex(simCon),
		RModID:      bson.ObjectIdHex(simMod),
		Inputs:      inputfilesId,
		Outputs:     outputfilesId,
		InputsName:  inputfilesName,
		OutputsName: outputfilesName,
	}

	err := MySimCaseModel.DB.Insert(&newCase)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Create new sim case and save data",
		Data:    outputfilesId,
	})
}

func GetSimList(c *gin.Context) {
	userId, _ := c.Get("G_userId")

	tmpSimCase := []SimCase_Full{}
	MySimCaseModel.DB.Find(bson.M{"user": bson.ObjectIdHex(userId.(string))}).All(&tmpSimCase)

	for i := 0; i < len(tmpSimCase); i++ {
		tmpCon := WorkingCon{}
		MyWConditionModel.DB.FindId(tmpSimCase[i].WConID).One(&tmpCon)
		tmpSimCase[i].WConName = tmpCon.ConName

		tmpMod := ReactorMod{}
		MyReactorModel.DB.FindId(tmpSimCase[i].RModID).One(&tmpMod)
		tmpSimCase[i].RModName = tmpMod.ModName
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Get Sim List",
		Data:    tmpSimCase,
	})
}

func GetSimIdList(c *gin.Context) {
	userId, _ := c.Get("G_userId")

	tmpSimCase := []SimCase{}
	MySimCaseModel.DB.Find(bson.M{"user": bson.ObjectIdHex(userId.(string))}).All(&tmpSimCase)

	Ids := []string{}

	for i := 0; i < len(tmpSimCase); i++ {
		Ids = append(Ids, (tmpSimCase[i].Id.Hex()))
	}

	c.JSON(http.StatusOK, &ApiResponse{
		Code:    200,
		Flag:    true,
		Message: "Get Sim Id List",
		Data:    Ids,
	})
}
