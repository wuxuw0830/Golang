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

	router.POST("/userLogin", func(c *gin.Context) { c.JSON(200, userLogin) })

	router.GET("/userInfo", func(c *gin.Context) { c.JSON(200, userInfo) })

	router.GET("/getImage", func(c *gin.Context) { c.File("./image/Img_debit_card.jpg") })
	//using go run index.go
	//router.Run("localhost:3001")

	//using docker
	router.Run(":3001")
}

/*
*Post JSON by local file
 */
