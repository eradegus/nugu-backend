package main

import (
	"github.com/gin-gonic/gin"
)

const serviceKey = "Uc70KK1K8bzhcwQ+y+durUkD2VMV8wyequ5hxhQ39ghB0fJ0v3/mtW2qB4l/YRTs3w9YFSP47MRfnSVVszwb6A=="
var userDB UserDB

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/", GetHomePage)
	r.GET("/ping", GetPing)
	r.GET("/health", GetHealth)

	r.POST("/userdb", PostUserDB)

	r.POST("/goodmorning", PostGoodmorning)
	r.POST("/seeya", PostSeeya)

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

