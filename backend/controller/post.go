package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/webforum/logic"
	"github.com/webforum/models"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数以及校验
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		zap.L().Debug("c.ShouldBindJSON(post) err", zap.Any("err", err))
		zap.L().Error("CREATE POST with invalid param")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	// 2. 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorID = userID
	// 3. 创建帖子
	err = logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 4. 返回响应
	ResponseSuccess(c, nil)

}
