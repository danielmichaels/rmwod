package server

import (
	"github.com/danielmichaels/rmwod/assets"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"net/http"
)

func (app *Application) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(app.recoverPanic)
	router.Use(app.securityHeaders)
	router.Use(middleware.RealIP)
	router.Use(middleware.Compress(5))
	router.Use(httplog.RequestLogger(app.Logger))

	router.NotFound(app.notFound)
	router.MethodNotAllowed(app.methodNotAllowed)

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	router.Handle("/static/*", fileServer)

	router.Get("/healthz", app.status)
	router.Group(func(web chi.Router) {
		web.Get("/", app.home)
		// web routes
	})
	router.Group(func(api chi.Router) {
		// api routes
	})

	return router
}
