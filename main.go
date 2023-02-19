package main

import (
	"example/gin-example/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	goappenv "github.com/natepboat/go-app-env"
)

func main() {
	app := goappenv.NewAppEnv(os.DirFS("."))
	log.Println(app)

	runServer(app)
}

func runServer(app *goappenv.AppEnv) {
	router := gin.Default()
	router.GET("/albums", service.ListAlbum)

	router.Run(app.Config()["app.port"].(string))
}
