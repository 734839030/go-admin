package repository

import (
	"encoding/json"
	"go-admin/internal/domain/sys/repository/po"
	"go-admin/internal/infrastructure/database"
	"testing"
)

var r = &SysUserRepo{database.NewDataSource()}

func TestSysUserRepo_Find(t *testing.T) {
	userId := uint(1)
	t.Run("Find", func(t *testing.T) {
		if sysUser, err := r.FindById(userId); err == nil {
			if marshal, err := json.Marshal(sysUser); err != nil {
				t.Error(err)
			} else {
				t.Logf("sysuser:%s", marshal)
			}
		} else {
			t.Error(err)
		}
	})
}
func TestSysUserRepo_Save(t *testing.T) {
	p := &po.SysUser{Username: "test", Name: "hello"}
	t.Run("Save", func(t *testing.T) {
		if _, err := r.Save(p); err != nil {
			t.Error(err)
		}
	})
}
