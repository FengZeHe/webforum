package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/webforum/dao/mysql"
	"github.com/webforum/logic"
	"github.com/webforum/models"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	fmt.Println("SingUpHandler ...")
	//1. 获取请求参数
	var fo *models.RegisterForm
	//2.校验数据有效性
	if err := c.ShouldBindJSON(&fo); err != nil {
		// 如果请求参数有误，返回响应
		fmt.Println("请求参数有误", err)
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParams) // 请求参数错误
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return // 翻译错误
	}
	fmt.Printf("fo: %v\n", fo)

	// 3.业务处理 -- 注册用户
	if err := logic.SignUp(fo); err != nil {
		fmt.Println("注册业务失败", err)
		zap.L().Error("logic.signup failed", zap.Error(err))
		if err.Error() == mysql.ErrorUserExit {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return

	}
	// 返回正确相应
	ResponseSuccess(c, nil)

}
