package endpoints

import (
	"net/http"

	"github.com/Kelvedler/GoBlog/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(context *gin.Context) {

	var user models.UserShort
	var createdUser models.UserFull
	err := context.ShouldBindJSON(&user)
	if err != nil {
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
	paramOrderBy := context.Query("order_by")
	orderBy := &models.UserColumns[0]
	for _, i := range models.UserColumns {
		if paramOrderBy == i {
			orderBy = &paramOrderBy
			break
		}
	}
	usersSlice, err := models.GetSlice(*orderBy)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if len(usersSlice) > 0 {
			context.JSON(http.StatusOK, usersSlice)
		}
	}
}

func Single(context *gin.Context) {
	id, err := uuid.Parse(context.Param("user_id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.GetByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, user)
}

func Update(context *gin.Context) {
	var user models.UserShort
	id, err := uuid.Parse(context.Param("user_id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := models.UpdateByID(id, user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedUser)
	}
}

func Delete(context *gin.Context) {
	id, err := uuid.Parse(context.Param("user_id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusNoContent, gin.H{})
	}
}
