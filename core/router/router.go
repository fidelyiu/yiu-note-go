package router

import (
	"github.com/fidelyiu/yiu-note/core/asset"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	asset.SetVueHistory(r)
	return r
}
