package endpoints

import (
	"context"
	"net/http"

	"github.com/Kelvedler/GoBlog/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func UserRegister(ginCtx *gin.Context) {
	var userInput models.UserShort
	var createdUser models.UserFull
	ctx, ok := ginCtx.MustGet("topContext").(context.Context)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	conn, ok := ginCtx.MustGet("dbConn").(*pgx.Conn)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	err := ginCtx.ShouldBindJSON(&userInput)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err = models.UserCreateNew(ctx, conn, userInput)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ginCtx.JSON(http.StatusCreated, createdUser)
	}
}

func UserList(ginCtx *gin.Context) {
	ctx, ok := ginCtx.MustGet("topContext").(context.Context)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	conn, ok := ginCtx.MustGet("dbConn").(*pgx.Conn)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	paramOrderBy := ginCtx.Query("order_by")
	orderBy := &models.UserColumns[0]
	for _, i := range models.UserColumns {
		if paramOrderBy == i {
			orderBy = &paramOrderBy
			break
		}
	}
	usersSlice, err := models.UserGetSlice(ctx, conn, *orderBy)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if len(usersSlice) > 0 {
			ginCtx.JSON(http.StatusOK, usersSlice)
		}
	}
}

func UserSingle(ginCtx *gin.Context) {
	ctx, ok := ginCtx.MustGet("topContext").(context.Context)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	conn, ok := ginCtx.MustGet("dbConn").(*pgx.Conn)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	id, err := uuid.Parse(ginCtx.Param("user_id"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.UserGetByID(ctx, conn, id)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ginCtx.JSON(http.StatusOK, user)
}

func UserUpdate(ginCtx *gin.Context) {
	ctx, ok := ginCtx.MustGet("topContext").(context.Context)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	conn, ok := ginCtx.MustGet("dbConn").(*pgx.Conn)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	var userInput models.UserShort
	id, err := uuid.Parse(ginCtx.Param("user_id"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ginCtx.ShouldBindJSON(&userInput)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := models.UserUpdateByID(ctx, conn, id, userInput)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ginCtx.JSON(http.StatusOK, updatedUser)
	}
}

func UserDelete(ginCtx *gin.Context) {
	ctx, ok := ginCtx.MustGet("topContext").(context.Context)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	conn, ok := ginCtx.MustGet("dbConn").(*pgx.Conn)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
	}
	id, err := uuid.Parse(ginCtx.Param("user_id"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.UserDeleteByID(ctx, conn, id)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ginCtx.JSON(http.StatusNoContent, gin.H{})
	}
}
