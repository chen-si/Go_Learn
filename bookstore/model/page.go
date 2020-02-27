package model

type Page struct {
	Books       []*Book //每页查询出来的图书存放的页数
	PageNo      int64   //当前页码
	PageSize    int64   //每页显示的条数
	TotalPageNo int64   //总页数 计算得到
	Totalrecord int64   //总记录数 查询数据库得到
	MinPrice    string
	MaxPrice    string
}

//IsHasPrev
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNo
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

//GetNextPageNo
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
