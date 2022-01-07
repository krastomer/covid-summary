package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/krastomer/covid-summary/internal/covid"
)

func Run() {

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	covidRouter := router.Group("/covid")
	covid.NewCovidHandler(covidRouter, nil)

	router.Run(":8080")
}
