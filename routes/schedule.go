package routes

import (
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/handlers"
	"github.com/gin-gonic/gin"
)

func addScheduleRoutes(rg *gin.RouterGroup) {
	schedule := rg.Group("/schedule")

	schedule.GET("/", func(ctx *gin.Context) {
		handlers.HelloWorld(ctx)
	})
}
