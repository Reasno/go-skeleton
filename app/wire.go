//+build wireinject

package app

import (
	"github.com/DoNewsCode/core/contract"
	"github.com/google/wire"
	"github.com/nfangxu/core-skeleton/app/handlers"
	"github.com/nfangxu/core-skeleton/app/repositories"
	"gorm.io/gorm"
)

func InjectKernel(db *gorm.DB, env contract.Env) (Kernel, error) {
	panic(
		wire.Build(
			NewGinTransport,
			repositories.NewUserRepository,
			wire.Struct(new(handlers.UsersHandler), "*"),
			wire.Struct(new(Kernel), "*"),
		),
	)
}
