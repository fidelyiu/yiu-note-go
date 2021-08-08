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

type Workspace struct {
	Id      string         `json:"id"`      // Uuid
	Path    string         `json:"path"`    // 绝对路径
	Name    string         `json:"name"`    // 名称
	Alias   string         `json:"alias"`   // 别名
	SortNum int            `json:"sortNum"` // 排序数
	Status  enum.ObjStatus `json:"status"`  // 状态
}

func (w *Workspace) CheckPath(y yiuLang.YiuLang) error {
	if !PathUtil.IsValidDir(w.Path) {
		w.Status = enum.ObjStatusInvalid
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "工作空间 '" + w.Path + "' 不是有效绝对路径",
			yiuLang.EnUS: "Workspace '" + w.Path + "' is not a valid absolute path",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	w.Status = enum.ObjStatusValid
	return nil
}

func (w *Workspace) CheckName(y yiuLang.YiuLang) error {
	if yiuStr.IsBlank(w.Name) {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "工作空间名称不能为空",
			yiuLang.EnUS: "Workspace name cannot be empty",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	return nil
}

func (w *Workspace) Check(y yiuLang.YiuLang) error {
	return yiuSErr.ToErrorBySep(" & ",
		w.CheckPath(y),
		w.CheckName(y),
	)
}
