package DbController

import (
	DbService "github.com/fidelyiu/yiu-note/service/db-service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(c *gin.Context) {
	c.JSON(http.StatusOK, DbService.Search(c))
}
