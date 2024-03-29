package dao

import (
	"dario.cat/mergo"
	"errors"
	"fmt"
	"go-parrot/src/model"
	"go-parrot/src/service/dto"
)

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

// 添加用户
func (u *UserDao) AddUser(dto *dto.UserAddDTO) error {
	var modelUser model.User
	err := mergo.Merge(&modelUser, model.User{
		Name:     dto.Name,
		RealName: dto.RealName,
		Avatar:   dto.Avatar,
		Phone:    dto.Phone,
		Email:    dto.Email,
		Password: dto.Password,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("参数合并发生错误：%s", err.Error()))
	}
	err = u.Orm.Save(&modelUser).Error
	if err == nil {
		dto.Password = ""     //移除password
		dto.ID = modelUser.ID //数据库中的id
	}
	return err
}

// 根据用户名查询用户
func (u *UserDao) GetUserByName(name string) (model.User, error) {
	var modelUser model.User
	err := u.Orm.Model(&model.User{}).Where("name = ?", name).First(&modelUser).Error
	return modelUser, err
}

// 根据用户ID查询用户
func (u *UserDao) GetUserById(id uint) (model.User, error) {
	var modelUser model.User
	err := u.Orm.Model(&model.User{}).Where("id = ?", id).First(&modelUser).Error
	return modelUser, err
}

// 查询用户列表
func (u *UserDao) GetUserList(dto dto.UserListDTO) ([]model.User, int64, error) {
	var userList []model.User
	var total int64
	db := u.Orm.Model(&model.User{})
	if dto.Name != "" {
		db.Where("name like ?", fmt.Sprintf("%%%s%%", dto.Name))
	}
	err := db.Scopes(Paginate(dto.Paginate)).Find(&userList).Offset(-1).Limit(-1).Count(&total).Error
	return userList, total, err
}

// 更新用户信息
func (u *UserDao) UpdateUser(dto *dto.UserUpdateDTO) error {
	targetUser, err := u.GetUserById(dto.ID)
	if err != nil {
		return err
	} else {
		sourceUser := model.User{
			Model:    targetUser.Model,
			RealName: dto.RealName,
			Avatar:   dto.Avatar,
			Phone:    dto.Phone,
			Email:    dto.Email,
			Password: dto.Password,
		}
		err := mergo.Merge(&targetUser, sourceUser, mergo.WithOverride)
		if err != nil {
			return err
		}
		return u.Orm.Save(&targetUser).Error
	}
}

func (u *UserDao) DeleteUserById(id uint) error {
	err := u.Orm.Delete(&model.User{}, id).Error
	return err
}
