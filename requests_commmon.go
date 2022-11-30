package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

////////////////////////////////////////////////////////////////////////////////
// Common Reqeust Handlers
////////////////////////////////////////////////////////////////////////////////
func GetHomePage(c *gin.Context) {
	fmt.Println("GET  \"" + c.Request.URL.String() + "\"")

	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "SKT NUGU Routine",
	})
}

func GetPing(c *gin.Context) {
	fmt.Println("GET  \"" + c.Request.URL.String() + "\"")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetHealth(c *gin.Context) {
	fmt.Println("GET  \"" + c.Request.URL.String() + "\"")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

