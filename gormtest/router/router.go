package router

import (
	"github.com/gin-gonic/gin"
	"gormtest/api"
)

func InitRouter(r *gin.Engine) {
	api.RegisterRouter(r)
}
