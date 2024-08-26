package controllers

import (
	"api/models"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Calculate(c *gin.Context) {
	var request models.OperationRequest

	// Liaison du corps de la requête JSON à la structure OperationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Évaluez l'expression via le service
	result, err := services.EvaluateExpression(request.Expression)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Envoie le résultat comme réponse JSON
	c.JSON(http.StatusOK, gin.H{"result": result})
}
