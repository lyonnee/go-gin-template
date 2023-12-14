package mysql

import (
	"context"
	"database/sql"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lyonnee/go-gin-template/infra/log"
)

var db *sql.DB

func Connection(dsn string) {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err.Error())
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
		log.Error("获取数据库conn失败,err: " + err.Error())
		return nil, err
	}
	return conn, nil
}

// @developer caller使用闭包的形式执行sql操作, 而无需关心sql连接的获取和释放
func Using(ctx context.Context, execute func(conn *sql.Conn)) {
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Error("获取数据库conn失败,err: " + err.Error())
	}

	execute(conn)

	conn.Close()
}

// @developer	仅用于单测时Mock数据库
func MockDB(mockDB *sql.DB) {
	db = mockDB
}
