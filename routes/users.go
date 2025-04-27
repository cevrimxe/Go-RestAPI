package routes

import (
	"net/http"

	"github.com/cevrimxe/Go-RestAPI/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt save"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Couldnt authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successfull"})
}
