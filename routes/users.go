package routes

import (
	"BookingApp/models"
	"BookingApp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't parse body"})
		return
	}
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't save user"})
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})

}

func Login(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't parse body"})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "User unauthorized",
		})

		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't generate token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User logged in",
		"token":   token,
	})
}
