package model

//分页
type Page struct {
	Records interface{} `json:"records"`
	Total   int64       `josn:"total"`
	PageNum int         `json:"pageNum"`
	PageInfo
}

//分页参数
type PageInfo struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

func (p *Page) SetPageNum(count int64) {
	c := int(count)
	if c%p.PageSize == 0 {
		p.PageNum = c / p.PageSize
	} else {
		p.PageNum = c/p.PageSize + 1
	}
}
