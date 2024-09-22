package routes

import (
	"BookingApp/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func GetEventById(context *gin.Context) {
	eventId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func CreateEvent(context *gin.Context) {

	userId := context.GetInt("userId")

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	event.UserID = userId
	log.Println(event)
	err := event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
	})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	userId := context.GetInt("userId")

	if userId == event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not registered!",
		})
		return
	}

	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	err = event.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event Updated",
	})

}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	if event.UserID != context.GetInt("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not registered!",
		})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event Deleted",
	})
}
