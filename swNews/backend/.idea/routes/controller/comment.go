package controller

import (
	"backend/service/comment"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var service comment.CreateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ListComment(c *gin.Context) {
	var service comment.ListService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.List()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
