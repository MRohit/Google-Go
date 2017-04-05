package main

import (
    "fmt"
    "log"
    "net/http"
	"mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)
	router.HandleFunc("/json/{todoId}", getJsonData)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
func getJsonData(w http.ResponseWriter, r* http.Request) {
	json := "{\"key\":\"12312312\",\"value\":\"Pass\"}"
	
}
func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}