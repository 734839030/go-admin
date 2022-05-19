package main

import (
	"go-admin/internal/infrastructure/middleware"
	"go-admin/internal/interfaces"
)

func main() {
	// 日志系统
	middleware.InitLog()
	// 初始化web
	interfaces.InitRouters()

}
