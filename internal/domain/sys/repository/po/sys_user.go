package po

import (
	"go-admin/internal/infrastructure/common"
)

type SysUser struct {
	UserId   int    `gorm:"primaryKey" json:"userId"`
	Username string `json:"username"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Photo    string `json:"photo"`
	Email    string `json:"email"`
	common.BasePO
}
