package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"io/ioutil"

	"os"

	"github.com/irgndsondepp/cleaningplan"
)

var cleaningPlan = cleaningplan.NewCleaningPlan()
var fileNameXML = "./plan.xml"
var fileNameJSON = "./plan.json"

func main() {
	loadPlanFromFile()
	http.HandleFunc("/", handleInput)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loadPlanFromFile() {
	ex, err := exists(fileNameXML)
	if err != nil {
		fmt.Println("Error checking if file exists: %v", err)
	}
	if !ex {
		cleaningPlan = cleaningplan.InitCleaningPlan()
	}
	bytes, err := ioutil.ReadFile(fileNameXML)
	if err != nil {
		return
	}
	cleaningPlan, err = cleaningplan.FromXML(bytes)
	if err != nil {
		fmt.Printf("Error decoding file: %v\n", err)
		cleaningPlan = cleaningplan.InitCleaningPlan()
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func savePlanToFile() {
	bytes, err := cleaningPlan.ToXML()
	if err != nil {
		fmt.Printf("Error trying to Encode Plan: %v\n", err)
	}
	err = ioutil.WriteFile(fileNameXML, bytes, 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v", err)
	}
}

func handleInput(w http.ResponseWriter, req *http.Request) {
	reqUri := req.RequestURI
	if strings.Contains(reqUri, "/done") {
		setJobAsDone(w, reqUri)
	} else {
		printCleaningPlan(w, req)
	}
}

func setJobAsDone(w http.ResponseWriter, reqUri string) {
	res := parseInput(reqUri)
	if len(res) != 2 {
		fmt.Fprintf(w, "Input was in wrong format: %v", reqUri)
	} else {
		err := cleaningPlan.MarkJobAsDone(res[0], res[1])
		if err != nil {
			fmt.Fprintf(w, "Error setting job as done: %v", err)
		} else {
			fmt.Fprintf(w, "Set job as done: %v for %v\n", res[1], res[0])
			fmt.Fprintln(w, cleaningPlan.ToString())
			savePlanToFile()
		}
	}
}

func parseInput(url string) []string {
	var res []string
	for i, input := range strings.Split(url, "/") {
		if i > 1 {
			res = append(res, input)
		}
	}
	return res
}

func printCleaningPlan(w http.ResponseWriter, req *http.Request) {
	err := printJSON(w)
	if err != nil {
		fmt.Fprintln(w, err)
	}
}

func printJSON(w http.ResponseWriter) error {
	w.Header().Add("Content", "application/json")
	bytes, err := cleaningPlan.ToJSON()
	if err != nil {
		return err
	}
	w.Write(bytes)
	return nil
}

func printXML(w http.ResponseWriter) error {
	w.Header().Add("Content", "text/xml")
	bytes, err := cleaningPlan.ToXML()
	if err != nil {
		return err
	}
	w.Write(bytes)
	return nil
}
