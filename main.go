package main

import (
	"github.com/OmarDardery/home-passwords/db"
	"github.com/OmarDardery/home-passwords/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.POST("/precious", func(c *gin.Context) {
		var precious models.Precious
		if err := c.ShouldBindJSON(&precious); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		precious.Save()
		c.JSON(200, precious)
	})
	server.GET("/listprecious", func(c *gin.Context) {
		c.JSON(200, models.GetAllPrecious())
	})
	server.GET("/precious/:name", func(c *gin.Context) {
		n := c.Param("name")
		p, err := models.GetPreciousByName(n)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "what the hell man",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "nya:3",
			"data":    p,
		})

	})
	server.Run(":8080")
}
