package main

import "github.com/LyonNee/app-layout/cmd"

//go:generate swag init
//go:generate swag fmt

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	cmd.Execute()
}
