package endpoints

import (
	"net/http"

	"github.com/Kelvedler/GoBlog/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserRegister(context *gin.Context) {
	var userInput models.UserShort
	var createdUser models.UserFull
	err := context.ShouldBindJSON(&userInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err = models.UserCreateNew(userInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, createdUser)
	}
}

func UserList(context *gin.Context) {
	paramOrderBy := context.Query("order_by")
	orderBy := &models.UserColumns[0]
	for _, i := range models.UserColumns {
		if paramOrderBy == i {
			orderBy = &paramOrderBy
			break
		}
	}
	usersSlice, err := models.UserGetSlice(*orderBy)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if len(usersSlice) > 0 {
			context.JSON(http.StatusOK, usersSlice)
		}
	}
}

func UserSingle(context *gin.Context) {
	id, err := uuid.Parse(context.Param("user_id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.UserGetByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, user)
}

func UserUpdate(context *gin.Context) {
	var userInput models.UserShort
	id, err := uuid.Parse(context.Param("user_id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = context.ShouldBindJSON(&userInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := models.UserUpdateByID(id, userInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedUser)
	}
}

func UserDelete(context *gin.Context) {
	id, err := uuid.Parse(context.Param("user_id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.UserDeleteByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusNoContent, gin.H{})
	}
}
