package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
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
func Login(p *models.ParamLogin) error {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
