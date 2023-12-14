package main

import "github.com/lyonnee/go-gin-template/cmd"

//go:generate swag init
//go:generate swag fmt
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	cmd.Execute()
}
