package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

////////////////////////////////////////////////////////////////////////////////
// Common Reqeust Handlers
////////////////////////////////////////////////////////////////////////////////
func GetHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "SKT NUGU Routine",
	})
}

func GetPing(c *gin.Context) {
	printLog("GET  \"" + c.Request.URL.String() + "\"")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetHealth(c *gin.Context) {
	printLog("GET  \"" + c.Request.URL.String() + "\"")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
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

