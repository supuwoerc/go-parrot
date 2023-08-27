package dao

import (
	"go-parrot/src/global"
	"gorm.io/gorm"
)

type BasicDao struct {
	Orm *gorm.DB
}

func NewBasicDao() BasicDao {
	return BasicDao{
		Orm: global.DB,
	}
}
