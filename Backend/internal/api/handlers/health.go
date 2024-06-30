package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck handles the health check endpoint.
// @Summary      Health Check
// @Description  Check the health of the API
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200   {string} healthCheck
// @Router       /health [get]
func (app *Application) HealthCheck(c *gin.Context) {
	// Perform checks to determine health status.
	// For simplicity, I will assume the API is always healthy
	status := "ok"

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
