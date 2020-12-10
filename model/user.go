package model

import (
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Index(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
