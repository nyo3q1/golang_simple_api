package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	peoples, err := GetPeoples()
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	fmt.Fprintf(w, "%v", peoples)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
