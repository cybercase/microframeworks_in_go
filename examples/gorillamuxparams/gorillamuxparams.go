package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT

func HelloName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, "Hello "+params["name"])
}

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/{name}", HelloName).Methods("GET")
	http.ListenAndServe(":8000", m)
}

// END OMIT
