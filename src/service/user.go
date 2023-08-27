package service

import (
	"errors"
	"go-parrot/src/dao"
	"go-parrot/src/model"
	"go-parrot/src/service/dto"
	"go-parrot/src/utils"
	"gorm.io/gorm"
)

type UserService struct {
	BasicService
	Dao *dao.UserDao
}

var userService *UserService

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			BasicService: NewBasicService(),
			Dao:          dao.NewUserDao(),
		}
	}
	return userService
}

func (u *UserService) Login(dto dto.UserLoginDTO) (model.User, string, error) {
	user, err := u.Dao.GetUserByNameAndPassword(dto.Name, dto.Password)
	if err == gorm.ErrRecordNotFound {
		return user, "", errors.New("用户名或账户错误")
	} else {
		token, err := utils.GenerateToken(user.ID, user.Name)
		//TODO:调整token写入redis策略
		err = u.RedisClient.Set(user.Name, token)
		return user, token, err
	}
}
