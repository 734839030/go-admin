package repository

import (
	"errors"
	"fmt"
	"go-admin/internal/domain/sys/repository/po"
	"go-admin/internal/infrastructure/database"
	"gorm.io/gorm"
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

func (r *SysUserRepo) UpdateById(po *po.SysUser) error {
	if po.UserId == 0 {
		return errors.New("Primary must set")
	}
	return r.dataSource.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Save(po)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected != 1 {
			return errors.New(fmt.Sprintf("update failed RowsAffected %d", result.RowsAffected))
		}
		return nil
	})
}

func (r *SysUserRepo) DeleteById(ids []int) error {
	count := len(ids)
	if count == 0 {
		return errors.New("ids must set")
	}
	return r.dataSource.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&po.SysUser{}, ids)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected != int64(count) {
			return errors.New(fmt.Sprintf("update failed RowsAffected %d ,Expected %d", result.RowsAffected, count))
		}
		return nil
	})
}
