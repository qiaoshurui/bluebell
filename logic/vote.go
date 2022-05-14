package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"

	"go.uber.org/zap"
)

//投票功能：
/*
投票的几种情况：

一.direction=1时，有两种情况：
	1.之前没有投过票，现在投赞票       -->更新分数和投票记录
	2.之前投反对票，现在改投赞成票     -->更新分数和投票记录

二.direction=0时，有两种情况：        -->更新分数和投票记录
	1.之前投过赞成票，现在投反对票		-->更新分数和投票记录

	2.之前投过反对票，现在要取消投票		-->更新分数和投票记录
三.direction=-1时，有两种情况			-->更新分数和投票记录

	1.之前没有投过票，现在投反对票		-->更新分数和投票记录
	2.之前投赞成票，现在改投反对票		-->更新分数和投票记录

投票限制：
每个帖子自发表之日起一个星期之内允许用户进行投票，超过一个星期就不在允许投票了。
	1.到期之后将redis中保存的赞成票数及反对票数储存到mysql表中
	2.到期之后删除那个 keyPostVotedZSetPF
*/

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
