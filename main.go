package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var logStream []string

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
