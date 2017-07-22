package impl

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irgndsondepp/cleaningplan/interfaces"
)

type Resthandler struct {
	router *mux.Router
	cp     interfaces.Plan
	conv   *JSONConverter
}

func NewResthandler(c interfaces.Plan) *Resthandler {
	return &Resthandler{
		router: mux.NewRouter().StrictSlash(true),
		cp:     c,
		conv:   NewJSONConverter(),
	}
}

func (r *Resthandler) StartServe(port string) {
	r.initHandlerFunctions()
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), r.router)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Resthandler) initHandlerFunctions() {
	r.router.HandleFunc("/", r.print)
	r.router.HandleFunc("/tasks", r.print)
	r.router.HandleFunc("/tasks/{taskname}/{assignee}", r.setJobAsDone)
}

func (r *Resthandler) setJobAsDone(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	task := NewSimpleTask(vars["taskname"], vars["assignee"])
	err := r.cp.MarkTaskAsDone(task)
	if err != nil {
		fmt.Fprintf(w, "Error setting job as done: %v\n", err)
	} else {
		fmt.Fprintf(w, "Set job as done: %v for %v\n", task.Name, task.Assignee)
		updatedPlan, err := r.conv.ConvertTo(r.cp)
		if err != nil {
			fmt.Fprintf(w, "Error saving file: %v\n", err)
		} else {
			w.Write(updatedPlan)
		}
	}
}

func (r *Resthandler) print(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content", "application/json")
	bytes, err := r.conv.ConvertTo(r.cp)
	if err != nil {
		//todo error handling
	}
	w.Write(bytes)
}
