package app

import (
	"github.com/DoNewsCode/core/otgorm"
	"github.com/gorilla/mux"
	"github.com/nfangxu/core-skeleton/app/repositories"
	"net/http"
)

type Kernel struct {
	users GinTransport
}

func (a Kernel) ProvideHttp(router *mux.Router) {
	router.PathPrefix("/users").Handler(http.StripPrefix("/users", a.users))
}

func (a Kernel) ProvideMigration() []*otgorm.Migration {
	return repositories.ProviderMigrations()
}

func (a Kernel) ProvideSeed() []*otgorm.Seed {
	return repositories.ProvideSeeders()
}
