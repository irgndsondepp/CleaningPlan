package main

import (
	"github.com/irgndsondepp/cleaningplan/impl"
	"github.com/irgndsondepp/cleaningplan/interfaces"
)

var benni = impl.NewFlatmate("Benni")
var markus = impl.NewFlatmate("Markus")
var robert = impl.NewFlatmate("Robert")
var people = []interfaces.Person{benni, markus, robert}
var tasks = []interfaces.Task{impl.NewCleanjob("Living Room", impl.SimpleDate(2017, 3, 27), benni), impl.NewCleanjob("Bath", impl.SimpleDate(2017, 5, 28), markus), impl.NewCleanjob("Kitchen", impl.SimpleDate(2016, 12, 31), markus)}
var filename = "./plan.json"

func main() {
	conv := impl.NewJSONConverter()
	persistence := impl.NewFilePersistence(filename, conv, people, tasks)
	cp := impl.NewRotatingCleaningPlan(persistence)
	persistence.Load(cp)
	r := impl.NewResthandler(cp)
	r.StartServe("8080")
}
