package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/services"
)

func SearchHandler(c *gin.Context) {
	var request struct {
		Terms []string `json:"terms"`
		Email string   `json:"email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.PerformSearch(request.Terms, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "O diagnóstico será processado e enviado por email."})
}
