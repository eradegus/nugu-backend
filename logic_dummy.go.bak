package main

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

func DummyInfo() string {
	return "창 밖을 보세요"
}

// Examples 1 - Search detailed station info by station name
func DummyStationName(input string) string {

	// Placeholder
	stSrch := input		// ex. "관악경찰서"
	printLog(" └ stSrch: " + stSrch)
	userId := "jeff"

	// Generate HTTP GET request
	api := "http://ws.bus.go.kr/api/rest/stationinfo/getStationByName?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("stSrch", stSrch)
	params.Add("resultType", "json")
	printLog(api + params.Encode())

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
	//jsonData := jsonstruct.OpenAPIResponse{}
	jsonData := OpenAPIResponse{}
	json.Unmarshal(resData, &jsonData)
	result := jsonData.MsgBody.ItemList[0]

	// Store DB
	station := new(dbStation)
	station.stNm = result.StNm
	station.arsId = result.ArsID
	station.stationId = result.StID

	db_station[userId] = station
	fmt.Println(station)

	return station.stNm
}

// Examples 2 - Search detailed bud info by bus number
func DummyBusNumber(input string) string {

	// Placeholder
	strSrch := input	// ex. "5511"
	printLog(" └ strSrch: " + strSrch)
	userId := "jeff"

	// Generate HTTP GET request
	api := "http://ws.bus.go.kr/api/rest/busRouteInfo/getBusRouteList?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("strSrch", strSrch)
	params.Add("resultType", "json")
	printLog(api + params.Encode())

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
	//jsonData := jsonstruct.OpenAPIResponse{}
	jsonData := OpenAPIResponse{}
	json.Unmarshal(resData, &jsonData)
	result := jsonData.MsgBody.ItemList[0]

	bus := new(dbBus)
	bus.busRouteNm = result.BusRouteNm
	bus.busRouteId = result.BusRouteID

	db_bus[userId] = bus
	fmt.Println(bus)

	return bus.busRouteNm
}

// Examples 3 - Get bus arrival info (need ex. 1 and 2 in advance)
func DummyBusTime() string {
	userId := "jeff"
	userStation, ok := db_station[userId]
	if !ok {
		fmt.Printf("db_station[" + userId + "] DB not exists\n")
		return ""
	}
	userBus, ok := db_bus[userId]
	if !ok {
		fmt.Printf("db_bus[" + userId + "] DB not exists\n")
		return ""
	}
	arsId := userStation.arsId
	busRouteNm := userBus.busRouteNm

	// Generate HTTP GET request
	api := "http://ws.bus.go.kr/api/rest/stationinfo/getStationByUid?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("arsId", arsId)
	params.Add("resultType", "json")
	printLog(api + params.Encode())

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
	//jsonData := jsonstruct.OpenAPIResponse{}
	jsonData := OpenAPIResponse{}
	json.Unmarshal(resData, &jsonData)

	var msg1, msg2 string
	for _, item := range jsonData.MsgBody.ItemList {
		if item.RtNm == busRouteNm {
			msg1 = item.Arrmsg1
			msg2 = item.Arrmsg2

			break;
		}
	}

	msg1 = makeSense(msg1)
	msg2 = makeSense(msg2)
	data := busRouteNm + " 버스는 " + msg1 + "입니다. 다음 버스는 " + msg2 + "입니다."

	return data
}

func makeSense(input string) string {
	msg := strings.Split(input, "[")[0]
	msg = strings.Replace(msg, " 도착", "", -1)

	if strings.Compare(msg, "운행종료") != 0 {
		msg += " 도착 예정"
	}

	return msg
}

