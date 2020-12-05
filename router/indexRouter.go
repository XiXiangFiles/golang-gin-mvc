package router

import (
	"webapp/model"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func IndexRouter(rg *gin.RouterGroup) {
	rg.GET("/ping", model.User)
}
