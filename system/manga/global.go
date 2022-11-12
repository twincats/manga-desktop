package manga

import "gorm.io/gorm"

var DB *gorm.DB

// Pagination struct for pagination need set Total
type Pagination struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	To          int `json:"to"`
	Total       int `json:"total"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"perpage"`
	Offset      int `json:"-"`
}

//set value for pagination struct
func (p *Pagination) Paginate() {
	//set From
	p.From = p.PerPage*p.CurrentPage - (p.PerPage - 1)

	//set To
	p.To = p.CurrentPage * p.PerPage
	if p.To > p.Total {
		p.To = p.Total
	}

	//set LastPage
	p.LastPage = p.Total / p.PerPage
	if p.LastPage > 0 {
		p.LastPage += 1
	}

	//limit currentpage to available page
	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}

	if p.CurrentPage > p.LastPage {
		p.CurrentPage = p.LastPage
	}

	//setting Offset for DB Query
	p.Offset = (p.CurrentPage - 1) * p.PerPage

}
