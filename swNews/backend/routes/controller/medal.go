package controller

import (
	"backend/service/medal"
	"github.com/gin-gonic/gin"
)

func GetMedals(c *gin.Context) {
	var service medal.Service
	if err := c.ShouldBindQuery(&service); err == nil {
		res := service.GetMedal()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
