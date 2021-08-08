package main

import (
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
	DbCore "github.com/fidelyiu/yiu-note/core/db"
	langCore "github.com/fidelyiu/yiu-note/core/lang"
	LogCore "github.com/fidelyiu/yiu-note/core/log"
	Router "github.com/fidelyiu/yiu-note/core/router"
)

func main() {
	langCore.SetYiuToolLang(yiuLang.EnUS)
	DbCore.CreateDB(".yiu/yiu-reader.db")
	defer DbCore.CloseDB()
	LogCore.CreateLogger()
	r := Router.InitRouter()
	_ = r.Run(":8080")
}
