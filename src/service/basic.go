package service

import (
	"go-parrot/src/conf"
	"go-parrot/src/global"
)

type BasicService struct {
	RedisClient *conf.RedisClient
}

func NewBasicService() BasicService {
	return BasicService{
		RedisClient: global.RedisClient,
	}
}
