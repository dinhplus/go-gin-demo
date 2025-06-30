package main

import (
	_ "my-gin-app/docs" // Import generated Swagger docs
	"my-gin-app/internal/handler"
	"my-gin-app/internal/repository"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My Gin App API
// @version 1.0
// @description This is a sample server for a Gin application.
// @host localhost:8080
// @BasePath /

func main() {
	repository.InitDB()
	r := handler.SetupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}
