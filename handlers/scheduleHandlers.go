package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// var fixturesByDateRangeURL = "https://api.sportmonks.com/v3/football/fixtures/between/2025-03-01/2025-03-03"

// Returns all matches occuring during the current matchday window.
func GetCurrentFixturesList(ctx *gin.Context, apiToken string) {
	// TODO: Add fetch to get all UPCOMING matches from sportmonks.

	// TODO: Return the current date/time at the time of request.
	ctx.JSON(http.StatusOK, gin.H{
		"current time": time.Now(),
	})
}
