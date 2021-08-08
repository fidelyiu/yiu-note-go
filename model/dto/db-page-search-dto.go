package dto

import FieldUtil "github.com/fidelyiu/yiu-note/core/field"

type DbPageSearchDto struct {
	Key      string   `json:"key" form:"key"`
	Str      []string `json:"str" form:"str"`
	Db       string   `json:"db" form:"db"`
	PageSize int      `json:"pageSize" form:"pageSize"`
	Page     int      `json:"page" form:"page"`
}

func (d *DbPageSearchDto) CheckDb() {
	switch d.Db {
	case FieldUtil.MainTable:
	case FieldUtil.LayoutTable:
	case FieldUtil.WorkspaceTable:
	case FieldUtil.NoteTable:
	case FieldUtil.ImageCacheTable:
	case FieldUtil.EditSoftTable:
	default:
		d.Db = FieldUtil.MainTable
	}
}

func (d *DbPageSearchDto) CheckPageSize() {
	if d.PageSize <= 0 {
		d.PageSize = 10
	}
}

func (d *DbPageSearchDto) CheckPage() {
	if d.Page <= 0 {
		d.Page = 1
	}
}

func (d *DbPageSearchDto) Check() {
	d.CheckDb()
	d.CheckPageSize()
	d.CheckPage()
}
