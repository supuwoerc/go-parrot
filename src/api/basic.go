package api

import (
	"go-parrot/src/global"
	"go.uber.org/zap"
)

type BasicApi struct {
	Logger *zap.SugaredLogger
}

func NewBasicApi() BasicApi {
	return BasicApi{Logger: global.Logger}
}
