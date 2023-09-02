package api

import (
	"go-parrot/src/service"
)

type RequestApi struct {
	BasicApi
	Service *service.RequestService
}

var requestApi *RequestApi

func NewRequestApi() *RequestApi {
	if requestApi == nil {
		requestApi = &RequestApi{
			BasicApi: NewBasicApi(),
			Service:  service.NewRequestService(),
		}
	}
	return requestApi
}
