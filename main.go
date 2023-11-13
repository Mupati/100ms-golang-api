package main

import (
	"net/http"

	"api/room"
	"api/token"

	"github.com/gin-gonic/gin"
)

func getHelp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}

func main() {
	router := gin.Default()
	router.POST("/token", token.CreateToken)
	router.POST("/room", room.CreateRoom)
	router.GET("/", getHelp)

	router.Run()
}
