package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/krastomer/covid-summary/internal/covid"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	covidRepo := covid.NewCovidRepository()
	covidService := covid.NewCovidService(covidRepo)

	covidRouter := router.Group("/covid")
	covid.NewCovidHandler(covidRouter, covidService)
	return router
}

func Run() {
	router := SetupRouter()
	router.Run(":8080")
}
