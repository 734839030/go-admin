package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/internal/domain/sys/repository"
	"go-admin/internal/infrastructure/common/api"
	"strconv"
)

type SysUserController struct {
	SysUserRepo *repository.SysUserRepo
}

func NewSysUserController(sysUserRepo *repository.SysUserRepo) SysUserController {
	return SysUserController{sysUserRepo}
}

func (w SysUserController) QueryByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		api.RespErrorWithCode(c, api.PARAM_INVALID, err.Error())
		return
	}
	sysUser, err := w.SysUserRepo.FindById(userId)
	if err != nil {
		api.RespError(c, err.Error())
		return
	}
	if sysUser != nil {
		api.RespOk(c, sysUser)
	} else {
		api.RespErrorWithCode(c, api.RECORD_NOT_FOUND, fmt.Sprintf("user [%d] not found", userId))
	}
}
