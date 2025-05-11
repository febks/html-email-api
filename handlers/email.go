package handlers

import (
	"go-send-email/models"
	"go-send-email/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendEmailHandler(c *gin.Context) {
	var req models.EmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := services.SendEmail(req.To, req.Cc, req.Subject, req.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
