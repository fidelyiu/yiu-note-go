package Router

import (
	DbController "github.com/fidelyiu/yiu-note/controller/db-controller"
	"github.com/gin-gonic/gin"
)

func SetDbRouter(r *gin.Engine) {
	dbGroup := r.Group("/db")
	{
		dbGroup.POST("", DbController.Search)
	}
}
