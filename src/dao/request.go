package dao

import "go-parrot/src/model"

type RequestDao struct {
	BasicDao
}

var requestDao *RequestDao

func NewRequestDao() *RequestDao {
	if requestDao == nil {
		requestDao = &RequestDao{NewBasicDao()}
	}
	return requestDao
}

// 添加一条请求记录
func (r *RequestDao) RequestRecordAdd(dto model.Request) error {
	return r.Orm.Model(&model.Request{}).Save(&dto).Error
}
