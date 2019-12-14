package main

import (
	"fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
    peoples := GetPeoples()
	fmt.Fprintf(w, "%s", peoples)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", Home)
    //r.HandleFunc("/products", ProductsHandler)
    //r.HandleFunc("/articles", ArticlesHandler)
	
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}