package main

import (
	"github.com/gin-gonic/gin"
	"guestbook/API"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/read-page", API.ReadPage)
	r.POST("/create-page", API.CreatePage)
	r.PUT("/update-page", API.UpdatePage)
	r.DELETE("/delete-page", API.DeletePage)
	return r
}
