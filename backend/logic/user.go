package logic

import (
	"github.com/webforum/dao/mysql"
	"github.com/webforum/models"
	"github.com/webforum/pkg/jwt"
	"github.com/webforum/pkg/snowflake"
)

func SignUp(p *models.RegisterForm) (error error) {
	// 1. 判断用户是否存在
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		// 数据库查询 已经存在
		return err
	}

	// 生成UID
	userId, err := snowflake.GetID()
	if err != nil {
		return mysql.ErrorGenIDFailed
	}

	// 构造一个user实例
	u := models.User{
		UserID:   userId,
		UserName: p.UserName,
		Password: p.Password,
		Email:    p.Email,
		Gender:   p.Gender,
	}
	// 3. 将数据保存到数据库
	return mysql.InsertUser(u)

}

func Login(p *models.LoginForm) (user *models.User, err error) {
	user = &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}

	// 生成JWT
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	return

}
