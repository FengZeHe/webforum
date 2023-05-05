package redis

import (
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600 // 一周的秒数
	VoteScore        float64 = 432           // 每一票的值432分
	PostPerAge               = 20            // 每页显示20条帖子
)

// CreatePost redis存储帖子信息，使用hash存储帖子信息
func CreatePost(postID, userID uint64, title, summary string, CommunityID uint64) (err error) {
	now := float64(time.Now().Unix())
	votedKey := KeyPostVotedZSetPrefix + strconv.Itoa(int(postID))
	communityKey := KeyCommunityPostSetPrefix + strconv.Itoa(int(CommunityID))
	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    0,
		"comments": 0,
	}
	// 事务操作
	pipeline := client.TxPipeline()
	pipeline.ZAdd(ctx, votedKey, redis.Z{ // 作者默认投赞成票
		Score:  1,
		Member: userID,
	})

	pipeline.Expire(ctx, votedKey, time.Second*OneWeekInSeconds)

	// 缓存 hash 文章
	pipeline.HMSet(ctx, KeyPostInfoHashPrefix+strconv.Itoa(int(postID)), postInfo)

	// 添加分数  Zset
	pipeline.ZAdd(ctx, KeyPostScoreZSet, redis.Z{
		Score:  now + VoteScore,
		Member: postID,
	})

	// 添加时间 Zset
	pipeline.ZAdd(ctx, KeyPostTimeZSet, redis.Z{
		Score:  now,
		Member: postID,
	})

	// 添加到对应板块
	pipeline.SAdd(ctx, communityKey, postID)
	_, err = pipeline.Exec(ctx)
	return
}
