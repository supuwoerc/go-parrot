package service

import (
	"go-parrot/src/dao"
	"go-parrot/src/model"
)

type RequestService struct {
	BasicService
	Dao *dao.RequestDao
}

var requestService *RequestService

func NewRequestService() *RequestService {
	if requestService == nil {
		requestService = &RequestService{
			BasicService: NewBasicService(),
			Dao:          dao.NewRequestDao(),
		}
	}
	return requestService
}

func (r *RequestService) RequestRecordAdd(dto model.Request) error {
	return r.Dao.RequestRecordAdd(dto)
}
