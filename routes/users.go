package routes

import (
	"net/http"

	"example.com/rest-api/model"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user model.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func getAllUsers(context *gin.Context) {
	events, err := model.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could fetch users.", "error": err})
		return
	}

	context.JSON(http.StatusOK, events)
}
