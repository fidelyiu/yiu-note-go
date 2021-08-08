package Router

import (
	"github.com/fidelyiu/yiu-note/core/asset"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	SetMaxMultipartMemory(r, 8<<20) // 8 MiB
	asset.SetVueHistory(r)
	SetDbRouter(r)
	return r
}

func SetMaxMultipartMemory(r *gin.Engine, m int64) {
	r.MaxMultipartMemory = m
}
