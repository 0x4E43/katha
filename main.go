package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

// Struct for odia english json data
type EnOdStruct struct {
	EN_key string // key will be english word
	OD_val string // value will be meaning of the englisg word in odia lipi
}

func main() {
	fmt.Println("Hello World!")
	//Read the given input Json file and store it in cache
	//IMPROVEMENT: do not reload the the file on each call
	data, err := os.ReadFile("En-Od-v3.json")

	if err != nil {
		panic(err)
	}
	//ummarshal the data to json format using Odia using EnOdStruct
	// var jsonData EnOdStruct

	var mapData = make(map[string]string)
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		panic(err)
	}
	startTime := time.Now().UnixMilli()
	fmt.Println("Start Time: ", startTime)
	fmt.Println(mapData["elephant"])
	fmt.Println("End Time: ", (time.Now().UnixMilli() - startTime))

	//render the HTML File
	http.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	log.Println("Server Running at : ", 8080)

}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.ParseFiles("htmx/index.html"))
	templ.Execute(w, nil)
}

func test() string {
	return "nimai"
}
