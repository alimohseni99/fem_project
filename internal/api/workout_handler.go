package api

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
)

type Workouthandler struct{}

func NewWorkoutHandler() *Workouthandler {
	return &Workouthandler{}
}

func (wh *Workouthandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request){
	paramsWorkoutID:= chi.URLParam(r, "id")

	if paramsWorkoutID == "" {
			http.NotFound(w,r)
			return
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
				http.NotFound(w, r)
				return
	}

	fmt.Fprintf(w, "this is the workout id %d\n", workoutID)
}

func (wh *Workouthandler) 