package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	Message string `json:"message"`
}

var pong = message{
	Message: "pong",
}

func main() {
	router := gin.Default()
	router.GET("/ping", getPong)

	router.Run("localhost:8080")
}

func getPong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pong)
}
