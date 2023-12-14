package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lyonnee/go-gin-template/infra/config"
	"github.com/lyonnee/go-gin-template/infra/database/entorm"
)

func Connect() {
	entorm.Connection("mysql", config.Mysql().DSN)
	//mysql.Connection(config.Mysql().DSN)
}

func Disconnect() {
	entorm.Disconnect()
	//mysql.Disconnect()
}
