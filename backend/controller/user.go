package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/webforum/dao/mysql"
	"github.com/webforum/logic"
	"github.com/webforum/models"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	fmt.Println("hello this is SignUpHandler")
	//1. 获取请求参数
	var fo *models.RegisterForm
	//2.校验数据有效性
	if err := c.ShouldBindJSON(&fo); err != nil {
		// 如果请求参数有误，返回响应&
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		return
	}
	fmt.Printf("fo: %v\n", fo)

	// 3.业务处理 -- 注册用户
	if err := logic.SignUp(fo); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))
		if err.Error() == mysql.ErrorUserExit {

		}

	}

}
