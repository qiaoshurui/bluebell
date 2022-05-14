package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUP) (err error) {
	//判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		//查询数据库出错
		return err
	}
	//生成UID
	userID := snowflake.GenID()
	//构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	mysql.InsertUser(user)
	return
}
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return " ", err
	}
	//生成jwt的token
	return jwt.GenToken(user.UserID, user.Username)
}
