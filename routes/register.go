package routes

import (
	"net/http"
	"strconv"

	"github.com/cevrimxe/Go-RestAPI/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}
	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}
	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "cancelled!"})
}
