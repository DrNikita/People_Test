package main

import (
	"github.com/DrNikita/People/docs"
	"github.com/DrNikita/People/internal/db"
	"github.com/DrNikita/People/internal/handler"
	"github.com/DrNikita/People/internal/model"
	"github.com/DrNikita/People/internal/service/migration"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title People service
// @version 3.0
// @basePath /api/people
// @description People service test task
// @host localhost:8080
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.Host = context.Request.Host
		ginSwagger.CustomWrapHandler(&ginSwagger.Config{URL: "/swagger/doc.json"}, swaggerFiles.Handler)(context)
	})

	configs := config.GetConfigurationInstance()
	if err := config.InitDB(configs); err != nil {
		log.Error("error initializing db connection")
		panic(err)
	}

	if err := migration.MigrateTable(&model.Persons{}); err != nil {
		log.Error("error creating table persons")
		panic(err)
	}

	api := router.Group("/api/people")
	api.GET("/find-users", handler.FindPeople)

	if err := router.Run(":" + configs.AppPort); err != nil {
		panic(err)
	}
}
