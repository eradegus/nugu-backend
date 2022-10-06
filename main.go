package main

import (
	"fmt"
	"strings"
	"time"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/eradegus/nugu-bus/struct/openapi"
)

var logStream []string
var serviceKey = "Uc70KK1K8bzhcwQ+y+durUkD2VMV8wyequ5hxhQ39ghB0fJ0v3/mtW2qB4l/YRTs3w9YFSP47MRfnSVVszwb6A=="

type dbStation struct {
	stNm		string
	arsId		string
	stationId	string
}
var db_station map[string]*dbStation

type dbBus struct {
	busRouteId	string
	busRouteNm	string
}
var db_bus map[string]*dbBus

func GetHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Bus Info Query Server",
	})
}

func GetPing(c *gin.Context) {
	printLog("GET  \"" + c.Request.URL.String() + "\"")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func PostStationName(c *gin.Context) {
	printLog("POST  \"" + c.Request.URL.String() + "\"")

	// Read request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var input map[string]interface{}
	json.Unmarshal([]byte(body), &input)
	userId := fmt.Sprintf("%v", input["userId"])
	userId = strings.Replace(userId, " ", "", -1)
	stSrch := fmt.Sprintf("%v", input["stationName"])
	stSrch = strings.Replace(stSrch, " ", "", -1)

	printLog(" └ userId: " + userId)
	printLog(" └ stationName: " + stSrch)

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
	jsonData := openapi.Response{}
	json.Unmarshal(resData, &jsonData)
	result := jsonData.MsgBody.ItemList[0]

	// Store DB
	station := new(dbStation)
	station.stNm = result.StNm
	station.arsId = result.ArsID
	station.stationId = result.StID

	db_station[userId] = station

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": station.stNm,
	})
}

func PostBusNumber(c *gin.Context) {
	printLog("POST  \"" + c.Request.URL.String() + "\"")

	// Read request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var input map[string]interface{}
	json.Unmarshal([]byte(body), &input)
	userId := fmt.Sprintf("%v", input["userId"])
	strSrch := fmt.Sprintf("%v", input["busNumber"])
	printLog(" └ userId: " + userId)
	printLog(" └ strSrch: " + strSrch)

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
	jsonData := openapi.Response{}
	json.Unmarshal(resData, &jsonData)
	result := jsonData.MsgBody.ItemList[0]

	bus := new(dbBus)
	bus.busRouteNm = result.BusRouteNm
	bus.busRouteId = result.BusRouteID

	db_bus[userId] = bus

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": bus.busRouteNm,
	})
}

func GetBusTime(c *gin.Context) {
	printLog("GET  \"" + c.Request.URL.String() + "\"")

	userId := c.Query("userId")
	printLog(" └ userId: " + userId)

	userStation, ok := db_station[userId]
	if !ok {
		fmt.Printf("db_station[" + userId + "] DB not exists\n")
		return
	}
	userBus, ok := db_bus[userId]
	if !ok {
		fmt.Printf("db_bus[" + userId + "] DB not exists\n")
		return
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
	jsonData := openapi.Response{}
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

	c.JSON(http.StatusOK, gin.H{
		"message" : "OK",
		"data" : data,
	})
}

func GetLogStream(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "OK",
		"log": logStream,
	})
}

func PostClearLog(c *gin.Context) {
	logStream = make([]string, 0)
	printLog("[GIN-debug] Listening and serving HTTP")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func makeSense(input string) string {
	msg := strings.Split(input, "[")[0]
	msg = strings.Replace(msg, " 도착", "", -1)

	if strings.Compare(msg, "운행종료") != 0 {
		msg += " 도착 예정"
	}

	return msg
}

func main() {
	printLog("[GIN-debug] Listening and serving HTTP")

	db_station = map[string]*dbStation{}
	db_bus = map[string]*dbBus{}

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", GetHomePage)
	r.GET("/ping", GetPing)
	r.POST("/stationName", PostStationName)
	r.POST("/busNumber", PostBusNumber)
	r.GET("/busTime", GetBusTime)
	r.GET("/logstream", GetLogStream)
	r.POST("/clearlog", PostClearLog)
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credential", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST")
		c.Next()
	}
}

func printLog(msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logStream = append(logStream, "[" + timestamp + "] " + msg)
	fmt.Println(msg)
}
