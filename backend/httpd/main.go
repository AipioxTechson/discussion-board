package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:3001 (for windows "localhost:3001")
	if err != nil {
		panic("Failed to start server")
	}
}
