package main

import (
	"example/gin-example/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	goappenv "github.com/natepboat/go-app-env"
)

func main() {
	config := initApp()
	log.Println(config)

	runServer(config)
}

func initApp() map[string]interface{} {
	app := goappenv.NewAppEnv(os.DirFS("."))
	return app.Config()
}

func runServer(cfg map[string]interface{}) {
	router := gin.Default()
	router.GET("/albums", service.ListAlbum)

	router.Run(cfg["app.port"].(string))
}
