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

var mapData = make(map[string]string)

func init() {
	// Read the given input JSON file and store it in cache
	// IMPROVEMENT: do not reload the file on each call
	data, err := os.ReadFile("En-Od-v3.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal the data to JSON format using EnOdStruct
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		panic(err)
	}

}
func main() {
	fmt.Println("Hello World!")

	startTime := time.Now().UnixMilli()
	fmt.Println("Start Time: ", startTime)
	fmt.Println(mapData["elephant"])
	fmt.Println("End Time: ", (time.Now().UnixMilli() - startTime))

	// Start file server for serving files
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	// Render the HTML file
	http.HandleFunc("/", indexHandler)

	go http.HandleFunc("/search", searchHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	log.Println("Server Running at : ", 8080)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.ParseFiles("htmx/index.html"))
	templ.Execute(w, nil)
}

type ServerResponseV1 struct {
	Key string
	Val string
}

func searchHandler(w http.ResponseWriter, req *http.Request) {
	//get the request from the urls
	searchParam := req.URL.Query().Get("search")
	//TODO: eed to do lower case for key
	//find in map dictionary and render the result
	res := ServerResponseV1{Key: searchParam, Val: getOdiaMeaning(searchParam)}
	templ := template.Must(template.ParseFiles("htmx/test.html"))
	templ.Execute(w, res)
}

func getOdiaMeaning(key string) string {
	return mapData[key]

}
