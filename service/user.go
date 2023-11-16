package service

import (
	"chagic/model"
	"chagic/pkg/e"
	"chagic/serializer"
)

type UserService struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *UserService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	model.DB.Where("username=?", r.Username).Find(&user)
	if user.Username == "" {
		code = e.ERROR_NOT_EXIST_USER
		return serializer.ResultErr(code)
	}
	if err := user.VerifyPassword(r.Password); err != nil {
		code = e.ERROR_PASSWORD_WRONG
		return serializer.ResultErr(code)
	}
	user.Password = ""
	token, err := user.GenerateToken()
	if err != nil {
		code = e.ERROR_TOKEN_GENERATE_FAILED
		return serializer.ResultErr(code)
	}
	return serializer.ResultOK(serializer.TokenData{
		User:  user,
		Token: token,
	})
}

func (r *UserService) Register() serializer.Response {
	var user model.User
	var count int64
	code := e.SUCCESS
	model.DB.Model(&model.User{}).Where("username=?", r.Username).Count(&count)
	if count > 0 {
		code = e.ERROR_EXIST_USERNAME
		return serializer.ResultErr(code)
	}

	user = model.User{
		Username: r.Username,
	}
	user.HashPassword(r.Password)
	if err := model.DB.Create(&user).Error; err != nil {
		code = e.ErrorDatabase
		return serializer.ResultErr(code)
	}
	user.Password = ""
	token, err := user.GenerateToken()
	if err != nil {
		code = e.ERROR_TOKEN_GENERATE_FAILED
		return serializer.ResultErr(code)
	}
	return serializer.ResultOK(serializer.TokenData{
		User:  user,
		Token: token,
	})
}

func GetUserInfo(id float64) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	user.Password = ""
	return serializer.ResultOK(user)
}
