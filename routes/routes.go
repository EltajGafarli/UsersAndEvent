package routes

import (
	"BookingApp/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	//events
	router.GET("/api/events", GetEvents)
	router.GET("/api/events/:id", GetEventById)

	events := router.Group("/api/events")
	events.Use(middleware.Authenticate)
	events.POST("/", CreateEvent)
	events.PUT("/:id", UpdateEvent)
	events.DELETE("/:id", DeleteEvent)

	//users
	router.POST("/api/auth/register", Register)
	router.POST("/api/auth/login", Login)

}
