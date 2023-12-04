package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/users", CreateNewUser)
	router.POST("/api/users/generateotp", GenerateOTP)
	router.POST("/api/users/verifyotp", VerifyOTP)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
