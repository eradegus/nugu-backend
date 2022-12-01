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

	// DB for bus
	userDB.StationInfo = GetStationInfoByStationName(userDB.BusStop)
	userDB.BusInfo = GetBusInfoByBusNumber(userDB.BusNum)

	// DB for weather
	userDB.HomeTown.LocationX, userDB.HomeTown.LocationY = GetTownLocationCoordinateByName(userDB.HomeAddr)
	userDB.DestTown.LocationX, userDB.DestTown.LocationY = GetTownLocationCoordinateByName(userDB.DestAddr)

	// DB for zip
	userDB.ZipInfo.TownName, userDB.ZipInfo.TownCode = GetLocationInfoByAptName(userDB.ZipName)

	userDB.TimeStamp = time.Now().String()
	fmt.Println(userDB)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func PostGoodmorning(c *gin.Context) {
	fmt.Printf("[SKT-NUGU] POST  \"%s\" from speaker.nugu.nu110.se ..\n",  c.Request.URL.String())

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
	fmt.Printf("[SKT-NUGU] Request: %v\n", nuguRequest)

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
		wHomeDesc, wHomeTemp, wHomeIsRain, wHomeIsSnow := GetWeatherInfoByTownLocation(userDB.HomeTown)

		switch wHomeDesc {
		case "1":
			wHomeDesc = "맑고"
		case "3":
			wHomeDesc = "구름이 많고"
		case "4":
			wHomeDesc = "흐리고"
		}

		result += " 현재 " + HomeTownNameFromNuguCandle + " 날씨는 " + wHomeDesc + " 기온은 " + wHomeTemp+ "도 입니다. "

		wDestDesc, _, wDestIsRain, wDestIsSnow := GetWeatherInfoByTownLocation(userDB.DestTown)
		switch wDestDesc {
		case "1":
			wDestDesc = "맑습니다. "
		case "3":
			wDestDesc = "구름이 많습니다. "
		case "4":
			wDestDesc = "흐립니다. "
		}

		result += userDB.DestAddr + "는 " + wDestDesc

		//////////////////////////////
		// NOTE: ONLY FOR DEMO
		wHomeIsSnow = true

		if wHomeIsRain || wDestIsRain {
			result += "오늘은 비"
			if wHomeIsSnow || wDestIsSnow {
				result += " 또는 눈소식이 있어요. "
			} else {
				result += "소식이 있어요. "
			}
		} else if wHomeIsSnow || wDestIsSnow {
			result += "오늘은 눈소식이 있어요. "
		}

		// Stock
		stockPrice := GetStockPriceByStockName(userDB.StockName)
		result += userDB.StockName + "의 어제 종가는 " + stockPrice + "원이에요."

		// Real Estate
		zipPrice := GetZipPriceByCodeAndName(userDB.ZipInfo.TownCode, userDB.ZipName)
		if zipPrice != "" {
			result += " 어제 " + userDB.ZipName + " 아파트에 새로운 실거래가 발생했어요."
			result += " 실거래가는 " + zipPrice + "원이에요."
		}

		// Anniversary
		dDayMessage := GetDDayInfoByDate(userDB.SpecialDay)
		if dDayMessage != "" {
			result += " 오늘은 기념일 " + dDayMessage
		}
	}

	result += " 좋은 하루 되세요!"

	nuguResponse.Output.ResultGoodmorning = result
	fmt.Printf("[SKT-NUGU] Response: %v\n", nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

func PostSeeya(c *gin.Context) {
	fmt.Printf("[SKT-NUGU] POST  \"%s\" from speaker.nugu.nu110.se ..\n",  c.Request.URL.String())

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
	fmt.Printf("[SKT-NUGU] Request: %v\n", nuguRequest)

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
		_, _, wHomeIsRain, wHomeIsSnow := GetWeatherInfoByTownLocation(userDB.HomeTown)
		_, _, wDescIsRain, wDescIsSnow := GetWeatherInfoByTownLocation(userDB.DestTown)

		//////////////////////////////
		// NOTE: ONLY FOR DEMO
		wHomeIsSnow = true

		if wHomeIsRain || wHomeIsSnow || wDescIsRain || wDescIsSnow {
			result += " 오늘 비 또는 눈소식이 있으니, 나갈때 우산 챙기는거 잊지마세요! "
		}
	}

	result += " 잘 다녀오세요!"

	nuguResponse.Output.ResultSeeya = result
	fmt.Printf("[SKT-NUGU] Response: %v\n", nuguResponse)

	// End Logic
	//////////////////////////////////////////////////

	c.JSON(http.StatusOK, nuguResponse)
}

