package main

import "gorm.io/gorm"

type Paginator struct {
	DefaultSize int
}

func NewPaginator(defaultSize int) *Paginator {
	return &Paginator{DefaultSize: defaultSize}
}

func (p *Paginator) Paginate(page, size int) func(s *gorm.Statement) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = p.DefaultSize
	}
	return func(s *gorm.Statement) {
		s.Offset((page - 1) * size).Limit(size)
	}
}
