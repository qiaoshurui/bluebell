package controller

import (
	"strconv"
	"web_app/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CommunityHandler(c *gin.Context) {
	//查询所有的社区 （community_id,community_name） 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端报错暴露在外边
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//查询所有的社区 （community_id,community_name） 以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端报错暴露在外边
		return
	}
	ResponseSuccess(c, data)
}
