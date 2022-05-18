package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/internal/domain/sys/repository"
	"go-admin/internal/infrastructure/common/api"
	"net/http"
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
		c.JSON(http.StatusOK, api.RespError(err.Error()))
		return
	}
	sysUser, err := w.SysUserRepo.FindById(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, api.RespError(err.Error()))
		return
	}
	if sysUser != nil {
		c.JSON(http.StatusOK, api.RespOk(sysUser))
	} else {
		c.JSON(http.StatusOK, api.RespErrorWithCode(api.RECORD_NOT_FOUND, fmt.Sprintf("user [%d] not found", userId)))
	}
}
