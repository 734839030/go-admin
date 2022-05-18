package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin/wire"
)

func InitRouters(r *gin.Engine) {
	sysUserController := wire.InitSysUserController()
	r.GET("/sys/user/query/:userId", sysUserController.QueryByUserId)
}
