package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-parrot/src/constant"
	"go-parrot/src/dao"
	"go-parrot/src/global"
	"go-parrot/src/model"
	"go-parrot/src/service/dto"
	"go-parrot/src/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
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

// 存储user的token到redis
func SetLoginUserToken2Redis(id uint, token string) error {
	expires := viper.GetDuration("jwt.expires") * time.Minute
	redisKey := strings.Replace(constant.LOGIN_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(id)), -1)
	return global.RedisClient.Set(redisKey, token, expires)
}

// 判断当前请求的用户
func IsCurrentUser(ctx *gin.Context, id uint) bool {
	loginUser := model.LoginUser{}
	value, exists := ctx.Get(constant.LOGIN_USER_KEY)
	if exists {
		loginUser.ID = value.(model.LoginUser).ID
		return loginUser.ID == id
	}
	return false
}

// 用户登录
func (u *UserService) Login(dto dto.UserLoginDTO) (model.User, string, error) {
	user, err := u.Dao.GetUserByName(dto.Name)
	if err != nil || !utils.CompareHashAndPassword(dto.Password, user.Password) {
		return user, "", errors.New("用户名或账户错误")
	} else {
		token, err := utils.GenerateToken(user.ID, user.Name)
		err = SetLoginUserToken2Redis(user.ID, token)
		return user, token, err
	}
}

// 添加用户
func (u *UserService) AddUser(dto *dto.UserAddDTO) error {
	_, err := u.Dao.GetUserByName(dto.Name)
	if err == gorm.ErrRecordNotFound {
		err = u.Dao.AddUser(dto)
		return err
	}
	return errors.New(fmt.Sprintf("用户名【%s】已存在,添加失败", dto.Name))
}

// 根据ID查询用户
func (u *UserService) GetUserById(dto *dto.BasicIdDTO) (model.User, error) {
	modelUser, err := u.Dao.GetUserById(dto.ID)
	if err == gorm.ErrRecordNotFound {
		return model.User{}, errors.New(fmt.Sprintf("未发现ID为%d的用户", dto.ID))
	}
	modelUser.Password = ""
	return modelUser, nil
}

// 查询用户列表
func (u *UserService) GetUserList(dto *dto.UserListDTO) ([]model.User, int64, error) {
	return u.Dao.GetUserList(*dto)
}

// 修改用户信息
func (u *UserService) UpdateUser(dto *dto.UserUpdateDTO) error {
	//TODO：添加权限控制，只允许用户修改自己的账户信息
	return u.Dao.UpdateUser(dto)
}

// 删除用户
func (u *UserService) DeleteUser(id uint) error {
	//TODO：添加权限控制，只允许用户删除自己的账户
	return u.Dao.DeleteUserById(id)
}
