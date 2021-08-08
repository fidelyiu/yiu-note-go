package entity

import (
	"errors"
	yiuSErr "github.com/fidelyiu/yiu-go-tool/error_s"
	yiuFile "github.com/fidelyiu/yiu-go-tool/file"
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
	FieldUtil "github.com/fidelyiu/yiu-note/core/field"
	langCore "github.com/fidelyiu/yiu-note/core/lang"
	"github.com/fidelyiu/yiu-note/model/enum"
)

type Image struct {
	Id     string         `json:"id"`
	Path   string         `json:"path"`   // 图片路径，相对于`/.yiu/image`的路径
	Status enum.ObjStatus `json:"status"` // 状态
	Src    string         `json:"src"`    // 源图片路径
}

func (i *Image) CheckPath(y yiuLang.YiuLang) error {
	if !yiuFile.IsExists(FieldUtil.ImageAdd + i.Path) {
		i.Status = enum.ObjStatusInvalid
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "图片'" + i.Path + "'不是有效文件",
			yiuLang.EnUS: "Image '" + i.Path + "' is not a valid file",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	i.Status = enum.ObjStatusValid
	return nil
}

func (i *Image) Check(y yiuLang.YiuLang) error {
	return yiuSErr.ToErrorBySep(" & ",
		i.CheckPath(y),
	)
}
