package pg

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/s2ar/unstable/config"
)

func RunMigrations(cfg *config.Configuration) error {
	if cfg.Database.MigrationsPath == "" {
		return nil
	}

	if cfg.Database.DSN == "" {
		return errors.New("no DSN provided")
	}

	m, err := migrate.New(
		cfg.Database.MigrationsPath,
		cfg.Database.DSN,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
