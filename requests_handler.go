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

	result := "좋은 아침이에요 "
	if userDB.TimeStamp != "" {

		// Weather
		HomeTownNameFromNuguCandle := nuguRequest.Action.Parameters.Location.Value
		wHomeDesc, wHomeTemp, wHomeIsRain, wHomeIsSnow := GetWeatherInfoByTownName(HomeTownNameFromNuguCandle)
		result += " 현재 " + HomeTownNameFromNuguCandle + " 날씨는 " + wHomeDesc + " 기온은 " + wHomeTemp+ "도 입니다. "
		
		wDestDesc, wDescTemp, wDescIsRain, wDescIsSnow := GetWeatherInfoByTownName(userDB.DestAddr)
		result += userDB.DestAddr + "는 " + wDestDesc + " 기온은 " + wDescTemp + "도 입니다. 오늘은 "

		if wHomeIsRain || wDescIsRain {
			result += "비"
			if wHomeIsSnow || wDescIsSnow {
				result += " 또는 눈"
			} else {
				result += "소식이 있어요. "
			}
		}
		
		if wHomeIsSnow || wDescIsSnow {
			result += "눈소식이 있어요. "
		}

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

	result += " 좋은 하루 되세요!"

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
		_, _, wHomeIsRain, wHomeIsSnow := GetWeatherInfoByTownName(HomeTownNameFromNuguCandle)
		_, _, wDescIsRain, wDescIsSnow := GetWeatherInfoByTownName(userDB.DestAddr)
		if wHomeIsRain || wHomeIsSnow || wDescIsRain || wDescIsSnow {
			result += " 오늘 비 또는 눈소식이 있으니, 나갈때 우산 챙기는거 잊지마세요! "
		}
	}

	result += " 잘 다녀오세요!"

	nuguResponse.Output.ResultSeeya = result
	fmt.Println(nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

