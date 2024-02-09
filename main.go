package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/0x4E43/katha/utils"
	"github.com/gorilla/mux"
)

var mapData = make(map[string]string)

var keys = make([]string, 0)

type ServerResponseV1 struct {
	Key string
	Val string
}

type ApiResponse struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// TODO: currently it seems like init() is being called 2 times
// need to check and fix
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

	// once mapData is loaded get all the keys and store them into the key variabele
	// at this specific point assuming that mapData is loaded, loding all the keys

	if len(mapData) > 0 {
		log.Println(utils.DEBUG("As mapdata size is more then 0, loading all the keys into keys"))
		for k := range mapData {
			keys = append(keys, k)
		}
	}

	log.Println("Keys : ", len(keys))
}

func main() {
	log.Println(utils.INFO("Hello World! This is katha📝"))
	router := mux.NewRouter() // mux from gorrila mux package

	// If any request coming as prefix /images it will be handled by file server
	// image will be served from the image directory
	// any image specified in html file needs to present in image dir
	// http.FileServer(http.Dir("images") -> Responsible for serving image files
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	router.HandleFunc("/", indexHandler)

	router.HandleFunc("/api/options/{opts}", optionHandler)

	router.HandleFunc("/search", searchHandler)
	// http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", router)
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

	val := mux.Vars(req)["opts"]
	// Generate search option based on the key
	option := generateSearchOption(val)

	// Write the generated option as the response
	w.WriteHeader(http.StatusOK)

	fmt.Println("Options: ", option, " Size :", len(option))
	res := ApiResponse{
		Msg:  "Autocomplete data fetched successfully",
		Data: option,
	}
	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Panic(err)
	}
}

func getOdiaMeaning(key string) string {
	return mapData[key]
}

func generateSearchOption(key string) []string {
	log.Print(utils.DEBUG("KEY: " + key))
	opts := make([]string, 0)
	count := 0
	for _, v := range keys {
		if count < 5 { // max 5 key is needed for auto complete
			if strings.Contains(v, key) {
				opts = append(opts, v)
				count = count + 1
			}
		} else {
			break
		}
	}
	return opts
}
