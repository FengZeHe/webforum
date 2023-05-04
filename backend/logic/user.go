package logic

import (
	"github.com/webforum/dao/mysql"
	"github.com/webforum/models"
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
