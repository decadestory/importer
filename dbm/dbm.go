package dbm

import (
	"fmt"

	"gitlab.smartee.cn/chenguorui/goutil/exception"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDb(dstType string, dsn string) {
	switch dstType {
	case "mysql":
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		exception.Errors.CheckErr(err)
		DB = db
	case "postgres":
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		exception.Errors.CheckErr(err)
		DB = db
	case "clickhouse":
		db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		exception.Errors.CheckErr(err)
		DB = db
	default:
		fmt.Println("unsupported database type")
	}
}
