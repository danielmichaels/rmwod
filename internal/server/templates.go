package server

import (
	"github.com/danielmichaels/rmwod/internal/version"
	"net/http"
)

func (app *Application) newTemplateData(r *http.Request) map[string]any {
	data := map[string]any{
		"Version":   version.Get(),
		"Plausible": app.Config.Secrets.Plausible,
	}

	return data
}
