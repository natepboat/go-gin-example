package main

import (
	"example/gin-example/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", service.ListAlbum)

	router.Run(":9000")
}
