package controllers

import (
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Historique(c *gin.Context) {

	results, err := services.LoadHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Envoie le résultat comme réponse JSON
	c.JSON(http.StatusOK, gin.H{"history": results})
}
