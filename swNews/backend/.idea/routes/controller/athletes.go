package controller

import (
	"backend/service/athletes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAthletes(c *gin.Context) {
	var service athletes.AthleteService
	res := service.GetAthletes()
	c.JSON(http.StatusOK, res)
}
