package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
)

func main() {
	jsonFile, err := os.Open("db.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]any

	json.Unmarshal([]byte(byteValue), &result)

	router := gin.Default()
	//set Https TSL
	//router := gin.Default()
	//go router.RunTLS(":443", config.SSL.CertFile, config.SSL.CertKey)
	//User Login
	userLogin := result["USERLOGIN"].(map[string]any)
	fmt.Println(reflect.TypeOf(userLogin))
	//User Information
	userInfo := result["USERINFO"].(map[string]any)

	homeInfo := result["HomeInfo"].(map[string]any)

	router.POST("/userLogin", func(c *gin.Context) { c.JSON(200, userLogin) })

	router.GET("/userInfo", func(c *gin.Context) { c.JSON(200, userInfo) })

	router.GET("/getImage", func(c *gin.Context) { c.File("./image/Img_debit_card.jpg") })

	router.GET("/homeInfo", func(c *gin.Context) { c.JSON(200, homeInfo) })

	router.POST("/addHomeInfo", func(c *gin.Context) {
		// 定義單筆 info 的結構
		var newInfo map[string]interface{}

		// 嘗試解析傳入的 JSON
		if err := c.ShouldBindJSON(&newInfo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON format", "detail": err.Error()})
			return
		}

		// 取出原本的 infoList
		homeResult := homeInfo["result"].(map[string]interface{})
		infoList := homeResult["infoList"].([]interface{})

		// 新增這筆 info
		infoList = append(infoList, newInfo)
		homeResult["infoList"] = infoList
		homeInfo["result"] = homeResult

		// 更新整個 JSON 結構
		result["HomeInfo"] = homeInfo
		updatedJson, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to serialize JSON"})
			return
		}

		// 寫回 db.json 檔案
		err = ioutil.WriteFile("db.json", updatedJson, 0644)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to write to db.json", "detail": err.Error()})
			return
		}

		c.JSON(200, gin.H{"status": "New info added successfully"})
	})

	//using go run index.go
	router.Run("localhost:3001")

	//using docker
	//router.Run(":3001")
}

/*
*Post JSON by local file
 */
