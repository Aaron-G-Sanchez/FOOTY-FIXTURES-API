package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	// TODO: Add fetch to get all UPCOMING matches from sportmonks.
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Schedule handler",
	})
}
