package main

import (
	"flag"
	"github.com/danielmichaels/rmwod/internal/config"
	"github.com/danielmichaels/rmwod/internal/server"
	"github.com/go-chi/httplog"
	"log"
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

	app := &server.Application{
		Config:   cfg,
		Logger:   logger,
		RouteDoc: *routeDoc,
	}
	err := app.Serve()
	if err != nil {
		app.Logger.Error().Err(err).Msg("server failed to start")
	}
	return nil
}
