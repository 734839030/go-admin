package interfaces

import (
	"go-admin/internal/infrastructure/middleware"
	"go-admin/wire"
)

func InitRouters() {
	r := middleware.InitGin()
	sysUserController := wire.InitSysUserController()
	r.GET("/sys/user/query/:userId", sysUserController.QueryByUserId)
	middleware.InitHttpServer(r)
}
