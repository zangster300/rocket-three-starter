package routes

import (
	"context"
	"math"
	"net/http"
	"time"

	"rocket-three-starter/templating"

	"github.com/starfederation/datastar-go/datastar"
)

func setupIndexRoute(ctx context.Context, mux *http.ServeMux) error {
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		if err := templating.Index().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("GET /updates", func(w http.ResponseWriter, r *http.Request) {

		const (
			amplitude = 3.0
			duration  = 30 * time.Second
			fps       = 60
		)

		ticker := time.NewTicker(time.Second / fps)
		defer ticker.Stop()

		sse := datastar.NewSSE(w, r, datastar.WithCompression(datastar.WithBrotli()))

		start := time.Now()
		for {
			select {
			case <-r.Context().Done():
				return
			case <-ticker.C:
				t := time.Since(start).Seconds()
				y := amplitude * math.Sin(2*math.Pi*t/duration.Seconds())

				signals := struct {
					Y float64 `json:"y"`
				}{
					Y: y,
				}

				sse.MarshalAndPatchSignals(signals)
			}
		}

	})

	return nil
}
