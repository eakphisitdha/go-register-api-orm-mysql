package main

import (
	"go/membersystem/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//Connect to Database
	auth.InitDB()

	r := gin.Default()
	r.POST("/register", auth.Register)
	r.Use(cors.Default())
	r.Run("localhost:8080")
}
