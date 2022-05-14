package controller

import (
	"web_app/logic"
	"web_app/models"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func PostVoteController(c *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //翻译报出的错误并且去掉前面的多余结构体
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userId, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeUserNotExist)
		return
	}
	logic.VoteForPost(userId, p)
	ResponseSuccess(c, nil)
}
