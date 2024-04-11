package routes

import (
	"backend/pkg/conf"
	"backend/pkg/logger/ginlogger"
	"backend/routes/controller"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	if conf.AppConf.Dev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(ginlogger.Logger())
	r.Use(ginlogger.Recovery())

	base := r.Group(conf.ServerConf.Prefix)
	{
		base.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		base.GET("/athletes", controller.GetAthletes)
		base.GET("/results", controller.GetResults)
		base.GET("/results/:id", controller.GetResultDetail)
		base.GET("/schedule", controller.GetSchedule)
		base.GET("/medals", controller.GetMedals)
		base.GET("/comments/:pageNo", controller.ListComment)
		base.POST("/comments", controller.CreateComment)
	}
	return r
}
