package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/danielmichaels/rmwod/assets"
	"github.com/danielmichaels/rmwod/internal/config"
	"github.com/danielmichaels/rmwod/internal/database"
	"github.com/danielmichaels/rmwod/internal/server"
	"github.com/go-chi/httplog"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

var routeDoc = flag.Bool("docgen", false, "Generate route documentation using docgen")

func main() {
	err := run()
	if err != nil {
		log.Fatalln("server failed to start:", err)
	}
}

func run() error {
	flag.Parse() // doc gen flag parser
	cfg := config.AppConfig()
	logger := httplog.NewLogger("rmwod", httplog.Options{
		JSON:     cfg.AppConf.LogJson,
		Concise:  cfg.AppConf.LogConcise,
		LogLevel: cfg.AppConf.LogLevel,
	})
	if cfg.AppConf.LogCaller {
		logger = logger.With().Caller().Logger()
	}

	db, err := openDB(cfg)
	if err != nil {
		logger.Error().Err(err).Msg("unable to connect to database")
		os.Exit(1)
	}

	app := &server.Application{
		Config:   cfg,
		Logger:   logger,
		RouteDoc: *routeDoc,
		Db:       database.New(db),
	}
	database.Seed()
	err = app.Serve()
	if err != nil {
		app.Logger.Error().Err(err).Msg("server failed to start")
	}
	return nil
}

func openDB(cfg *config.Conf) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.Db.DbName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Db.DatabaseConnectionContext)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	err = goose.SetDialect("sqlite3")
	if err != nil {
		return nil, err
	}
	goose.SetBaseFS(assets.EmbeddedFiles)
	err = goose.Up(db, "migrations")
	if err != nil {
		return nil, err
	}

	return db, nil
}
