package dao

import (
	"go-parrot/src/global"
	"go-parrot/src/service/dto"
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

// 通用的分页参数处理scope
func Paginate(dto dto.Paginate) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := dto.GetPage()
		size := dto.GetPageSize()
		return db.Offset((page - 1) * size).Limit(size)
	}
}
