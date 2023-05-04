package mysql

import (
	"github.com/webforum/models"
	"go.uber.org/zap"
)

// 创建帖子，插入到数据库中
func CreatePost(post *models.Post) (err error) {
	// 写sql语句并执行
	sqlStr := `insert into post(post_id,title,content,author_id,community_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, post.PostID, post.Title, post.Content, post.AuthorID, post.CommunityID)

	if err != nil {
		zap.L().Error("INSERT post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return nil
}
