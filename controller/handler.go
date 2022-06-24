package controller

import "github.com/LyonNee/app-layout/controller/user"

// 初始化controller
func Initialize() {
	user.Initialize()
	// 每当添加新的controller, 需要在这里执行初始化
	// order.Initialize()
}
