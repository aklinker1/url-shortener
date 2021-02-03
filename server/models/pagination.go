package models

type Pagination struct {
	Page int
	Size int
}

func (it *Pagination) Limit() int {
	return it.Size
}

func (it *Pagination) Offset() int {
	return it.Size * it.Page
}
