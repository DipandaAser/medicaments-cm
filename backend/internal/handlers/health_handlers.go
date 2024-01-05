package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health always returns 200 status code
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Medicament.cm, backend system good!"})
}
