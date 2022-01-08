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
	response, err := h.service.GetSummary()
	if err != nil {
		switch err {
		case ErrGetRequestError:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server can't pull data.",
			})
			return
		case ErrConvertToStructError:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server can't convert data.",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error.",
			})
			return
		}
	}
	c.JSON(http.StatusOK, response)
}
