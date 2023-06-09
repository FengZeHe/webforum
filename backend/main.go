package main

import (
	"fmt"
	"github.com/webforum/dao/mysql"
	"github.com/webforum/dao/redis"
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

	// 注册路由
	r := routers.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
