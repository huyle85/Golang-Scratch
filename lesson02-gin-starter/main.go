package main

import (

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/demo", func (ctx *gin.Context)  {
		ctx.JSON(200, gin.H{"message": "hello Le Van Huy"})
	})

	r.GET("/user/:user_id", func (ctx *gin.Context)  {
		user_id := ctx.Param("user_id")	
		age := ctx.Query("age")
		ctx.JSON(200, gin.H{
			"data": "thong tin user",
			"user_id": user_id,
			"age": age,
		})
	})

	r.GET("/products", func (ctx *gin.Context)  {
		ctx.JSON(200, gin.H{"data": "Danh sach san pham"})
	})
	r.Run(":8080")
}