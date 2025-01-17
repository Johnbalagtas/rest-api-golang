package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.com/rest/models"
	"golang.com/rest/utils"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
