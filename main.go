package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Main.go Start")
	initDB()

	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/getUsers", getUser)
	r.PUT("/updateUser", updateUser)
	r.DELETE("/deleteUser", deleteUser)
	r.Run(":8080")
	fmt.Println("Main.go End")
}
