package mysql

import "github.com/pkg/errors"

func CheckUserExist(username string) (error error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
	}
	return
}
