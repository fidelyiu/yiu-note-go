package entity

import (
	"errors"
	yiuSErr "github.com/fidelyiu/yiu-go-tool/error_s"
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
	langCore "github.com/fidelyiu/yiu-note/core/lang"
	PathUtil "github.com/fidelyiu/yiu-note/core/path"
	"github.com/fidelyiu/yiu-note/model/enum"
)

type Note struct {
	Id               string         `json:"id"`               // Uuid
	AbsPath          string         `json:"absPath"`          // 绝对路径
	Path             string         `json:"path"`             // 相对于工作空间相对路径
	Name             string         `json:"name"`             // 名称
	Alias            string         `json:"alias"`            // 别名
	SortNum          int            `json:"sortNum"`          // 排序数
	Status           enum.ObjStatus `json:"status"`           // 状态
	WorkspaceId      string         `json:"workspaceId"`      // 所属工作空间Id
	ParentId         string         `json:"parentId"`         // 父级目录Id
	ParentAbsPath    string         `json:"parentAbsPath"`    // 父级的绝对路径
	ParentPath       string         `json:"parentPath"`       // 相对于父级的相对路径
	Level            int            `json:"level"`            // 等级
	Show             bool           `json:"show"`             // 是否展示
	IsDir            bool           `json:"isDir"`            // 是否是文件夹
	ShowDocumentNum  bool           `json:"showDocumentNum"`  // 笔记页面是否展示文档序号
	ShowDirNum       bool           `json:"showDirNum"`       // 笔记页面是否展示目录序号
	ShowMainPointNum bool           `json:"showMainPointNum"` // 笔记页面是否展示大纲序号
	// ShowNum       int                   `json:"showNum"`       // 排除隐藏文件后的标题顺序
	// DefStatus     enum.DefinitionStatus `json:"defStatus"`     // 是否定义过顺序，如果没定义过顺序，那么就是本地刚导入的
}

func (n *Note) CheckPath(y yiuLang.YiuLang) error {
	if n.IsDir {
		if !PathUtil.IsValidDir(n.AbsPath) {
			n.Status = enum.ObjStatusInvalid
			lMap := map[yiuLang.YiuLang]string{
				yiuLang.ZhCN: "工作空间 '" + n.AbsPath + "' 不是有效文件夹的绝对路径",
				yiuLang.EnUS: "Workspace '" + n.AbsPath + "' is not an absolute path to a valid folder",
			}
			return errors.New(langCore.GetLangByKey(y, lMap))
		}
	} else {
		if !PathUtil.IsValidFile(n.AbsPath) {
			n.Status = enum.ObjStatusInvalid
			lMap := map[yiuLang.YiuLang]string{
				yiuLang.ZhCN: "工作空间 '" + n.AbsPath + "' 不是有效文件的绝对路径",
				yiuLang.EnUS: "Workspace '" + n.AbsPath + "' is not an absolute path to a valid file",
			}
			return errors.New(langCore.GetLangByKey(y, lMap))
		}
	}
	n.Status = enum.ObjStatusValid
	return nil
}

func (n *Note) Check(y yiuLang.YiuLang) error {
	return yiuSErr.ToErrorBySep(" & ",
		n.CheckPath(y),
	)
}
