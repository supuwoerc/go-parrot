package dao

import "go-parrot/src/model"

type UserDao struct {
	BasicDao
}

var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{NewBasicDao()}
	}
	return userDao
}

// 根据用户名和密码查询用户
func (u *UserDao) GetUserByNameAndPassword(name string, password string) (model.User, error) {
	var user model.User
	err := u.Orm.Model(&model.User{}).Where("name = ? and password = ?", name, password).First(&user).Error
	return user, err
}
