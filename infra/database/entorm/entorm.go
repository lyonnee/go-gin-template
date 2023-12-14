package entorm

import (
	"context"
	"log"

	"github.com/lyonnee/go-gin-template/ent"
)

var client *ent.Client

func Connection(dbtype, dsn string) {
	var err error
	client, err = ent.Open(dbtype, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// 运行自动迁移工具。
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func Disconnect() {
	client.Close()
}

func GetClient() *ent.Client {
	return client
}
