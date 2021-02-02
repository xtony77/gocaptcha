package main

import (
	"gocaptcha/internal/gin/route"
	config "gocaptcha/configs"
)

func main() {
	config := config.NewConfig()
	router := route.SetupRouter(config)
	router.Run(config.Gin.Host + ":" + config.Gin.Port)
}
