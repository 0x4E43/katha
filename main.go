package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

var mapData = make(map[string]string)

type ServerResponseV1 struct {
	Key string
	Val string
}

func init() {
	// Read the given input JSON file and store it in cache
	// IMPROVEMENT: do not reload the file on each call
	data, err := os.ReadFile("data/En-Od-v3.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the data to JSON format using EnOdStruct
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	log.Println("Hello World! This is katha📝")

	// Start file server for serving files
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	// Render the HTML file
	http.HandleFunc("/", indexHandler)

	go http.HandleFunc("/search", searchHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.ParseFiles("htmx/index.html"))
	templ.Execute(w, nil)
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
