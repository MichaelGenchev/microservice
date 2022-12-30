package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MichaelGenchev/microservice/internal/businesslogic"
)


type JSONAPIServer struct {
	listenAddr string
	ws  businesslogic.WorkoutsService
}

func NewJSONAPIServer(listenAddr string, ws businesslogic.WorkoutsService) *JSONAPIServer {
	return &JSONAPIServer{
		ws:        ws,
		listenAddr: listenAddr,
	}
}


func (s *JSONAPIServer) HandleGetWorkout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		return fmt.Errorf("invalid id")
	}

	workout, err := s.ws.GetWorkout(id)
	if err != nil {
		return err
	}

	resp := struct {
		Name string
		Exercises []string
	}{
		Name: workout.Name,
		Exercises: workout.Exercises,
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}