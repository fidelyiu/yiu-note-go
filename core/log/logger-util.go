package LogCore

import (
	yiuLog "github.com/fidelyiu/yiu-go-tool/log"
	"github.com/fidelyiu/yiu-note/core/bean"
	"go.uber.org/zap"
)

func CreateLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		yiuLog.ErrorLn("日志初始化出错：")
		yiuLog.ErrorLn(err)
		return
	}
	bean.SetLoggerBean(logger)
}
