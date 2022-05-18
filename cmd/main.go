package main

import (
	"go-admin/configs"
	"go-admin/internal/infrastructure/middleware"
	"go-admin/internal/interfaces"
)

func main() {
	// 始化配置
	configs.InitConfig()
	// 日志系统
	middleware.InitLog()
	// 初始化web
	interfaces.InitRouters()

}
