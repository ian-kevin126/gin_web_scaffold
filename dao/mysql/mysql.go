package mysql

import (
	"bluebell/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(mysqlConf *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DB,
	)

	// 也可以使用MustConnect连接，不成功就panic
	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		zap.L().Error("Connect DB failed", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(mysqlConf.MaxOpenConns)
	db.SetMaxIdleConns(mysqlConf.MaxIdleConns)

	return
}

// Close 对外暴露一个close数据库连接的方法
func Close() {
	_ = db.Close()
}
