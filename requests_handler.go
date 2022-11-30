package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func PostUserDB(c *gin.Context) {
	fmt.Println("POST  \"" + c.Request.URL.String() + "\"")

	// Read request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal([]byte(body), &userDB)
	if err != nil {
		fmt.Println(err.Error())
	}

	userDB.StationInfo = GetStationInfoByStationName(userDB.BusStop)
	userDB.BusInfo = GetBusInfoByBusNumber(userDB.BusNum)

	/* Start Debug Log */
	// TODO Remove this block later

	fmt.Println(userDB)

	fmt.Println("")
	fmt.Println("# Weather Info #")
	strWeather1 := GetWeatherInfoByTownName("관악구")
	fmt.Println(strWeather1)
	strWeather1, strWeather2, strWeather3, strWeather4 := GetWeatherInfoByTownName(userDB.DestAddr)
	fmt.Println(strWeather1 + " / " + strWeather2 + " / " + strWeather3 + " / " + strWeather4)

	fmt.Println("")
	fmt.Println("# Bus Info #")
	strBus := GetBusArrivalTimeByCodes(userDB.StationInfo.arsId, userDB.BusInfo.busRouteNm)
	fmt.Println(strBus)

	fmt.Println("")
	fmt.Println("# Stock Info #")
	strStock := GetStockPriceByStockName(userDB.StockName)
	fmt.Println(strStock)

	fmt.Println("")
	fmt.Println("# Anniversary Info ")
	strAnniversary := GetDDayInfoByDate(userDB.SpecialDay)
	/* End Debug Log */

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func PostGoodmorning(c *gin.Context) {
	fmt.Println("POST  \"" + c.Request.URL.String() + "\"")

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

	nuguResponse.Output.ResultGoodmorning = result
	fmt.Println(nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

func PostSeeya(c *gin.Context) {
	fmt.Println("POST  \"" + c.Request.URL.String() + "\"")

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

	nuguResponse.Output.ResultSeeya = result
	fmt.Println(nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

