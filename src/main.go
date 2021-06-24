package main

import (
	"cmdb-go/api"
	"cmdb-go/config"
	_ "cmdb-go/docs"
	"cmdb-go/modules"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// @title CDMB平台API
	// @version 1.0
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name token
	f, _ := os.Create(config.LogFile)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	url := ginSwagger.URL("http://xxx:8888/swagger/doc.json")
	router := gin.Default()
	CorsConfig := cors.DefaultConfig()
	CorsConfig.AllowAllOrigins = true
	CorsConfig.AllowCredentials = true
	CorsConfig.AllowMethods = []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "HEADER"}
	CorsConfig.AllowHeaders = []string{"*"}
	CorsConfig.ExposeHeaders = []string{"Authorization", "pro", "reg"}
	router.Use(cors.New(CorsConfig))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Use(modules.AuthMiddleWare())
	v1 := router.Group("/web/v1/cmdb")
	{
		v1.POST("/agent", api.AgentUpdate)
	}
	_ = router.Run(":8888")
}
