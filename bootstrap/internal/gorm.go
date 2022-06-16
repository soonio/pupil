package internal

import (
	"errors"
	"github.com/soonio/pupil/app"
	"github.com/soonio/pupil/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() {
	app.DB = connect(app.Config.DB)
}

func connect(config *config.DB) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       config.Dsn(), // DSN data source name
		DefaultStringSize:         191,          // string 类型字段的默认长度
		DisableDatetimePrecision:  true,         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,        // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		panic(errors.New("db connect failed, err:" + err.Error()))
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
		return db
	}
}
