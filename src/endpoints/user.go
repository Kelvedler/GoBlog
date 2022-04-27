package endpoints

import (
	"fmt"
	"net/http"

	"github.com/Kelvedler/GoBlog/models"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {

	var user models.UserShort
	var createdUser models.UserFull
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err = models.CreateNewUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, createdUser)
	}
}

func List(context *gin.Context) {
	usersSlice, err := models.GetSlice("id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if len(usersSlice) > 0 {
			context.JSON(http.StatusOK, usersSlice)
		}
	}
}
