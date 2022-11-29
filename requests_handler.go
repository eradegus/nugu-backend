package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func PostUserDB(c *gin.Context) {
	printLog("POST  \"" + c.Request.URL.String() + "\"")

	// Read request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal([]byte(body), &userDB)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userDB)


	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func PostDummy(c *gin.Context) {
	printLog("POST  \"" + c.Request.URL.String() + "\"")

	// Read request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var input map[string]interface{}
	err = json.Unmarshal([]byte(body), &input)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(input)

	var nuguRequest NuguRequest
	err = json.Unmarshal([]byte(body), &nuguRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(nuguRequest)

	// Create response skeleton
	var nuguResponse NuguResponse
	nuguResponse.Version = nuguRequest.Version
	nuguResponse.ResultCode = "OK"

	//////////////////////////////////////////////////
	// Start Logic (logic_*.go file)

	stName := DummyStationName("관악경찰서")
	busNum := DummyBusNumber("5511")
	result := DummyBusTime()
	fmt.Println("[ ", stName, " / ", busNum, " ] - ", result)

	nuguResponse.Output.ResultString = result
	fmt.Println(nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

func PostGoodmorning(c *gin.Context) {
	printLog("POST  \"" + c.Request.URL.String() + "\"")

	// Read request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var nuguRequest NuguRequest
	err = json.Unmarshal([]byte(body), &nuguRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(nuguRequest)

	// Create response skeleton
	var nuguResponse NuguResponse
	nuguResponse.Version = nuguRequest.Version
	nuguResponse.ResultCode = "OK"

	//////////////////////////////////////////////////
	// Start Logic (logic_*.go file)

	result := "좋은아침 로직을 여기에 작성하세요"

	nuguResponse.Output.ResultString = result
	fmt.Println(nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

func PostSeeya(c *gin.Context) {
	printLog("POST  \"" + c.Request.URL.String() + "\"")

	// Read request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var nuguRequest NuguRequest
	err = json.Unmarshal([]byte(body), &nuguRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(nuguRequest)

	// Create response skeleton
	var nuguResponse NuguResponse
	nuguResponse.Version = nuguRequest.Version
	nuguResponse.ResultCode = "OK"

	//////////////////////////////////////////////////
	// Start Logic (logic_*.go file)

	result := "다녀올게 로직을 여기에 작성하세요"

	nuguResponse.Output.ResultString = result
	fmt.Println(nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

