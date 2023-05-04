package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/webforum/controller"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}

	// 初始化 gin Engine
	r := gin.New()
	v1 := r.Group("/api/v1")

	// 登录注册业务
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	return r
}
