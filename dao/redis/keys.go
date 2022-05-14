package redis

//redis key
//注意使用命名空间的方式方便查询和区分
const (
	Prefix              = "bluebell:"
	KeyPostTImeZSet     = "post:time"   //zset：帖子以及发帖的时间
	keyPostScoreZSet    = "post:score"  //zset：帖子以及投票的分数
	LKeyPostVotedZSetPF = "post:voted:" //zset：记录用户以及投票类型,参数是post id
	KeyCommunitySetPF   = "community:"  //set;保存每个分区下帖子的id
)

//给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
