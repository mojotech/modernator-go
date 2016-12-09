package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
)

func main() {
	fmt.Println("heyo!")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("getThings", getting)

	router.Run(":" + port)
}

func getting(c *gin.Context) {
	c.JSON(200, gin.H{"stuff": "thaangs"})
}
