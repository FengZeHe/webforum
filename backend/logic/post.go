package logic

import (
	"fmt"
	"github.com/webforum/dao/mysql"
	"github.com/webforum/dao/redis"
	"github.com/webforum/models"
	"github.com/webforum/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	//1. 生成post_id
	postID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}

	post.PostID = postID

	// 2. 创建帖子并保存到数据库
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}

	community, err := mysql.GetCommunityNameByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
		return err
	}

	// redis 存储帖子信息
	if err := redis.CreatePost(
		post.PostID,
		post.AuthorID,
		post.Title,
		TruncateByWords(post.Content, 120),
		community.CommunityID); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return
}
