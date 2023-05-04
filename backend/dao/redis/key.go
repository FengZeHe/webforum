package redis

const (
	KeyPostInfoHashPrefix = "webforum:post:"
	KeyPostTimeZSet       = "webforum:post:time"  // zset;帖子及发帖时间定义
	KeyPostScoreZSet      = "webforum:post:score" // zset;帖子及投票分数定义
	//KeyPostVotedUpSetPrefix   = "webforum:post:voted:down:"
	//KeyPostVotedDownSetPrefix = "webforum:post:voted:up:"
	KeyPostVotedZSetPrefix    = "webforum:post:voted:" // zSet;记录用户及投票类型;参数是post_id
	KeyCommunityPostSetPrefix = "webforum:community:"  // set保存每个分区下帖子的id
)
