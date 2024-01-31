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
	EN_key string // key will be the English word
	OD_val string // value will be the meaning of the English word in Odia lipi
}

func main() {
	fmt.Println("Hello World!")

	// Read the given input JSON file and store it in cache
	// IMPROVEMENT: do not reload the file on each call
	data, err := os.ReadFile("En-Od-v3.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal the data to JSON format using EnOdStruct
	var mapData = make(map[string]string)
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		panic(err)
	}

	startTime := time.Now().UnixMilli()
	fmt.Println("Start Time: ", startTime)
	fmt.Println(mapData["elephant"])
	fmt.Println("End Time: ", (time.Now().UnixMilli() - startTime))

	// Start file server for serving files
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	// Render the HTML file
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
