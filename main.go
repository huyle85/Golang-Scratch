package main

import (
	"github.com/gin-gonic/gin"
	v1handler "huyle.com/golang-scratch/internal/api/v1/handler"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		UserHandlerV1 := v1handler.NewUserHandler()
		v1.GET("/users", UserHandlerV1.GetUserV1)
		v1.GET("/users/:user_id", UserHandlerV1.GetUserById)
		v1.POST("/users", UserHandlerV1.PostUsers)
		v1.PUT("/user/:id", UserHandlerV1.PutUsers)
		v1.DELETE("/users/delete/:id", UserHandlerV1.DeleteUser)
	}

	r.Run(":8080")
}
