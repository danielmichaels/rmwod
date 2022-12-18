package server

import (
	"context"
	"github.com/danielmichaels/rmwod/internal/response"
	"github.com/danielmichaels/rmwod/internal/version"
	"net/http"
)

func (app *Application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status":  "OK",
		"Version": version.Get(),
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	video, err := app.Db.RandomWod(context.Background())
	if err != nil {
		app.Logger.Error().Err(err).Msg("error: could not retrieve video")
		app.notFound(w, r)
		return
	}

	data["Video"] = video

	err = response.Page(w, http.StatusOK, data, "pages/home.tmpl")
	if err != nil {
		app.serverError(w, r, err)
	}
}
