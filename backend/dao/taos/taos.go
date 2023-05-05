package taos

import (
	"database/sql"
	"fmt"
	"github.com/webforum/setting"

	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

var taos *sql.DB

func Init(cfg *setting.TaosConfig) (err error) {
	var taosDNS = fmt.Sprintf("%s:%s@http(%s:%s)/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	taos, err = sql.Open("taosRestful", taosDNS)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return
	}
	fmt.Println("taos init success!")
	return
}

func Close() {
	taos.Close()
}
