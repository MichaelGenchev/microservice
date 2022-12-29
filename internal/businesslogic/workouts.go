package businesslogic

type Workout struct {
	Name string `json:"name"`
	Exercises []string `json:"exercises"`
}

type WorkoutsService interface {
	GetWorkout(id string) (Workout, error)
}