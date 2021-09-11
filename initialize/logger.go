package initialize

import (
	"fmt"

	"github.com/LRboyz/gin-api/global"
	"github.com/LRboyz/gin-api/utils"
	"go.uber.org/zap"
)

func InitLogger() {
	// 实例化zap配置
	cfg := zap.NewDevelopmentConfig()
	// 配置日志的输出地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()),
	}

	logg, _ := cfg.Build()
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logg)
	global.Lg = logg
}
