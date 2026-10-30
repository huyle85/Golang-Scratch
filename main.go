package main

import (
	"github.com/gin-gonic/gin"
	v1handler "huyle.com/golang-scratch/internal/api/v1/handler"
)

func main() {
	r := gin.Default()

	UserHandlerV1 := v1handler.NewUserHandler()
	r.GET("/api/v1/users", UserHandlerV1.GetUserV1)
	r.GET("/api/v1/users/:user_id", UserHandlerV1.GetUserById)
	r.POST("/api/v1/users", UserHandlerV1.PostUsers)
	r.PUT("/api/v1/user/:id", UserHandlerV1.PutUsers)
	r.DELETE("api/v1/users/delete/:id", UserHandlerV1.DeleteUser)
	r.Run(":8080")
}