package covid

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service CovidService
}

func NewCovidHandler(router *gin.RouterGroup, service CovidService) {
	handler := &handler{service: service}

	router.GET("/summary", handler.getSummary)
}

func (h *handler) getSummary(c *gin.Context) {
	response, _ := h.service.GetSummary()
	c.JSON(http.StatusOK, response)
}
