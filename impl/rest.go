package impl

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irgndsondepp/cleaningplan/impl/rest"
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
	routes := r.createRoutes()
	for _, route := range routes {
		var hf http.Handler
		hf = route.HandlerFunc
		hf = rest.Logger(hf, route.Name)

		r.router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(hf)
	}
}

func (r *Resthandler) createRoutes() []*rest.Route {
	var routes []*rest.Route
	routes = append(routes, rest.NewRoute("Get tasks", "GET", "/", r.print))
	routes = append(routes, rest.NewRoute("Get tasks", "GET", "/tasks", r.print))
	routes = append(routes, rest.NewRoute("Set task done", "PUT", "/tasks/{taskname}/{assignee}", r.setJobAsDone))
	return routes
}

func (r *Resthandler) setJobAsDone(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	task := NewSimpleTask(vars["taskname"], vars["assignee"])
	w.Header().Add("Content-Type", "application/json")
	err := r.cp.MarkTaskAsDone(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error setting job as done: %v\n", err)
	} else {
		fmt.Fprintf(w, "Set job as done: %v for %v\n", task.Name, task.Assignee.GetName())
		updatedPlan, err := r.conv.ConvertTo(r.cp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error saving file: %v\n", err)
		} else {
			w.WriteHeader(http.StatusAccepted)
			w.Write(updatedPlan)
		}
	}
}

func (r *Resthandler) print(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, err := r.conv.ConvertTo(r.cp)
	if err != nil {
		//todo error handling
	}
	w.Write(bytes)
}
