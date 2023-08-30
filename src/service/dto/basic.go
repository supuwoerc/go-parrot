package dto

type BasicIdDTO struct {
	ID uint `json:"id" form:"id" uri:"id" binding:"required"`
}

type Paginate struct {
	Page     int `json:"page,omitempty" form:"page"`
	PageSize int `json:"page_size,omitempty" form:"page_size"`
}

func (p Paginate) GetPage() int {
	if p.Page > 0 {
		return p.Page
	}
	return 1
}

func (p Paginate) GetPageSize() int {
	if p.PageSize <= 0 {
		return 10
	}
	return p.PageSize
}
