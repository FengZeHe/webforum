package main

import (
	"fmt"
	"github.com/webforum/dao/mysql"
	"github.com/webforum/dao/redis"
	"github.com/webforum/logger"
	"github.com/webforum/pkg/snowflake"
	"github.com/webforum/routers"
	"github.com/webforum/setting"
)

func main() {
	if err := setting.Init(); err != nil {
		fmt.Println("load config file error %v", err)
	}

	if err := mysql.Init(setting.Conf.MysqlConfig); err != nil {
		fmt.Println("mysql init failed; err:%v", err)
	}
	defer mysql.Close()

	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Println("redis init failed; err:%v", err)
	}
	defer redis.Close()

	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Println("Log init failed; err:%v", err)
		return
	}

	// 雪花算法生成分布式ID
	if err := snowflake.Init(1); err != nil {
		fmt.Println("init snowflake failed, err:%v", err)
		return
	}

	// 注册路由
	r := routers.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
