package main

import (
	routes "run/question1/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/users", routes.CreateNewUser)
	router.POST("/api/users/generateotp", routes.GenerateOTP)
	router.POST("/api/users/verifyotp", routes.VerifyOTP)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
