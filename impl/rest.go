package impl

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/irgndsondepp/cleaningplan/impl/rest"
	"github.com/irgndsondepp/cleaningplan/interfaces"
)

type Resthandler struct {
	router *mux.Router
	cp     interfaces.Plan
}

func NewResthandler(c interfaces.Plan) *Resthandler {
	return &Resthandler{
		router: mux.NewRouter().StrictSlash(true),
		cp:     c,
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
	routes = append(routes, rest.NewRoute("Get filtered tasks", "GET", "/tasks/{assignee}", r.filterTasks))
	routes = append(routes, rest.NewRoute("Set task done", "PUT", "/tasks/{taskname}/{assignee}", r.setJobAsDone))
	return routes
}

func (r *Resthandler) filterTasks(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["assignee"]
	tasks, err := r.cp.FilterTasks(name)
	if err != nil {
		returnError(err, http.StatusBadRequest, w)
	}
	encodeResponse(tasks, http.StatusOK, w)
}

func (r *Resthandler) setJobAsDone(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	task := NewSimpleTask(vars["taskname"], vars["assignee"])
	err := r.cp.MarkTaskAsDone(task)
	if err != nil {
		returnError(fmt.Errorf("Error setting job as done: %v", err), http.StatusBadRequest, w)
	} else {
		encodeResponse(r.cp.GetTasks(), http.StatusAccepted, w)
	}
}

func (r *Resthandler) print(w http.ResponseWriter, req *http.Request) {
	encodeResponse(r.cp.GetTasks(), http.StatusOK, w)
}

func encodeResponse(v interface{}, statusCode int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	bytes, err := json.Marshal(v)
	if err != nil {
		returnError(err, http.StatusInternalServerError, w)
	}
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

func returnError(err error, errorCode int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	b, _ := json.Marshal(&ErrorResponse{ErrorCode: errorCode, Message: err.Error()})
	w.Write(b)
}

type ErrorResponse struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
}
