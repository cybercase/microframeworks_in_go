package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Golab")
}

func main() {
	m := http.NewServeMux() // Standard mux
	m.HandleFunc("/", Hello)

	n := negroni.New(negroni.Recovery(), negroni.NewLogger(), negroni.Static())
	n.UseHandler(m)

	http.ListenAndServe(":8000", n)
}
