package main

import (
	"webapp/router"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func main() {
	user := r.Group("/user")
	router.IndexRouter(user)
	r.Run()
}
