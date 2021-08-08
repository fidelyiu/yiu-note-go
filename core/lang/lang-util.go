package langCore

import (
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
	"github.com/gin-gonic/gin"
)

func GetContextLang(c *gin.Context) yiuLang.YiuLang {
	lang := yiuLang.YiuLang(c.Request.Header.Get("Accept-Language"))
	switch lang {
	case yiuLang.ZhCN:
		return yiuLang.ZhCN
	default:
		return yiuLang.EnUS
	}
}

func GetLangByMap(c *gin.Context, m map[yiuLang.YiuLang]string) string {
	return GetLangByKey(GetContextLang(c), m)
}

func GetLangByKey(y yiuLang.YiuLang, m map[yiuLang.YiuLang]string) string {
	if len(m) == 0 {
		return ""
	} else {
		result := m[y]
		if result == "" {
			for lang := range m {
				return m[lang]
			}
		}
		return result
	}
}
