//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"go-admin/internal/domain/sys/repository"
	"go-admin/internal/infrastructure/database"
	"go-admin/internal/interfaces/sys/web"
)

var dataSource = database.NewDataSource()

func InitSysUserController() web.SysUserController {
	panic(wire.Build(wire.Value(dataSource), repository.NewSysUserRepo, web.NewSysUserController))
}
