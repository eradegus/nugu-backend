package main

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

// Search detailed station info by station name
func GetStationInfoByStationName(input string) StationInfo {

	stSrch := input		// ex. "관악경찰서"

	// Generate HTTP GET request
	api := "http://ws.bus.go.kr/api/rest/stationinfo/getStationByName?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("stSrch", stSrch)
	params.Add("resultType", "json")
	fmt.Println(api + params.Encode())

	resp, err := http.Get(api + params.Encode())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Unmarshal json string
	jsonData := OpenAPIResponse_bus{}
	json.Unmarshal(resData, &jsonData)
	result := jsonData.MsgBody.ItemList[0]

	// Store DB
	stationInfo := new(StationInfo)
	stationInfo.stNm = result.StNm
	stationInfo.arsId = result.ArsID
	stationInfo.stationId = result.StID

	return *stationInfo
}

// Search detailed bud info by bus number
func GetBusInfoByBusNumber(input string) BusInfo {

	strSrch := input	// ex. "5511"

	// Generate HTTP GET request
	api := "http://ws.bus.go.kr/api/rest/busRouteInfo/getBusRouteList?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("strSrch", strSrch)
	params.Add("resultType", "json")
	fmt.Println(api + params.Encode())

	resp, err := http.Get(api + params.Encode())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Unmarshal json string
	jsonData := OpenAPIResponse_bus{}
	json.Unmarshal(resData, &jsonData)
	result := jsonData.MsgBody.ItemList[0]

	busInfo := new(BusInfo)
	busInfo.busRouteNm = result.BusRouteNm
	busInfo.busRouteId = result.BusRouteID

	return *busInfo
}

// Get bus arrival info
func GetBusArrivalTimeByCodes(stationCode string, busCode string) string {
	arsId := stationCode
	busRouteNm := busCode

	// Generate HTTP GET request
	api := "http://ws.bus.go.kr/api/rest/stationinfo/getStationByUid?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("arsId", arsId)
	params.Add("resultType", "json")
	fmt.Println(api + params.Encode())

	resp, err := http.Get(api + params.Encode())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Unmarshal json string
	jsonData := OpenAPIResponse_bus{}
	json.Unmarshal(resData, &jsonData)

	var arrmsg1, arrmsg2 string
	for _, item := range jsonData.MsgBody.ItemList {
		if item.RtNm == busRouteNm {
			arrmsg1 = item.Arrmsg1
			arrmsg2 = item.Arrmsg2

			break;
		}
	}

	arrmsg1 = makeBusString(arrmsg1)
	arrmsg2 = makeBusString(arrmsg2)
	data := busRouteNm + " 버스는 " + arrmsg1 + "입니다. 다음 버스는 " + arrmsg2 + "입니다."

	return data
}

func makeBusString(input string) string {
	msg := strings.Split(input, "[")[0]
	msg = strings.Replace(msg, " 도착", "", -1)

	if strings.Compare(msg, "운행종료") != 0 {
		msg += " 도착 예정"
	}

	return msg
}

