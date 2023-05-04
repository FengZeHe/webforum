package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/webforum/models"
)

const secret = "web.forum"

// 对密码进行加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

func CheckUserExist(username string) (error error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New(ErrorUserExit)
	}
	return
}

// InsertUser 注册业务-向数据库中插入新的用户
func InsertUser(user models.User) (error error) {
	// 对密码进行加密
	user.Password = encryptPassword([]byte(user.Password))
	// 执行SQL语句
	sqlSQL := `insert into user(user_id,username,password,email,gender) values (?,?,?,?,?)`
	_, err := db.Exec(sqlSQL, user.UserID, user.UserName, user.Password, user.Email, user.Gender)
	return err
}

// Login 登录业务
func Login(user *models.User) (err error) {
	// 记录一下用户登录密码
	originPassword := user.Password
	sqlStr := `select user_id,username,password from user where username = ?`
	err = db.Get(user, sqlStr, user.UserName)
	// 查询数据库异常
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	// 如果用户不存在
	if err == sql.ErrNoRows {
		return errors.New(ErrorUserNotExit)
	}
	// 生成加密密码与查询到的秘密进行比对
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return errors.New(ErrorPasswordWrong)
	}
	return nil

}
