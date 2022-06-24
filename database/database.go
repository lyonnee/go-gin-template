package database

import (
	"github.com/LyonNee/app-layout/database/mysql"
)

func Connect() {
	mysql.Connection()
}

func Disconnect() {
	mysql.Disconnect()
}
