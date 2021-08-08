package vo

type DbPageSearchVo struct {
	Data        []DbPageSearchItemVo `json:"data"`
	TotalRow    int                  `json:"totalRow"`
	TotalPage   int                  `json:"totalPage"`
	CurrentPage int                  `json:"currentPage"`
	PageSize    int                  `json:"pageSize"`
	HasPrevious bool                 `json:"hasPrevious"`
	HasNext     bool                 `json:"hasNext"`
}
