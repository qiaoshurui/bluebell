package controller

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//获取参数和参数校验
	var p models.ParamSignUP
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//请求err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist) //用户已存在
			return
		}
		ResponseError(c, CodeServerBusy) //服务繁忙
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
func LoginHandler(c *gin.Context) {
	//请求参数和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		//请求err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNoExist) {
			ResponseError(c, CodeUserNotExist) //用户不存在
			return
		}
		ResponseError(c, CodeInvalidPassword) //用户名或密码错误
		return
	}
	//返回响应
	ResponseSuccess(c, nil)

}
