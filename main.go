package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/0x4E43/katha/utils"
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
	log.Println(utils.INFO("Hello World! This is katha📝"))
	// Start file server for serving files
	// file handler kept as it is might not work if uses mux, will see later
	// TODO: handle Image later
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	// Render the HTML file

	// TODO: rewrite server with http requests multiplexere
	mux := http.NewServeMux() // will be possible to route efficiently wilth server mux

	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/api/options", optionHandler)
	mux.HandleFunc("/search", searchHandler)
	// http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Panic(utils.ERROR(err.Error()))
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

func optionHandler(w http.ResponseWriter, req *http.Request) {
	// Extract the key from the URL path
	key := strings.TrimPrefix(req.URL.Path, "/api/options/")
	key = strings.ToLower(key) // Optionally convert to lowercase

	// Generate search option based on the key
	option := generateSearchOption(key)

	// Write the generated option as the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(option))
}

func getOdiaMeaning(key string) string {
	return mapData[key]
}

func generateSearchOption(key string) string {
	log.Println(utils.DEBUG("KEY: " + key))
	return "hello"
}
