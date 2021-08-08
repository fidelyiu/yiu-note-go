package entity

import (
	"errors"
	yiuSErr "github.com/fidelyiu/yiu-go-tool/error_s"
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
	langCore "github.com/fidelyiu/yiu-note/core/lang"
	"github.com/fidelyiu/yiu-note/model/enum"
	"time"
)

type Layout struct {
	Id         string                 `json:"id"`         // Uuid
	Type       enum.LayoutType        `json:"type"`       // 类型
	Status     enum.ObjStatus         `json:"status"`     // 状态
	Width      int                    `json:"width"`      // 宽度
	Height     int                    `json:"height"`     // 高度
	Left       int                    `json:"left"`       // 距离左边
	Top        int                    `json:"top"`        // 距离顶部
	Setting    map[string]interface{} `json:"setting"`    // 设置，根据类型而定
	UpdateTime time.Time              `json:"updateTime"` // 最后更新时间
}

func (l *Layout) CheckType(y yiuLang.YiuLang) error {
	if l.Type <= enum.LayoutTypeLink-1 || l.Type >= enum.LayoutTypeMainBox+1 {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "布局[Type]无效",
			yiuLang.EnUS: "Layout [Type] is invalid",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	return nil
}

func (l *Layout) CheckStatus(y yiuLang.YiuLang) error {
	if l.Status <= enum.ObjStatusInvalid-1 || l.Status >= enum.ObjStatusValid+1 {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "布局状态无效",
			yiuLang.EnUS: "Layout status is invalid",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	return nil
}

func (l *Layout) CheckWidth(y yiuLang.YiuLang) error {
	if l.Width <= 0 {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "布局宽度无效",
			yiuLang.EnUS: "Layout width is invalid",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	return nil
}

func (l *Layout) CheckHeight(y yiuLang.YiuLang) error {
	if l.Width <= 0 {
		lMap := map[yiuLang.YiuLang]string{
			yiuLang.ZhCN: "布局高度无效",
			yiuLang.EnUS: "Layout height is invalid",
		}
		return errors.New(langCore.GetLangByKey(y, lMap))
	}
	return nil
}

func (l *Layout) CheckSetting(y yiuLang.YiuLang) error {
	switch l.Type {
	case enum.LayoutTypeLink:
		if l.Setting["name"] == nil {
			lMap := map[yiuLang.YiuLang]string{
				yiuLang.ZhCN: "链接名称不能为空",
				yiuLang.EnUS: "Link name cannot be empty",
			}
			return errors.New(langCore.GetLangByKey(y, lMap))
		}
		if l.Setting["url"] == nil {
			lMap := map[yiuLang.YiuLang]string{
				yiuLang.ZhCN: "链接地址不能为空",
				yiuLang.EnUS: "Link address cannot be empty",
			}
			return errors.New(langCore.GetLangByKey(y, lMap))
		}
	}
	return nil
}

func (l *Layout) Check(y yiuLang.YiuLang) error {
	return yiuSErr.ToErrorBySep(" & ",
		l.CheckType(y),
		l.CheckStatus(y),
		l.CheckWidth(y),
		l.CheckHeight(y),
	)
}
