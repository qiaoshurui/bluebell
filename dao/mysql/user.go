package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNoExist     = errors.New("用户不存在")
	ErrorInvalidPassWord = errors.New("用户名或密码错误")
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil

}
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) value(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

//密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	//h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登陆的密码
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username) //将从数据库中获得数据赋给user数据
	if err == sql.ErrNoRows {
		return ErrorUserNoExist
	}
	if err != nil {
		//查询数据库错误
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassWord
	}
	return
}
