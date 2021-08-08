package DbService

import (
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
	"github.com/fidelyiu/yiu-note/core/bean"
	langCore "github.com/fidelyiu/yiu-note/core/lang"
	DbDao "github.com/fidelyiu/yiu-note/dao/db-dao"
	"github.com/fidelyiu/yiu-note/model/dto"
	"github.com/fidelyiu/yiu-note/model/enum"
	"github.com/fidelyiu/yiu-note/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var serviceName = map[yiuLang.YiuLang]string{
	yiuLang.ZhCN: "数据库",
	yiuLang.EnUS: "dataBase",
}

func Search(c *gin.Context) response.YiuReaderResponse {
	result := response.YiuReaderResponse{}
	var searchDto dto.DbPageSearchDto
	err := c.ShouldBindJSON(&searchDto)
	if err != nil {
		bean.GetLoggerBean().Error(
			langCore.GetLangByMap(c, map[yiuLang.YiuLang]string{
				yiuLang.ZhCN: "查询" + langCore.GetLangByMap(c, serviceName) + "出错，Body参数转换出错!",
				yiuLang.EnUS: "Search " + langCore.GetLangByMap(c, serviceName) + " error，Body parameter conversion error!",
			}),
			zap.Error(err),
		)
		result.ToError(err.Error())
		return result
	}

	searchDto.Check()

	dbSearchVo, err := DbDao.FindBySearchDto(langCore.GetContextLang(c), searchDto)
	if err != nil {
		bean.GetLoggerBean().Error(
			langCore.GetLangByMap(c, map[yiuLang.YiuLang]string{
				yiuLang.ZhCN: "查询" + langCore.GetLangByMap(c, serviceName) + "出错!",
				yiuLang.EnUS: "Search " + langCore.GetLangByMap(c, serviceName) + " error!",
			}),
			zap.Error(err),
		)
		result.ToError(err.Error())
		return result
	}

	result.Result = dbSearchVo
	result.SetType(enum.ResultTypeSuccess)
	return result
}
