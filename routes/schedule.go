package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addScheduleRoutes(rg *gin.RouterGroup) {
	schedule := rg.Group("/schedule")

	schedule.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Schedule route hit!",
		})
	})
}
