package entity

import (
	"errors"
	yiuSErr "github.com/fidelyiu/yiu-go-tool/error_s"
	yiuStr "github.com/fidelyiu/yiu-go-tool/string"
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
	langCore "github.com/fidelyiu/yiu-note/core/lang"
	PathUtil "github.com/fidelyiu/yiu-note/core/path"
	"github.com/fidelyiu/yiu-note/model/enum"
)

type EditSoft struct {
	Id      string         `json:"id"`      // Uuid
	Name    string         `json:"name"`    // 名称
	Path    string         `json:"path"`    // 绝对路径
	Img     string         `json:"img"`     // 软件图标地址
	SortNum int            `json:"sortNum"` // 排序数
	Status  enum.ObjStatus `json:"status"`  // 状态
}

func (e *EditSoft) CheckPath(y yiuLang.YiuLang) error {
	if !PathUtil.IsValidFile(e.Path) {
		e.Status = enum.ObjStatusInvalid
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "编辑软件'" + e.Path + "'不是有效文件的绝对路径",
			yiuLang.EnUS: "Edit software '" + e.Path + "' is not an absolute path to a valid file",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	e.Status = enum.ObjStatusValid
	return nil
}

func (e *EditSoft) CheckName(y yiuLang.YiuLang) error {
	if yiuStr.IsBlank(e.Name) {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "编辑软件名称不能为空",
			yiuLang.EnUS: "The name of the editing software cannot be empty",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	return nil
}

func (e *EditSoft) Check(y yiuLang.YiuLang) error {
	return yiuSErr.ToErrorBySep(" & ",
		e.CheckPath(y),
		e.CheckName(y),
	)
}
