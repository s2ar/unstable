//go:generate mockgen -destination=./mocks.go -source=./app.go -package=application
package application

import (
	"github.com/pkg/errors"
	"github.com/s2ar/unstable/config"
	"github.com/s2ar/unstable/internal/repository"
	repositoryOpendota "github.com/s2ar/unstable/internal/repository/opendota"
	"github.com/s2ar/unstable/internal/service/opendota"
)

type Application interface {
	RepositoryOpendota() repositoryOpendota.Repository
	Config() *config.Configuration
	Connection() *repository.DB
	ServiceOpendota() opendota.Opendota
}

type app struct {
	connection *repository.DB
	cfg        *config.Configuration
}

func (a *app) Connection() *repository.DB {
	return a.connection
}

func (a *app) Config() *config.Configuration {
	return a.cfg
}

func (a *app) RepositoryOpendota() repositoryOpendota.Repository {
	return repositoryOpendota.New(a.connection)
}

func (a *app) ServiceOpendota() opendota.Opendota {
	return opendota.New(a.RepositoryOpendota(), a.Config())
}

func NewApp(cfg *config.Configuration) (Application, error) {
	connection, err := repository.New(cfg.Database.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "cannot new app")
	}
	app := &app{
		connection: connection,
		cfg:        cfg,
	}
	return app, nil
}
