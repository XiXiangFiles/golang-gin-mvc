package main

import (
	// "html/template"
	"webapp/router"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func main() {
	user := r.Group("/user")
	router.IndexRouter(user)
	r.LoadHTMLGlob("views/*")
	r.StaticFile("/favicon.ico", "./public/images/favicon.ico")
	r.Static("/images", "./public/images/")
	r.Run()
}
