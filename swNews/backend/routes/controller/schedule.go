package controller

import (
	"backend/service/schedule"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSchedule(c *gin.Context) {
	var service schedule.Service
	res := service.GetSchedule()
	c.JSON(http.StatusOK, res)
}
