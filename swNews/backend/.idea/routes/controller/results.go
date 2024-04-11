package controller

import (
	"backend/service/results"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetResults(c *gin.Context) {
	var service results.ResultService
	res := service.GetResults()
	c.JSON(http.StatusOK, res)
}

func GetResultDetail(c *gin.Context) {
	var service results.DetailService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetDetail()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
