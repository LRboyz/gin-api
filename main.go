package main

import (
	"fmt"

	"github.com/LRboyz/gin-api/global"
	"github.com/LRboyz/gin-api/initialize"
	"github.com/fatih/color"
	"go.uber.org/zap"
)

func main() {
	// 1.初始化yaml配置
	initialize.InitConfig()
	// 2.初始化routers
	Router := initialize.Routers()
	// 3. 初始化日志信息
	initialize.InitLogger()
	// 4. 初始化翻譯
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	// 5. 初始化mysql
	initialize.InitMysqlDB()
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	color.Yellow("go-gin服务开始了")

	// 启动Gin 并配置端口
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("This is Hello func", zap.String("error", "启动错误"))
	}

}
