package routes

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
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

	mux.HandleFunc("GET /stream", func(w http.ResponseWriter, r *http.Request) {
		ticker := time.NewTicker(1000 * time.Millisecond)
		defer ticker.Stop()

		sse := datastar.NewSSE(w, r)
		element, err := buildElement()
		if err != nil {
			slog.Error("error generating element", slog.String("error", err.Error()))
			return
		}
		if err := sse.PatchElements(element); err != nil {
			slog.Error("failed to patch elements", "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		for {
			select {
			case <-ctx.Done():
				slog.Debug("server closed connection")
				return

			case <-r.Context().Done():
				slog.Debug("client closed connection")
				return

			case <-ticker.C:
				element, err := buildElement()
				if err != nil {
					slog.Error("error generating element", slog.String("error", err.Error()))
					return
				}

				if err := sse.PatchElements(element); err != nil {
					slog.Error("failed to patch elements", "error", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}
	})

	return nil
}

func buildElement() (string, error) {
	bytes := make([]byte, 3)

	n, err := rand.Read(bytes)
	if err != nil || n != len(bytes) {
		return "", err
	}

	hexString := hex.EncodeToString(bytes)

	return fmt.Sprintf(`<span id="outline" style="color:#%s;border:1px solid #%s;border-radius:0.25rem;padding:1rem;display:flex;justify-content:center;">%s</span>`, hexString, hexString, hexString), nil
}
