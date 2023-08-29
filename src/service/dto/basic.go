package dto

type BasicIdDTO struct {
	ID uint `json:"id" form:"id" uri:"id" binding:"required"`
}
