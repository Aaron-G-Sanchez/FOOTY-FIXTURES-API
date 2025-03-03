package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Starts the server and assigns sets the available routes.
func Run(apiToken string) {
	getRoutes(apiToken)
	fmt.Println("Server listening on localhost:3001")
	router.Run(":3001")
}

// Gets all available routes and their groups.
func getRoutes(apiToken string) {
	v1 := router.Group("/v1")
	addScheduleRoutes(v1, apiToken)
}
