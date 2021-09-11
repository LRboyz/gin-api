package initialize

import (
	"github.com/LRboyz/gin-api/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() *gorm.DB {
	mysqlInfo := global.Settings.MysqlInfo
	dsn := mysqlInfo.Name + ":" + mysqlInfo.Password + "@tcp(" + mysqlInfo.Host + ")/" + mysqlInfo.DBName + "?" + "charset=utf8&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if db, err := gorm.Open(mysql.New(mysqlConfig), config); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		return db
	}
}
