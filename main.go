package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const serviceKey = "Uc70KK1K8bzhcwQ+y+durUkD2VMV8wyequ5hxhQ39ghB0fJ0v3/mtW2qB4l/YRTs3w9YFSP47MRfnSVVszwb6A=="

var logStream []string

var db_station map[string]*dbStation
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
	r.GET("/health", GetHealth)

	r.POST("/dummy", PostDummy)
	r.POST("/goodmorning", PostGoodmorning)
	r.POST("/seeya", PostSeeya)

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
