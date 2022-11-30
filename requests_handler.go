package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"

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

	userDB.TimeStamp = time.Now().String()
	fmt.Println(userDB)

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

	result := "좋은아침이에요 "
	if userDB.TimeStamp != "" {

		// Weather
		HomeTownNameFromNuguCandle := nuguRequest.Action.Parameters.Location.Value
		wHomeDesc, wHomeTemp, wHomeNowRain, wHomeFutureRain := GetWeatherInfoByTownName(HomeTownNameFromNuguCandle)
		result += " 현재 " + HomeTownNameFromNuguCandle + " 날씨는 " + wHomeTemp + "도로 " + wHomeDesc + " " + wHomeNowRain + " " + wHomeFutureRain

		wDestDesc, wDescTemp, wDescNowRain, wDescFutureRain := GetWeatherInfoByTownName(userDB.DestAddr)
		result += userDB.DestAddr + "에는 " + wDestDesc + " " + wDescTemp + " " + wDescNowRain + " " + wDescFutureRain + "입니다. "

		// Stock
		stockPrice := GetStockPriceByStockName(userDB.StockName)
		result += userDB.StockName + "의 어제 종가는 " + stockPrice + "원 이에요."

		// Anniversary
		dDayMessage := GetDDayInfoByDate(userDB.SpecialDay)
		fmt.Println(dDayMessage)
		if dDayMessage != "" {
			result += " 오늘은 기념일 " + dDayMessage
		}
	}

	result += " 종은하루되세요!"

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

	result := "네! "
	if userDB.TimeStamp != "" {

		// Bus
		busArrivalMessage := GetBusArrivalTimeByCodes(userDB.StationInfo.arsId, userDB.BusInfo.busRouteNm)
		result += userDB.StationInfo.stNm + " 정류장에 " + busArrivalMessage

		// Weather
		HomeTownNameFromNuguCandle := nuguRequest.Action.Parameters.Location.Value
		_, _, wHomeNowRain, wHomeFutureRain := GetWeatherInfoByTownName(HomeTownNameFromNuguCandle)
		_, _, wDescNowRain, wDescFutureRain := GetWeatherInfoByTownName(userDB.DestAddr)
		if wHomeNowRain != "" || wHomeFutureRain != "" || wDescNowRain != "" || wDescFutureRain != "" {
			result += "나갈때 우산 챙기는거 잊지마세요! "
		}
	}

	result += " 잘 다녀오세요!"

	nuguResponse.Output.ResultSeeya = result
	fmt.Println(nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

