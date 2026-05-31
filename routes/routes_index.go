package routes

import (
	"context"
	"net/http"

	"rocket-three-starter/templating"
)

func setupIndexRoute(ctx context.Context, mux *http.ServeMux) error {
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		if err := templating.Index().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	return nil
}
