package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Starts the server and assigns sets the available routes.
func Run() {
	getRoutes()
	fmt.Println("Server listening on localhost:3001")
	router.Run(":3001")
}

func getRoutes() {
	v1 := router.Group("/v1")
	addScheduleRoutes(v1)
}
