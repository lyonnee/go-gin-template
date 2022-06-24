package mysql

import (
	"context"
	"database/sql"

	"time"

	"github.com/LyonNee/app-layout/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var db *sql.DB

func Connection() {
	var err error
	db, err = sql.Open("mysql",
		viper.GetString("mysql.dsn"))
	if err != nil {
		log.ZapLogger().Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.ZapLogger().Fatal(err.Error())
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
}

func Disconnect() {
	db.Close()
}

func GetConn(ctx context.Context) (*sql.Conn, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		log.ZapLogger().Error("获取数据库conn失败,err: " + err.Error())
		return nil, err
	}
	return conn, nil
}

//@developer  仅用于单测时Mock数据库
func MockDB(mockDB *sql.DB) {
	db = mockDB
}
