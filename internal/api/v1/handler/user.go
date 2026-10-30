package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user_handler struct {
}

func NewUserHandler() *user_handler {
	return &user_handler{}
}

func (u *user_handler) GetUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "list message"})
}

func (u *user_handler) GetUserById(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	ctx.JSON(200, gin.H{"user": user_id})
}

func (u *user_handler) PostUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
func (u *user_handler) PutUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "update user success"})
}
func (u *user_handler) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "delete user"})
}
