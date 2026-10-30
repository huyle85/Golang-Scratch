package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get_user_v1 (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, gin.H{"message": "list message"})
	}

func get_user_byid(ctx *gin.Context){
	user_id := ctx.Param("user_id")
	ctx.JSON(200, gin.H{"user": user_id})
}

func post_users(ctx *gin.Context){
	ctx.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
func put_users(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"message": "update user success"})
}
func delete_user(ctx *gin.Context){
	ctx.JSON(http.StatusNoContent, gin.H{"message": "delete user"})
}
func main() {
	r := gin.Default()

	r.GET("/api/v1/users", get_user_v1)
	r.GET("/api/v1/users/:user_id", get_user_byid)
	r.POST("/api/v1/users", post_users)
	r.PUT("/api/v1/user/:id", put_users)
	r.DELETE("api/v1/users/delete/:id", delete_user)
	r.Run(":8080")
}