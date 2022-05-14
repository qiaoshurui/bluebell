package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) error {
	//事物处理
	pipeline := client.Pipeline()
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTImeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//帖子分数
	pipeline.ZAdd(getRedisKey(keyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}
func VoteForPost(userID, postID string, value float64) error {
	//1. 判断投票的限制
	//去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTImeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//2和3需要放在一个pipeline事物中处理
	//2. 更新分数
	//先查当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(LKeyPostVotedZSetPF+postID), userID).Val()
	var dir float64
	//如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value)
	pipeline := client.Pipeline()
	pipeline.ZIncrBy(getRedisKey(keyPostScoreZSet), dir*diff*scorePerVote, postID)
	//3. 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(LKeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(LKeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
