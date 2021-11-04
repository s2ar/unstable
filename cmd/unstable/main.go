package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/s2ar/unstable/config"
	"github.com/s2ar/unstable/internal/application"
	"github.com/s2ar/unstable/internal/handler"
	"github.com/s2ar/unstable/internal/pg"
	"github.com/s2ar/unstable/internal/repository"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
)

func InitConfig(filename string) (*config.Configuration, error) {
	cfg, err := config.InitConfig(filename)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load config")
	}

	return cfg, nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	decimal.MarshalJSONWithoutQuotes = true

	app := &cli.App{
		Name:  "unstable",
		Usage: "unstable service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Usage:   "set config file",
				Aliases: []string{"c"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "server",
				Usage:  "start app in server mode",
				Action: cmdStartServer,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		return errors.Wrap(err, "cannot run app")
	}
	return nil
}

func startServer(ctx context.Context, app application.Application, cfg *config.Configuration) {
	r := mux.NewRouter()
	generalMiddleware := []mux.MiddlewareFunc{
		application.WithApp(app),
	}
	api := r.PathPrefix("/api/").Subrouter()
	api.Use(generalMiddleware...)
	handler.GenRouting(api)

	srv := &http.Server{
		Addr:         cfg.Server.Listen,
		WriteTimeout: time.Second * cfg.Server.WriteTimeout,
		ReadTimeout:  time.Second * cfg.Server.ReadTimeout,
		IdleTimeout:  time.Second * cfg.Server.IdleTimeout,
		Handler:      r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(errors.Wrap(err, "cannot start http server"))
		}
	}()

	<-ctx.Done()

	ctxShutdown, cancel := context.WithTimeout(ctx, time.Second*cfg.Server.GracefulShutdown)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Println(errors.Wrap(err, "cannot shutdown http server"))
	}
}

func cmdStartServer(cliContext *cli.Context) error {
	cfg, err := InitConfig(cliContext.String("config"))
	if err != nil {
		return err
	}

	connection, err := repository.NewPG(cfg.Database.DSN)
	if err != nil {
		return errors.Wrap(err, "cannot new connection")
	}

	repos := repository.New(connection)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app, err := application.NewApp(cfg)
	if err != nil {
		return err
	}

	// start migrations
	log.Println("running PostgreSQL migrations")
	err = pg.RunMigrations(cfg)
	if err != nil {
		return fmt.Errorf("RunMigrations failed: %e", err)
	}

	// upload data
	serviceOpendota := app.ServiceOpendota()
	err = serviceOpendota.AddTeams()
	if err != nil {
		return errors.Wrap(err, "cannot add teams")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go startServer(ctx, app, cfg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println("closing")

	return nil
}
