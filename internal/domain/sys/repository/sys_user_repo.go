package repository

import (
	"go-admin/internal/domain/sys/repository/po"
	"go-admin/internal/infrastructure/database"
)

type SysUserRepo struct {
	dataSource *database.DataSource
}

func NewSysUserRepo(dataSource *database.DataSource) *SysUserRepo {
	return &SysUserRepo{dataSource}
}

func (r *SysUserRepo) FindById(userId int) (*po.SysUser, error) {
	var sysUser po.SysUser
	result := r.dataSource.DB.Limit(1).Find(&sysUser, &po.SysUser{
		UserId: userId,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, result.Error
	}
	return &sysUser, result.Error
}

func (r *SysUserRepo) Save(po *po.SysUser) (int64, error) {
	result := r.dataSource.DB.Create(po)
	return result.RowsAffected, result.Error
}
