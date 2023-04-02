package pagination

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

const MaxPerPage = 100

type Pagination struct {
	Page         int    `json:"page,omitempty" query:"currentPage"`
	ItemsPerPage int    `json:"perPage,omitempty" query:"itemsPerPage"`
	TotalPages   int    `json:"totalPages"`
	TotalItems   int64  `json:"totalItems"`
	Sort         string `json:"sort,omitempty" query:"sort"`
}

// Interface declaration is only needed to automate mock generation.
// This interface is not relevant anywhere else.
type IPagination interface {
	GetOffset() int
	GetLimit() int
	GetPage() int
	GetSort() string
	Paginate(query *gorm.DB) func(db *gorm.DB) *gorm.DB
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.ItemsPerPage <= 0 {
		p.ItemsPerPage = 10
	}
	if p.ItemsPerPage > MaxPerPage {
		p.ItemsPerPage = MaxPerPage
	}
	return p.ItemsPerPage
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func (p *Pagination) Paginate(query *gorm.DB) func(db *gorm.DB) *gorm.DB {
	query.Where(`"deleted_at" is NULL`).Count(&p.TotalItems)

	p.TotalPages = int(math.Ceil(float64(p.TotalItems) / float64(p.GetLimit())))

	if p.GetPage() > p.TotalPages && p.TotalPages != 0 {
		return func(db *gorm.DB) *gorm.DB {
			db.AddError(fmt.Errorf("page input exceeds the maximum page limit, input:%v max-page:%v", p.Page, p.TotalPages))
			return db
		}
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}
}