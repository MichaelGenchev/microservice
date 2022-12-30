package api

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/MichaelGenchev/microservice/internal/businesslogic"
	"github.com/MichaelGenchev/microservice/internal/service/server"
	"github.com/go-chi/chi/v5"
)
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func MakeAPIFunc(fn APIFunc) http.HandlerFunc {
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request) {
		ctx = context.WithValue(ctx, "requestID", rand.Intn(100000000))

		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func SetupRoutesAndRun(router *chi.Mux, listenAddr string, ws businesslogic.WorkoutsService) {
	server := server.NewJSONAPIServer(listenAddr, ws)
	router.Get("/workout/:id", MakeAPIFunc(server.HandleGetWorkout))
	http.ListenAndServe(listenAddr, router)
}


func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}