package rwsql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormConfig struct {
	Source string
}

func NewGormDb(config *GormConfig) (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(config.Source), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		messqge := fmt.Errorf("failed opening connection to db: %v", err)
		panic(messqge)
	}

	mysql, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	mysql.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	mysql.SetMaxOpenConns(100)
	mysql.Stats()

	return
}
